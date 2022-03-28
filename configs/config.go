package configs

import "os"

type Config struct {
	Port string
}

func GetEnvs() Config {
	var conf Config
	conf.Port = os.Getenv("PORT")
	return conf
}
