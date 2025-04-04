package config

type Configuration struct {
	ServerPort          int
	DBMaxConnection     int
	DBMaxIdleConnection int
	DBPath              string
	EventServer         string
	EventServerPort     int
}

func LoadConfiguration() (*Configuration, error) {

	cfg := &Configuration{
		ServerPort:          8080,
		DBMaxConnection:     10,
		DBMaxIdleConnection: 5,
		DBPath:              "pajarit.db",
		EventServer:         "nats://localhost",
		EventServerPort:     4222,
	}

	// TODO - por la simpleza de la configuración
	// tal vez no sea necesario devolver error
	return cfg, nil
}
