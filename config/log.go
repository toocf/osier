package config

type Log struct {
	Level string `mapstructure:"level"`
	Path  string `mapstructure:"path"`
}
