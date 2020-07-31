package config

import (
	"github.com/tkanos/gonfig"
	"log"
	"os"
)

type Config struct {
	Enable         bool
	ServerPort     string
	DatabaseConfig DatabaseConfig
	EurekaConfig   EurekaConfig
}

type DatabaseConfig struct {
	DatabaseHost     string
	DatabasePort     int
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DdlAuto          string
}

type EurekaConfig struct {
	EurekaUrl      string
	InstanceId     string
	HostName       string
	SecurePort     int
	App            string
	IPAddr         string
	VipAddress     string
	HomePageUrl    string
	StatusPageUrl  string
	HealthCheckUrl string
	EnableEureka   bool
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
