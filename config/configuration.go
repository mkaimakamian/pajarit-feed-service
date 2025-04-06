package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Server struct {
	Port int `yaml:"port"`
}

type Database struct {
	Path              string `yaml:"path"`
	MaxConnection     int    `yaml:"maxConnection"`
	MaxIdleConnection int    `yaml:"maxIdleConnection"`
}

type Event struct {
	Server string `yaml:"serverUrl"`
	Port   int    `yaml:"port"`
}

type Configuration struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Event    Event    `yaml:"event"`
}

func LoadConfiguration() (*Configuration, error) {
	cfg, err := loadYaml()
	if err != nil {
		return nil, err
	}

	// TODO - Esto podría encapsularse y generalizarse, de modo
	// que los parámetros de configuración puedan ser sobreescritos
	// por las variables de entorno del desarrollador
	if value := os.Getenv("EVENT_SERVER_URL"); value != "" {
		cfg.Event.Server = value
	}

	if value := os.Getenv("DB_PATH"); value != "" {
		cfg.Database.Path = value
	}

	return cfg, nil
}

func loadYaml() (*Configuration, error) {
	cfg := &Configuration{}

	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
