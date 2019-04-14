package dbcli

import (
	"database/sql"
	"fmt"
	"log"
	"server/pkg/config"
	"strings"

	// postgres driver
	_ "github.com/lib/pq"
)

var dbCons = map[string]*sql.DB{}

func initDbCon(name string, dbConf config.DbType) error {
	state := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConf.Host,
		dbConf.Port,
		dbConf.User,
		dbConf.Pass,
		name,
	)
	var err error
	if dbCons[name], err = sql.Open(dbConf.Driver, state); err != nil {
		return fmt.Errorf("cannot open db '%s': %s", name, err)
	}
	if err = dbCons[name].Ping(); err != nil {
		return fmt.Errorf("cannot ping db '%s': %s", name, err)
	}

	log.Printf("connected to database '%s'", name)
	return nil
}

// Create given database, or all if [name] is empty
func Create(name string, conf *config.Config) error {
	if name == "" || name == "all" {
		for db := range conf.Databases {
			if err := CreateDb(db, conf.Databases[db]); err != nil {
				return err
			}
		}
	} else {
		if err := CreateDb(name, conf.Databases[name]); err != nil {
			return err
		}
	}
	return nil
}

// CreateDb with given [name], using [conf]
func CreateDb(name string, conf config.DbType) error {
	// using user's default database to create our one
	if err := initDbCon(conf.User, conf); err != nil {
		return err
	}

	query := fmt.Sprintf("CREATE DATABASE %s", name)
	if _, err := dbCons[conf.User].Exec(query); err != nil {
		return fmt.Errorf("cannot create database '%s': %s", name, err)
	}
	log.Printf("created database '%s'", name)

	// set connection using [name] database
	if err := initDbCon(name, conf); err != nil {
		return err
	}

	if err := CreateTables(name, conf.Tables); err != nil {
		return err
	}
	return nil
}

// CreateTables for [db] database, using a map of configs
func CreateTables(db string, tables map[string]config.TableType) error {
	for table := range tables {
		var query strings.Builder
		query.WriteString(fmt.Sprintf("CREATE TABLE %s (", table))
		i := 0
		for field := range tables[table].Fields {
			f := tables[table].Fields[field]
			qry := strings.Join([]string{field, f.Type, strings.Join(f.Constraints, " ")}, " ")
			if i == len(tables[table].Fields)-1 {
				query.WriteString(qry)
				continue
			}
			query.WriteString(qry + ", ")
			i++
		}
		query.WriteString(")")

		if _, err := dbCons[db].Exec(query.String()); err != nil {
			return fmt.Errorf("cannot create table '%s' in '%s': %s", table, db, err)
		}
		log.Printf("created table '%s' in '%s'", table, db)
	}
	return nil
}

// Drop given database, or all if [name] is empty
func Drop(name string, conf *config.Config) error {
	if name == "" || name == "all" {
		for db := range conf.Databases {
			if err := DropDb(db, conf.Databases[db]); err != nil {
				return err
			}
		}
	} else {
		if err := DropDb(name, conf.Databases[name]); err != nil {
			return err
		}
	}
	return nil
}

// DropDb with given [name], using [conf]
func DropDb(name string, conf config.DbType) error {
	if err := initDbCon(conf.User, conf); err != nil {
		return fmt.Errorf("using '%s' settings: %s", name, err)
	}
	query := fmt.Sprintf("DROP DATABASE %s", name)
	if _, err := dbCons[conf.User].Exec(query); err != nil {
		log.Printf("cannot drop database '%s': %s", name, err)
		return nil
	}
	log.Printf("dropped database '%s'", name)
	return nil
}

// Fill given database, or all if [name] is empty
func Fill(name string, conf *config.Config) error {
	if name == "" || name == "all" {
		for db := range conf.Databases {
			if err := FillDb(db, conf.Databases[db]); err != nil {
				return err
			}
		}
	} else {
		if err := FillDb(name, conf.Databases[name]); err != nil {
			return err
		}
	}
	return nil
}

// FillDb with given [name], using [conf]
func FillDb(name string, conf config.DbType) error {
	switch name {
	case "users":
		if err := fillUsers(conf); err != nil {
			return err
		}
	default:
		log.Printf("no mock data for '%s' db", name)
	}
	return nil
}

func fillUsers(conf config.DbType) error {
	return nil
}
