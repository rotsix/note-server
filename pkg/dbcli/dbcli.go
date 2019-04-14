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

func initDbCon(db string, dbConf config.DbType) error {
	state := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConf.Host,
		dbConf.Port,
		dbConf.User,
		dbConf.Pass,
		db,
	)
	var err error
	if dbCons[db], err = sql.Open(dbConf.Driver, state); err != nil {
		return fmt.Errorf("cannot open db '%s': %s", db, err)
	}
	if err = dbCons[db].Ping(); err != nil {
		return fmt.Errorf("cannot ping db '%s': %s", db, err)
	}

	return nil
}

// Create db:table (all if empty)
func Create(db, table string, conf *config.Config) error {
	if db == "" {
		for db := range conf.Databases {
			if err := createDb(db, conf.Databases[db]); err != nil {
				return err
			}
			for table := range conf.Databases[db].Tables {
				if err := createTable(db, table, conf.Databases[db].Tables[table]); err != nil {
					return err
				}
			}
		}
	} else {
		if table == "" {
			for table := range conf.Databases[db].Tables {
				if err := createTable(db, table, conf.Databases[db].Tables[table]); err != nil {
					return err
				}
			}
		} else {
			if err := createTable(db, table, conf.Databases[db].Tables[table]); err != nil {
				return err
			}
		}
	}
	return nil
}

func createDb(db string, conf config.DbType) error {
	if err := initDbCon(conf.User, conf); err != nil {
		return err
	}
	query := fmt.Sprintf("CREATE DATABASE %s", db)
	if _, err := dbCons[conf.User].Exec(query); err != nil {
		return fmt.Errorf("cannot create database '%s': %s", db, err)
	}
	if err := initDbCon(db, conf); err != nil {
		// set connection to 'db', could be useful when creating tables
		return err
	}
	log.Printf("created database '%s'", db)
	return nil
}

func createTable(db, table string, conf config.TableType) error {
	var query strings.Builder
	query.WriteString(fmt.Sprintf("CREATE TABLE %s (", table))
	i := 0
	for field := range conf.Fields {
		f := conf.Fields[field]
		qry := strings.Join([]string{field, f.Type, strings.Join(f.Constraints, " ")}, " ")
		if i == len(conf.Fields)-1 {
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
	return nil
}

// Drop db:table (all if empty)
func Drop(db, table string, conf *config.Config) error {
	if db == "" {
		for db := range conf.Databases {
			if err := dropDb(db, conf.Databases[db]); err != nil {
				return err
			}
		}
	} else {
		if table != "" {
			if err := dropTable(db, table, conf.Databases[db]); err != nil {
				return err
			}
		}
		if err := dropDb(db, conf.Databases[db]); err != nil {
			return err
		}
	}
	return nil
}

func dropDb(db string, conf config.DbType) error {
	if err := initDbCon(conf.User, conf); err != nil {
		return err
	}
	query := fmt.Sprintf("DROP DATABASE %s", db)
	if _, err := dbCons[conf.User].Exec(query); err != nil {
		return fmt.Errorf("cannot drop database '%s': %s", db, err)
	}
	log.Printf("dropped database '%s'", db)
	return nil
}

func dropTable(db, table string, conf config.DbType) error {
	if err := initDbCon(db, conf); err != nil {
		return err
	}
	query := fmt.Sprintf("DROP TABLE %s", table)
	if _, err := dbCons[db].Exec(query); err != nil {
		return fmt.Errorf("cannot drop table '%s' in '%s': %s", table, db, err)
	}
	log.Printf("dropped table '%s'", table)
	return nil
}

// Fill db:table with mock data (all if empty)
func Fill(db, table string, conf *config.Config) error {
	if db == "" {
		for db := range conf.Databases {
			if err := fillDb(db, table, conf.Databases[db]); err != nil {
				return err
			}
		}
	} else {
		if err := fillDb(db, table, conf.Databases[db]); err != nil {
			return err
		}
	}
	return nil
}

func fillDb(db, table string, conf config.DbType) error {
	if err := initDbCon(db, conf); err != nil {
		return err
	}
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

		if _, err := dbCons[db].Exec(query.String()); err != nil {
			log.Printf("cannot insert into '%s/%s': %s", db, table, err)
		}
	}

	log.Printf("filled table '%s' in '%s' (mock data)", table, db)
	return nil
}
