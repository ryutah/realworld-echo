package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port      string `envconfig:"PORT" default:"8080"`
	ProjectID string `envconfig:"GOOGLE_CLOUD_PROJECT" default:""`
}

var _config Config

func init() {
	if err := envconfig.Process("myapp", &_config); err != nil {
		panic(err)
	}
}

func GetConfig() Config {
	return _config
}
