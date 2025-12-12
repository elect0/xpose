package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

type Config struct {
	App struct {
		Environment string
		Port        string
	}
	DB struct {
		Source string
	}
	Paseto struct {
		Asymmetrical         string
		DurationMinutes int
	}
}

func LoadConfig() (*Config, error) {
	var (
		cfg Config
		k   = koanf.New(".")
	)

	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		return nil, err
	}

	if err := k.Unmarshal("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
