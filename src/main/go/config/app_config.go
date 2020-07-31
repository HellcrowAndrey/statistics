package config

import (
	"github.com/tkanos/gonfig"
	"log"
	"os"
)

type Config struct {
	Enable       bool
	ServerPort   string
	Host         string
	Port         int
	User         string
	Password     string
	Dbname       string
	EnableEureka bool
	DdlAuto      string
}

func NewConfig() *Config {
	config := Config{}
	var err interface{}
	profile := os.Getenv("ENVIRONMENT")
	switch profile {
	case "dev":
		err = gonfig.GetConf("src/main/resources/config.development.json", &config)
	case "prod":
		err = gonfig.GetConf("src/main/resources/config.production.json", &config)
	case "test":
		err = gonfig.GetConf("src/main/resources/config.production.json", &config)
	case "default":
		err = gonfig.GetConf("src/main/resources/config.default.json", &config)
	default:
		err = gonfig.GetConf("src/main/resources/config.default.json", &config)
	}
	log.Println("Profile name: ", profile, "Profile parameters: ", config)
	if err != nil {
		panic(err)
	}
	return &config
}
