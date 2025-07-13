package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string
}

//env-default: "production", use this in Env, this help us to directly deployed on production with no failure and maintain the logs
type Config struct{
	Env string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storgae_path" env-required:"true"`
	HTTPServer `yaml:"http_server"`
}


func MustLoad() *Config {
	var configPath string;

	configPath = os.Getenv("CONFIG_PATH");

	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file");
		flag.Parse();

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config Path is not set");
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath);
	}


	var cfg Config;

	err := cleanenv.ReadConfig(configPath, &cfg);

	if err != nil {
		log.Fatalf("Can not read config file: %s", err.Error())
	}

	return &cfg;
}