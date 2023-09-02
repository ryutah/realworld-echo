package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

const envPrefix = "realworld"

type Config struct {
	Port              string `envconfig:"PORT" default:"8080"`
	ProjectID         string `envconfig:"GOOGLE_CLOUD_PROJECT" default:""`
	RequestTimeOutSec int    `default:"60"`
	Service           string `default:"realworld"`
	Version           string `default:"v1.0.0"`
	DBConnection      string `envconfig:"DB_CONNECTION"`
}

var _config Config

func init() {
	if err := envconfig.Process(envPrefix, &_config); err != nil {
		panic(err)
	}
}

func GetConfig() Config {
	return _config
}

func (c Config) RequestTimeOut() time.Duration {
	return time.Duration(c.RequestTimeOutSec) * time.Second
}
