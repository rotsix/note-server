package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Config describes an app configuration
type Config struct {
	Location string `json:"location"`
	DbUser   string `json:"db_user"`
	DbPass   string `json:"db_pass"`
}

// Parse parses args, config file, and returns a Config
func Parse() *Config {
	config := new(Config)

	log.Println("/pkg/config/config.go:Parse: fix how to get config location")
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

// ToString converts a Config into a readable string
func ToString(conf Config) string {
	var buff bytes.Buffer
	buff.WriteString("\ncurrent config:")

	buff.WriteString("\n  location: ")
	buff.WriteString(conf.Location)
	buff.WriteString("\n  db_user:  ")
	buff.WriteString(conf.DbUser)

	return buff.String()
}
