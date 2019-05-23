package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Config describes an app configuration
type Config struct {
	Databases map[string]DbType `json:"databases"`
}

// DbType represents a database in config
type DbType struct {
	Host        string               `json:"host"`
	Port        int                  `json:"port"`
	Driver      string               `json:"driver"`
	User        string               `json:"user"`
	Pass        string               `json:"pass"`
	Description string               `json:"desc"`
	Tables      map[string]TableType `json:"tables"`
}

// TableType represents a database's tables in config
type TableType struct {
	Description string               `json:"desc"`
	Fields      map[string]FieldType `json:"fields"`
	MockData    mockType             `json:"mock"`
}

// map which associates string with a map of string -> string
type mockType = map[string]map[string]string

// FieldType represents a table's field in config
type FieldType struct {
	Type        string   `json:"type"`
	Constraints []string `json:"constraints"`
}

func getenv(env string) string {
	res := os.Getenv(env)
	if res == "" {
		log.Printf("couldn't read %s from env", env)
	}
	return res
}

// Parse parses args, config file, and returns a Config
func Parse() *Config {
	log.Println("pkg/config/config.go:Parse: # TODO: get config location properly")
	confLocation := os.Getenv("GOPATH")
	if confLocation == "" {
		log.Panicln("couldn't read GOPATH in env")
	}

	confLocation += "/src/server/configs/config.json"
	file, err := ioutil.ReadFile(confLocation)
	if err != nil {
		log.Println("error while reading your config file")
		panic(err)
	}

	config := new(Config)
	if err = json.Unmarshal(file, config); err != nil {
		panic(err)
	}

	return config
}
