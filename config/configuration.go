package config

type Configuration struct {
	ServerPort          int
	DBMaxConnection     int
	DBMaxIdleConnection int
	DBPath              string
}

func LoadConfiguration() (*Configuration, error) {

	cfg := &Configuration{
		ServerPort:          8080,
		DBMaxConnection:     10,
		DBMaxIdleConnection: 5,
		DBPath:              "pajarit.db",
	}

	// TODO - por la simpleza de la configuración
	// tal vez no sea necesario devolver error
	return cfg, nil
}
