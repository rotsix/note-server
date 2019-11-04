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

func superExec(dbName, query string) (sql.Result, error) {
	dbConf := config.Config.Databases[dbName]
	con := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConf.Host,
		dbConf.Port,
		dbConf.User,
		dbConf.Pass,
		dbConf.User,
	)

	var db *sql.DB
	var err error
	if db, err = sql.Open(dbConf.Driver, con); err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db.Exec(query)
}

// Create db:table (all if empty)
func Create(db, table string) error {
	if db == "" {
		for db := range config.Config.Databases {
			if err := createDb(db); err != nil {
				return err
			}
			for table := range config.Config.Databases[db].Tables {
				if err := createTable(db, table); err != nil {
					return err
				}
			}
		}
	} else {
		if err := createDb(db); err != nil {
			return err
		}
		if table == "" {
			for table := range config.Config.Databases[db].Tables {
				if err := createTable(db, table); err != nil {
					return err
				}
			}
		} else {
			if err := createTable(db, table); err != nil {
				return err
			}
		}
	}
	return nil
}

func createDb(dbName string) error {
	query := fmt.Sprintf("CREATE DATABASE %s", dbName)
	if _, err := superExec(dbName, query); err != nil {
		return fmt.Errorf("cannot create database '%s': %s", dbName, err)
	}
	log.Printf("created database '%s'", dbName)
	if err := config.InitDb(dbName); err != nil {
		return fmt.Errorf("db init: %s", err)
	}
	return nil
}

func createTable(db, table string) error {
	var query strings.Builder
	query.WriteString(fmt.Sprintf("CREATE TABLE %s (", table))
	i := 0
	for field := range config.Config.Databases[db].Tables[table].Fields {
		f := config.Config.Databases[db].Tables[table].Fields[field]
		qry := strings.Join([]string{field, f.Type, strings.Join(f.Constraints, " ")}, " ")
		if i == len(config.Config.Databases[db].Tables[table].Fields)-1 {
			query.WriteString(qry)
			continue
		}
		query.WriteString(qry + ", ")
		i++
	}
	query.WriteString(")")

	if _, err := config.Db[db].Exec(query.String()); err != nil {
		return fmt.Errorf("cannot create table '%s' in '%s': %s", table, db, err)
	}
	log.Printf("created table '%s' in '%s'", table, db)
	return nil
}

// Drop db:table (all if empty)
func Drop(db, table string) error {
	if db == "" {
		for db := range config.Config.Databases {
			dropDb(db, config.Config.Databases[db])
		}
	} else {
		if table != "" {
			if err := dropTable(db, table, config.Config.Databases[db]); err != nil {
				return err
			}
		}
		dropDb(db, config.Config.Databases[db])
	}
	return nil
}

func dropDb(db string, conf config.DbType) {
	query := fmt.Sprintf("DROP DATABASE %s", db)
	if _, err := superExec(db, query); err != nil {
		log.Printf("cannot drop database '%s': %s", db, err)
		return
	}
	log.Printf("dropped database '%s'", db)
}

func dropTable(db, table string, conf config.DbType) error {
	query := fmt.Sprintf("DROP TABLE %s", table)
	if _, err := config.Db[db].Exec(query); err != nil {
		return fmt.Errorf("cannot drop table '%s' in '%s': %s", table, db, err)
	}
	log.Printf("dropped table '%s'", table)
	return nil
}

// Fill db:table with mock data (all if empty)
func Fill(db, table string) error {
	if db == "" {
		for db := range config.Config.Databases {
			if err := fillDb(db, table, config.Config.Databases[db]); err != nil {
				return err
			}
		}
	} else {
		if err := fillDb(db, table, config.Config.Databases[db]); err != nil {
			return err
		}
	}
	return nil
}

func fillDb(db, table string, conf config.DbType) error {
	if table == "" {
		for table := range conf.Tables {
			if err := fillDbTable(db, table, conf.Tables[table]); err != nil {
				return fmt.Errorf("cannot fill table '%s' in '%s': %s", table, db, err)
			}
		}
	} else {
		if err := fillDbTable(db, table, conf.Tables[table]); err != nil {
			return fmt.Errorf("cannot fill table '%s' in '%s': %s", table, db, err)
		}
	}
	return nil
}

func fillDbTable(db, table string, conf config.TableType) error {
	for user := range conf.MockData {
		var query, bk, bv strings.Builder

		userConf := conf.MockData[user]
		query.WriteString("INSERT INTO " + table + "(")
		for k, v := range userConf {
			bk.WriteString(k + ",")
			bv.WriteString("'" + v + "',")
		}
		query.WriteString(bk.String()[:len(bk.String())-1])
		query.WriteString(") VALUES (")
		query.WriteString(bv.String()[:len(bv.String())-1])
		query.WriteString(")")

		if _, err := config.Db[db].Exec(query.String()); err != nil {
			log.Printf("cannot insert into '%s/%s': %s", db, table, err)
		}
	}

	log.Printf("filled table '%s' in '%s' (mock data)", table, db)
	return nil
}
