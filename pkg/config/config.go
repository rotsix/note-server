package config

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	// postgres driver
	_ "github.com/lib/pq"
)

// Config describes an app configuration
var Config *config

// Db stores databases connections
var Db map[string]*sql.DB = map[string]*sql.DB{}

type config struct {
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

// InitDb initialises databases connections
func InitDb(name string) error {
	f := func(dbName string) error {
		dbConf := Config.Databases[dbName]
		con := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			dbConf.Host,
			dbConf.Port,
			dbConf.User,
			dbConf.Pass,
			dbName,
		)

		var db *sql.DB
		var err error
		if db, err = sql.Open(dbConf.Driver, con); err != nil {
			return err
		}
		if err = db.Ping(); err != nil {
			return err
		}
		if db == nil {
			return errors.New("db is nil")
		}
		Db[dbName] = db
		return nil
	}

	if name == "" {
		for dbName := range Config.Databases {
			if err := f(dbName); err != nil {
				return err
			}
			log.Printf("connected to database: %s", dbName)
		}
	} else {
		if err := f(name); err != nil {
			return err
		}
		log.Printf("connected to database: %s", name)
	}
	return nil
}

// Parse parses args, config file, and returns a Config
func Parse() error {
	confLocation := os.Getenv("CONF_LOCATION")
	if confLocation == "" {
		return errors.New("couldn't read CONF_LOCATION in env")
	}

	file, err := ioutil.ReadFile(confLocation)
	if err != nil {
		return errors.New("error while reading your config file: " + err.Error())
	}

	tmp := new(config)
	if err = json.Unmarshal(file, tmp); err != nil {
		return errors.New("error while parsing your config file: " + err.Error())
	}

	Config = tmp
	return nil
}
