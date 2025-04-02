package config

type Configuration struct {
	Port int
}

func LoadConfiguration() (*Configuration, error) {

	cfg := &Configuration{
		Port: 8080,
	}
	// TODO -
	return cfg, nil
}
