package config

type App struct {
	// 运行模式
	Debug string `mapstructure:"debug"`
	// 语言
	Lang string `mapstructure:"lang"`
	// 服务器
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
