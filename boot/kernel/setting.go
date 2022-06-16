package kernel

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"osier/config"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

// 设置路由日志
func setRouteLog(m, p, n string, h int) {

	fmt.Fprintf(gin.DefaultWriter, "[%v]API: %v -> %v %v\n", m, p, n, h)
}

// Zap-日志编码器
func getEncoder() zapcore.Encoder {

	enConfig := zap.NewProductionEncoderConfig()

	enConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	enConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewConsoleEncoder(enConfig)
}

// Zap-日志位置
func getWrite(l config.Log) zapcore.WriteSyncer {

	now := time.Now()
	name := fmt.Sprintf("%s/%d%02d/%02d.log", l.Path, now.Year(), now.Month(), now.Day())
	info := &lumberjack.Logger{
		Filename:   name,
		MaxSize:    2,
		MaxBackups: 0,
		MaxAge:     0,
		Compress:   false,
	}

	return zapcore.AddSync(info)
}

// Zap-日志级别
func getLevel(l config.Log) zapcore.Level {

	level := zapcore.DebugLevel
	switch l.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "dpanic":
		level = zapcore.DPanicLevel
	case "panic":
		level = zapcore.PanicLevel
	case "fatal":
		level = zapcore.FatalLevel
	}

	return level
}

// 日志中间件
func getLogger(l *zap.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				l.Error(e)
			}
		} else {
			fields := []zapcore.Field{
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.Duration("latency", latency),
			}
			l.Info(path, fields...)
		}
	}
}

// 中间件-重启
func getRecovery(l *zap.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					l.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)

					c.Error(err.(error))
					c.Abort()
					return
				}

				l.Error("[Recovery from panic]",
					zap.Time("time", time.Now()),
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
					zap.String("stack", string(debug.Stack())),
				)

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

// 获取配置文件信息
func getCfg[T any](cls T, vpr *viper.Viper) T {

	dfKey := "df"
	rt := reflect.ValueOf(&cls).Elem()
	vt := rt.Type()
	name := strings.ToLower(vt.Name())

	for i := 0; i < rt.NumField(); i++ {

		f := rt.Field(i)
		vf := vt.Field(i)
		ckey := name + "." + strings.ToLower(vf.Name)

		switch f.Kind() {

		case reflect.Bool:
			if vpr.IsSet(ckey) {
				i := vpr.GetBool(ckey)
				f.SetBool(i)
			} else {
				tag, ok := vf.Tag.Lookup(dfKey)
				if i, err := strconv.ParseBool(tag); ok && err == nil {
					f.SetBool(i)
				}
			}

		case reflect.String:
			if vpr.IsSet(ckey) {
				i := vpr.GetString(ckey)
				f.SetString(i)
			} else {
				tag, ok := vf.Tag.Lookup(dfKey)
				if ok {
					f.SetString(tag)
				}
			}

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if vpr.IsSet(ckey) {
				i := vpr.GetInt64(ckey)
				f.SetInt(i)
			} else {
				tag, ok := vf.Tag.Lookup(dfKey)
				if i, err := strconv.ParseInt(tag, 10, 64); ok && err == nil {
					f.SetInt(i)
				}
			}

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if vpr.IsSet(ckey) {
				i := vpr.GetUint64(ckey)
				f.SetUint(i)
			} else {
				tag, ok := vf.Tag.Lookup(dfKey)
				if i, err := strconv.ParseUint(tag, 10, 64); ok && err == nil {
					f.SetUint(i)
				}
			}

		case reflect.Float32, reflect.Float64:
			if vpr.IsSet(ckey) {
				i := vpr.GetFloat64(ckey)
				f.SetFloat(i)
			} else {
				tag, ok := vf.Tag.Lookup(dfKey)
				if i, err := strconv.ParseFloat(tag, 64); ok && err == nil {
					f.SetFloat(i)
				}
			}
		}
	}

	return cls
}
