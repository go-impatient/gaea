// Package conf 提供最基础的配置加载功能
package conf

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"moocss.com/gaea/pkg/conf/file"
	"moocss.com/gaea/pkg/log"
)

var files map[string]*Conf

var (
	// Hostname 主机名
	Hostname = "localhost"
	// AppID 获取 APP_ID
	AppID = "gaea_app"
	// IsDevEnv 开发环境标志
	IsDevEnv = false
	// IsUatEnv 集成环境标志
	IsUatEnv = false
	// IsProdEnv 生产环境标志
	IsProdEnv = false
	// Env 运行环境
	Env = "dev"
	// Zone 服务区域
	Zone = "sh001"
)

// Conf .
type Conf struct {
	viper *viper.Viper
}

// Init .
func Init(logger log.Logger) {
	log := log.NewHelper("conf", logger)

	Hostname, _ = os.Hostname()
	if appID := os.Getenv("APP_ID"); appID != "" {
		AppID = appID
	} else {
		log.Warn("env APP_ID is empty")
	}

	if env := os.Getenv("DEPLOY_ENV"); env != "" {
		Env = env
	} else {
		log.Warn("env DEPLOY_ENV is empty")
	}

	if zone := os.Getenv("ZONE"); zone != "" {
		Zone = zone
	} else {
		log.Warn("env ZONE is empty")
	}

	switch Env {
	case "prod", "pre":
		IsProdEnv = true
	case "uat":
		IsUatEnv = true
	default:
		IsDevEnv = true
	}

	confPath := os.Getenv("CONF_PATH")

	var err error
	if confPath == "" {
		log.Warn("env CONF_PATH is empty")
		if confPath, err = os.Getwd(); err != nil {
			panic(err)
		}
		log.Infow(
			"use default conf path",
			confPath,
		)
		confPath += "/config"
	}

	Load(confPath, log)
}

// Load .
func Load(confPath string, log *log.Helper) {
	// 目标
	src := file.NewFile(confPath)
	// 目标下的所有配置文件
	fs, err := src.Load()
	if err != nil {
		panic(err)
	}
	files = make(map[string]*Conf, len(fs))
	for _, f := range fs {
		if !strings.HasSuffix(f, ".yaml") {
			continue
		}

		log.Infof("config file %s", f)

		v := viper.New()
		// Config's format: "json" | "toml" | "yaml" | "yml"
		// v.SetConfigType("yaml")
		v.SetConfigFile(f)
		if err := v.ReadInConfig(); err != nil {
			log.Warnf("Using config file: %s [%s]\n", viper.ConfigFileUsed(), err)
			panic(err)
		}

		// 读取匹配的环境变量
		v.AutomaticEnv()

		name := strings.TrimSuffix(path.Base(f), ".yaml")
		files[name] = &Conf{v}
	}
}

// GetFloat64 获取浮点数配置
func GetFloat64(key string) float64 { return File("gaea").GetFloat64(key) }
func (c *Conf) GetFloat64(key string) float64 {
	return c.viper.GetFloat64(key)
}

// Get 获取字符串配置
func Get(key string) string { return File("gaea").Get(key) }
func (c *Conf) Get(key string) string {
	return c.viper.GetString(key)
}

// GetStrings 获取字符串列表
func GetStrings(key string) (s []string) { return File("gaea").GetStrings(key) }
func (c *Conf) GetStrings(key string) (s []string) {
	value := Get(key)
	if value == "" {
		return
	}

	for _, v := range strings.Split(value, ",") {
		s = append(s, v)
	}
	return
}

// GetInt32s 获取数字列表
// 1,2,3 => []int32{1,2,3}
func GetInt32s(key string) (s []int32, err error) { return File("gaea").GetInt32s(key) }
func (c *Conf) GetInt32s(key string) (s []int32, err error) {
	s64, err := GetInt64s(key)
	for _, v := range s64 {
		s = append(s, int32(v))
	}
	return
}

// GetInt64s 获取数字列表
func GetInt64s(key string) (s []int64, err error) { return File("gaea").GetInt64s(key) }
func (c *Conf) GetInt64s(key string) (s []int64, err error) {
	value := Get(key)
	if value == "" {
		return
	}

	var i int64
	for _, v := range strings.Split(value, ",") {
		i, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			return
		}
		s = append(s, i)
	}
	return
}

// GetInt 获取整数配置
func GetInt(key string) int { return File("gaea").GetInt(key) }
func (c *Conf) GetInt(key string) int {
	return c.viper.GetInt(key)
}

// GetInt32 获取 int32 配置
func GetInt32(key string) int32 { return File("gaea").GetInt32(key) }
func (c *Conf) GetInt32(key string) int32 {
	return c.viper.GetInt32(key)
}

// GetInt64 获取 int64 配置
func GetInt64(key string) int64 { return File("gaea").GetInt64(key) }
func (c *Conf) GetInt64(key string) int64 {
	return c.viper.GetInt64(key)
}

// GetDuration 获取时间配置
func GetDuration(key string) time.Duration { return File("gaea").GetDuration(key) }
func (c *Conf) GetDuration(key string) time.Duration {
	return c.viper.GetDuration(key)
}

// GetTime 查询时间配置
// 默认时间格式为 "2006-01-02 15:04:05"，conf.GetTime("FOO_BEGIN")
// 如果需要指定时间格式，则可以多传一个参数，conf.GetString("FOO_BEGIN", "2006")
//
// 配置不存在或时间格式错误返回**空时间对象**
// 使用本地时区
func GetTime(key string, args ...string) time.Time { return File("gaea").GetTime(key, args...) }
func (c *Conf) GetTime(key string, args ...string) time.Time {
	fmt := "2006-01-02 15:04:05"
	if len(args) == 1 {
		fmt = args[0]
	}

	t, _ := time.ParseInLocation(fmt, c.viper.GetString(key), time.Local)
	return t
}

// GetBool 获取配置布尔配置
func GetBool(key string) bool { return File("gaea").GetBool(key) }
func (c *Conf) GetBool(key string) bool {
	return c.viper.GetBool(key)
}

// Sub 返回新的Viper实例，代表该实例的子节点。
func Sub(key string) (*viper.Viper, error) { return File("gaea").Sub(key) }
func (c *Conf) Sub(key string) (*viper.Viper, error) {
	if app := c.viper.Sub(key); app != nil {
		return app, nil
	}
	return nil, fmt.Errorf("No found `%s` in the configuration", key)
}

// Set 设置配置，仅用于测试
func Set(key string, value string) { File("gaea").Set(key, value) }
func (c *Conf) Set(key string, value string) {
	c.viper.Set(key, value)
}

// File 根据文件名获取对应配置对象
// 目前仅支持 toml 文件，不用传扩展名
// 如果要读取 foo.toml 配置，可以 File("foo").Get("bar")
func File(name string) *Conf {
	res, _ := files[name]
	if res == nil {
		res = &Conf{viper: &viper.Viper{}}
	}
	return res
}

// OnConfigChange 注册配置文件变更回调
// 需要在 WatchConfig 之前调用
func OnConfigChange(run func()) {
	for _, v := range files {
		v.viper.OnConfigChange(func(in fsnotify.Event) { run() })
	}
}

// WatchConfig 启动配置变更监听，业务代码不要调用。
func WatchConfig() {
	for _, v := range files {
		v.viper.WatchConfig()
	}
}
