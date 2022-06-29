package config

type Swag struct {
	Title       string `mapstructure:"title"`
	Description string `mapstructure:"description"`
	Version     string `mapstructure:"version"`
	BasePath    string `mapstructure:"basepath"`
}
