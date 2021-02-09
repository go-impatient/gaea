package middleware

import (
	"net/http"
	"strings"

	"moocss.com/gaea/pkg/conf"
)

// Cors middleware
func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log.Get(context.Background()).Println("cors before")

		origin := r.Header.Get("Origin")
		suffix := conf.Get("app.cors_origin_suffix")

		if origin != "" && suffix != "" && strings.HasSuffix(origin, suffix) {
			w.Header().Add("Access-Control-Allow-Origin", origin)
			w.Header().Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS") //允许请求方法
			w.Header().Add("Access-Control-Allow-Credentials", "true")         //设置为true，允许ajax异步请求带cookie信息
			w.Header().Add("Access-Control-Allow-Headers", "Origin,No-Cache,X-Requested-With,If-Modified-Since,Pragma,Last-Modified,Cache-Control,Expires,Content-Type,Access-Control-Allow-Credentials,DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Cache-Webcdn,Content-Length,Authorization,Token,Cookie")
		}

		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)

		// log.Get(context.Background()).Println("cors after")
	})
}
