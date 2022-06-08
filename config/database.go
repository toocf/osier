package config

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Dbname   string `mapstructure:"dbname"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Charset  string `mapstructure:"charset"`
	MaxIdle  int    `mapstructure:"maxidle"`
	MaxOpen  int    `mapstructure:"maxopen"`
}
