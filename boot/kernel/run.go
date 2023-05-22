package kernel

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"osier/app/middle"
	"osier/boot"
	"osier/config"
	"osier/docs"
	"osier/routes"
	"reflect"
)

// 运行
func Run() {

	// 优先的顺序初始化
	boot.Vpr = initVpr()
	boot.Cfg = initCfg(boot.Vpr)
	boot.Gdb = initGdb(boot.Cfg)
	boot.Rds = initRds(boot.Cfg)
	boot.Log = initLog(boot.Cfg)
	boot.Ege = initEge(boot.Cfg)

	// 依托全局加载
	loadLang()
	loadSwag()
	loadMiddle()
	loadRouter()
	loadFinish()

}

// 初始化Viper
func initVpr() *viper.Viper {

	v := viper.New()
	v.SetConfigFile("./config.ini")
	_ = v.ReadInConfig()

	return v
}

// 初始化配置信息
func initCfg(v *viper.Viper) *config.Config {

	c := config.Config{}
	err := v.Unmarshal(&c)
	if err != nil {
		panic("【配置信息】初始化失败：" + err.Error())
	}

	return &c
}

// 初始化数据库
func initGdb(c *config.Config) *gorm.DB {

	cfg := c.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname,
		cfg.Charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("【数据库】连接失败: " + err.Error())
	}

	sql, _ := db.DB()
	// 最大空闲连接数
	sql.SetMaxIdleConns(cfg.MaxIdle)
	// 最大连接数
	sql.SetMaxOpenConns(cfg.MaxOpen)

	return db
}

// 初始化缓存
func initRds(c *config.Config) *redis.Client {

	if !c.Redis.Enable {
		return nil
	}

	rds := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host + ":" + c.Redis.Port,
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	})

	_, err := rds.Ping(context.Background()).Result()
	if err != nil {
		panic("【Redis】连接失败：" + err.Error())
	}

	return rds
}

// 初始化日志Zap
func initLog(c *config.Config) *zap.Logger {

	core := zapcore.NewCore(getEncoder(), getWrite(c.Log), getLevel(c.Log))

	return zap.New(core, zap.AddCaller())
}

// 初始化路由
func initEge(c *config.Config) *gin.Engine {

	gin.SetMode(c.App.Debug)
	gin.DisableConsoleColor()

	gin.DebugPrintRouteFunc = setRouteLog

	return gin.New()
}

// 加载语言
func loadLang() {

	lang := boot.FirstUpper(boot.Cfg.App.Lang)
	method := reflect.ValueOf(boot.LangText{}).MethodByName(lang)
	if method.IsValid() {
		method.Call(nil)
	}
}

// 加载文档
func loadSwag() {

	// 文档页配置
	docs.SwaggerInfo.Title = boot.Cfg.Swag.Title
	docs.SwaggerInfo.Description = boot.Cfg.Swag.Title
	docs.SwaggerInfo.Version = boot.Cfg.Swag.Version
	docs.SwaggerInfo.Host = "http://" + boot.Cfg.App.Host + ":" + boot.Cfg.App.Port
	docs.SwaggerInfo.BasePath = boot.Cfg.Swag.BasePath

	// 文档地址 /swagger/index.html
	boot.Ege.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// 加载中间件
func loadMiddle() {

	boot.Ege.Use(getLogger(boot.Log), getRecovery(boot.Log))

	mids := []gin.HandlerFunc{}
	rt := reflect.ValueOf(middle.Middle{})
	for i := 0; i < rt.NumMethod(); i++ {

		method, ok := rt.Method(i).Interface().(func(*gin.Context))
		if ok {
			mids = append(mids, method)
		}
	}

	boot.Ege.Use(mids...)
}

// 加载路由
func loadRouter() {

	rt := reflect.ValueOf(routes.Route{})
	for i := 0; i < rt.NumMethod(); i++ {
		rt.Method(i).Call([]reflect.Value{reflect.ValueOf(boot.Ege)})
	}

	boot.Ege.NoRoute(func(c *gin.Context) {
		boot.Res(c, "url not found", "Illegal Access!", 4)
	})
}

// 收尾
func loadFinish() {

	addr := fmt.Sprintf("%s:%s",
		boot.Cfg.App.Host,
		boot.Cfg.App.Port)
	_ = boot.Ege.Run(addr)
}
