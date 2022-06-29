package config

type Config struct {
	App      App      `mapstructure:"app"`
	Log      Log      `mapstructure:"log"`
	Database Database `mapstructure:"database"`
	Redis    Redis    `mapstructure:"redis"`
	Swag     Swag     `mapstructure:"swag"`
}
