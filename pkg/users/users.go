package users

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"server/pkg/config"

	_ "github.com/lib/pq"
)

// Config stores current app configuration
var Config *config.Config

// Db is a database connection
var Db *sql.DB

// Init starts connection to SQL
func Init(conf *config.Config) {
	Config = conf

	usersDb := Config.Databases["users"]

	src := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		usersDb.Host,
		usersDb.Port,
		usersDb.User,
		usersDb.Pass,
		usersDb.Name)

	Db, _ = sql.Open(usersDb.Driver, src)
	if err := Db.Ping(); err != nil {
		panic(err)
	}
	log.Println("connected to database")
}

// Login returns a token, or an error in case of fail
func Login(user, password string) (string, error) {
	if user == "" || password == "" {
		return "", errors.New("user not found")
	}

	log.Println("pkg/users/users.go:Login: # TODO")
	return "t0k3n", nil
}

// Logout removes [token] from logged users
func Logout(token string) {
	if token == "" {
		return
	}

	log.Println("pkg/users/users.go:Logout: # TODO")
}
