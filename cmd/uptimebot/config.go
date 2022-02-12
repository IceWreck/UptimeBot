package main

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Interval int       `yaml:"interval"`
	Services []Service `yaml:"services"`
}

type Service struct {
	Name            string            `yaml:"name"`
	URL             string            `yaml:"url"`
	RespCode        int               `yaml:"resp_code"`
	Alert           bool              `yaml:"alert"`
	AlertWebHook    string            `yaml:"alert_webhook"`
	AlertFormParams map[string]string `yaml:"alert_form_parms"`
}

func loadConfig(configFileLoc string) Config {
	var settings Config
	yamlFile, err := os.Open(configFileLoc)
	if err != nil {
		log.Fatalln("Config file not found")
	}
	defer yamlFile.Close()
	log.Println("Loaded config.yaml")
	byteValue, _ := io.ReadAll(yamlFile)
	err = yaml.Unmarshal(byteValue, &settings)
	if err != nil {
		log.Fatalln("Error parsing config file", err)
	}

	// Set default values.

	if settings.Interval < 10 {
		settings.Interval = 300
		log.Println("Interval cannot be less than 10 seconds, using default value of 300 seconds")
	}

	for index, service := range settings.Services {
		if service.RespCode == 0 {
			settings.Services[index].RespCode = 200
		}
	}

	return settings

}
