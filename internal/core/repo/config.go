package core_repo

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host        string        `envconfig:"HOST" required:"true"`
	Port        int        `envconfig:"PORT" required:"true"`
	User        string        `envconfig:"USER" required:"true"`
	Password    string        `envconfig:"PASSWORD" required:"true"`
	DatabaseName string        `envconfig:"DB" required:"true"`
	Timeout     time.Duration `envconfig:"TIMEOUT" default:"30s"`
}

func newConfig() (Config, error) {
	var cfg Config
	err := envconfig.Process("POSTGRES", &cfg)
    if err != nil {
		return Config{}, fmt.Errorf("process repository cfg: %w", err)
	}

	return cfg, nil
}

func NewConfigMustLoad() Config {
	cfg, err := newConfig()
	if err != nil {
		err = fmt.Errorf("load repository cfg: %w", err)
		panic(err)
	}

	return cfg
}