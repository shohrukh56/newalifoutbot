package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
)

type TelegramConfigType struct {
	Token      string `yaml:"token"`
	WebHookURL string `yaml:"web-hook-url"`
}

type DBConfigType struct {
	DriverName       string `yaml:"DriverName" env:"DRIVER_NAME" env-default:"postgres"`
	DataBaseUser     string `yaml:"DataBaseUser" env:"DB_USER" env-default:"postgres"`
	DataBaseName     string `yaml:"DataBaseName" env:"DB_NAME" env-default:"postgres"`
	DataBasePassword string `yaml:"DataBasePassword" env:"DB_PASSWORD" env-default:"pass"`
	SSLMode          string `yaml:"SSLMode" env:"SSL_MODE" env-default:"disable"`
	Port             string `yaml:"Port" env:"PORT" env-default:"5433"`
	Host             string `yaml:"Host" env:"HOST" env-default:"localhost"`
	URI              string`yaml:"URI" env-default:""`
	HerokuCLI   	 string`yaml:"HerokuCLI"`
}


//ReadConfig parse configuration file according to YAML format (yaml tag in this case);
//reads environment variables and overwrites values from the file with the values which was found in the environment (env tag);
//if no value was found on the first two steps, the field will be filled with the default value (env-default tag) if it is set.

func TGConfig() TelegramConfigType {
	var config TelegramConfigType

	err := cleanenv.ReadConfig("configs/configs.yaml", &config)
	if err != nil {
		log.Panic(err)
	}

	return config
}

func DBConfig() string {
	var config DBConfigType

	err := cleanenv.ReadConfig("configs/configs.example.yaml", &config)
	if err != nil {
		log.Panic(err)
	}

	dbUrl := "host=" + config.Host + " user=" + config.DataBaseUser + " password=" + config.DataBasePassword + " database=" + config.DataBaseName + " sslmode=" + config.SSLMode + " port=" + config.Port + "URI="+config.URI+"HerokuCLI="+config.HerokuCLI

	return dbUrl
}
