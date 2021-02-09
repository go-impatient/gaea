package server

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"moocss.com/gaea/pkg"
	"moocss.com/gaea/pkg/conf"
	"moocss.com/gaea/pkg/log"
	"moocss.com/gaea/pkg/middleware"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var server *http.Server

// 从 http 标准库搬来的
type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (net.Conn, error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return nil, err
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}

var logger = log.Get(context.Background())

func main() {
	reload := make(chan int, 1)
	stop := make(chan os.Signal, 1)

	conf.OnConfigChange(func() { reload <- 1 })
	conf.WatchConfig()

	// 注入App需要的依赖
	// InitApp()

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)

	startServer()

	for {
		select {
		case <-reload:
			pkg.Reset()
		case sg := <-stop:
			stopServer()
			// 仿 nginx 使用 HUP 信号重载配置
			if sg == syscall.SIGHUP {
				startServer()
			} else {
				pkg.Stop()
				return
			}
		}
	}
}

func startServer() {
	logger.Info("start server")

	rand.Seed(int64(time.Now().Nanosecond()))

	mux := http.NewServeMux()

	timeout := 600 * time.Millisecond
	initMux(mux, isInternal)

	if isInternal {
		initInternalMux(mux)

		if d := conf.GetDuration("INTERNAL_API_TIMEOUT"); d > 0 {
			timeout = d
		}
	} else {
		if d := conf.GetDuration("OUTER_API_TIMEOUT"); d > 0 {
			timeout = d
		}
	}

	// 中间件
	m := middleware.NewMiddleware()
	m.Use(middleware.Recovery)
	m.Use(middleware.Cors)
	panicHandler := m.Add(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r)
	}))

	handler := http.TimeoutHandler(panicHandler, timeout, "timeout")

	prefix := conf.Get("RPC_PREFIX")
	if prefix == "" {
		prefix = "/api"
	}
	if prefix != "/" {
		handler = http.StripPrefix(prefix, handler)
	}
	http.Handle("/", handler)

	metricsHandler := promhttp.Handler()
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		pkg.GatherMetrics()

		metricsHandler.ServeHTTP(w, r)
	})

	http.HandleFunc("/monitor/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	addr := fmt.Sprintf(":%d", port)
	server = &http.Server{
		IdleTimeout: 60 * time.Second,
	}

	// 配置下发可能会多次触发重启，必须等待 Listen() 调用成功
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		// 本段代码基本搬自 http 标准库
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			panic(err)
		}
		wg.Done()

		err = server.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
		if err != http.ErrServerClosed {
			panic(err)
		}
	}()

	wg.Wait()
}

func stopServer() {
	logger.Info("stop server")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal(err)
	}

	pkg.Reset()
}
