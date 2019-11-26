package users

import (
	"database/sql"
	"log"

	"server/pkg/config"
	"server/pkg/errors"
)

// Login returns a token, or an error in case of fail
func Login(username, password string) (string, error) {
	if username == "" || password == "" {
		return "", new(errors.BadRequest)
	}

	var id int
	row := config.Db["users"].QueryRow("SELECT id FROM informations WHERE (username=$1 AND password=$2)", username, password)
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

	log.Println("pkg/users/users.go:Login: # TODO generate token")
	log.Println("pkg/users/users.go:Login: # TODO insert created session into db")
	/*
		lastSeen := "1970-01-01"
		expiration := "1970-01-01"

		if _, err := config.Db["users"].Exec(
			"INSERT INTO sessions(uid, token, last_seen, expiration) VALUES ($1, $2, $3, $4)",
			id, tok, lastSeen, expiration,
		); err != nil {
			log.Printf("pkg/users/users.go:Login: cannot create session: %s", err)
			return "", new(errors.Internal)
		}
	*/

	tok := "t0k3n"
	return tok, nil
}

// Logout removes [token] from logged users
func Logout(tok string) error {
	if tok == "" {
		return new(errors.NotFound)
	}

	// NOTE remove from sessions
	log.Printf("pkg/users/users.go:Logout: # TODO")

	return nil
}
