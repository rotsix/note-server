package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Config describes an app configuration
type Config struct {
	Location  string            `json:"location"`
	Prod      bool              `json:"prod"`
	Databases map[string]DbType `json:"databases"`
}

type DbType struct {
	Host   string `json:"host"`
	Port   int    `json:"port"`
	Driver string `json:"driver"`
	Name   string `json:"name"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
}

// Parse parses args, config file, and returns a Config
func Parse() *Config {
	config := new(Config)

	log.Println("/pkg/config/config.go:Parse: # TODO fix how to get config location")
	config.Location = os.Getenv("GOPATH") + "/src/server/configs/config.json"
	file, err := ioutil.ReadFile(config.Location)
	if err != nil {
		log.Println("you forgot to create your config file")
		panic(err)
	}

	if err := json.Unmarshal(file, config); err != nil {
		panic(err)
	}

	return config
}
