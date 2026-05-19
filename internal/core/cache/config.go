package core_cache

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port		string		  `envconfig:"PORT" required:"true"`
	Host 		string		  `envconfig:"HOST" required:"true"`
	Password    string        `envconfig:"PASSWORD" required:"true"`
	User        string        `envconfig:"USER" required:"true"`
	DB          int           `envconfig:"DB" required:"true" default:"0"`
	MaxRetries  int           `envconfig:"MAX_RETRIES" default:"5"`
	DialTimeout time.Duration `envconfig:"DIAL_TIMEOUT" default:"10s"`
	Timeout     time.Duration `envconfig:"TIMEOUT" default:"5s"`
}

func NewConfig() (Config, error) {
	var cfg Config

	if err := envconfig.Process("REDIS", &cfg); err != nil {
		return Config{}, fmt.Errorf("process cache cfg: %w", err)
	}

	return cfg, nil
}

func NewConfigMust() (Config) {
	cfg, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("load repository cfg: %w", err)
		panic(err)
	}

	return cfg
}