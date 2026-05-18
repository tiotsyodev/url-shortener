package core_http_server

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host     string        `envconfig:"HOST" default:"localhost"`
	Port        int           `envconfig:"PORT" required:"true"`
	Timeout     time.Duration `envconfig:"TIMEOUT" required:"true"`
	IdleTimeout time.Duration `envconfig:"IDLE_TIMEOUT" required:"true"`
}

func NewCofig() (Config, error) {
	var config Config
	err := envconfig.Process("HTTP", &config)
	if err != nil {
		return Config{}, fmt.Errorf("process http-server config: %w", err)
	}

	return config, nil
}

func ConfigMustLoad() Config {
	config, err := NewCofig()
	if err != nil {
		err = fmt.Errorf("get http server config: %w" , err)
		panic(err)
	}

	return config
}