package users

import (
	"database/sql"
	"fmt"
	"log"

	"server/pkg/config"
	"server/pkg/errors"

	// postgres driver
	_ "github.com/lib/pq"
)

// Config stores app configuration
var Config *config.Config

var db *sql.DB

// Init starts connection to SQL
func Init(conf *config.Config) error {
	Config = conf

	dbConf := Config.Databases["users"]
	usersCon := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConf.Host,
		dbConf.Port,
		dbConf.User,
		dbConf.Pass,
		"users",
	)

	var err error
	if db, err = sql.Open(dbConf.Driver, usersCon); err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	log.Printf("connected to databases")
	return nil
}

// Login returns a token, or an error in case of fail
func Login(username, password string) (string, error) {
	if username == "" || password == "" {
		return "", new(errors.BadRequest)
	}

	var id int
	row := db.QueryRow("SELECT id FROM informations WHERE (username=$1 AND password=$2)", username, password)
	err := row.Scan(&id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			// wrong username/password
			log.Printf("failed to login: %s:%s", username, password)
			return "", new(errors.Unauthorized)
		default:
			log.Printf("pkg/users/users.go:Login: error during results' scan: %s", err)
			return "", new(errors.Internal)
		}
	}

	// now, we have user's id, let's create a token, and make a new session
	log.Printf("pkg/users/users.go:Login: # TODO generate token")
	token := "t0k3n"
	lastSeen := "1970-01-01"
	expiration := "1970-01-01"

	if _, err := db.Exec(
		"INSERT INTO sessions(uid, token, last_seen, expiration) VALUES ($1, $2, $3, $4)",
		id, token, lastSeen, expiration,
	); err != nil {
		log.Printf("pkg/users/users.go:Login: cannot create session: %s", err)
		return "", new(errors.Internal)
	}

	return token, nil
}

// Logout removes [token] from logged users
func Logout(token string) error {
	if token == "" {
		return new(errors.NotFound)
	}

	log.Printf("pkg/users/users.go:Logout: # TODO")
	return nil
}
