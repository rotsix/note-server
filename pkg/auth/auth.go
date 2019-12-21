package auth

import (
	"database/sql"
	"log"

	"server/pkg/config"
	"server/pkg/errors"
	"server/pkg/token"
)

// Login returns a token, or an error in case of fail
func Login(username, password string) (string, error) {
	if username == "" || password == "" {
		return "", new(errors.BadRequest)
	}

	var id int
	row := config.Db["note"].QueryRow("SELECT id FROM accounts WHERE (username=$1 AND password=$2)", username, password)
	err := row.Scan(&id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			// wrong username/password
			log.Printf("failed to login: %s:%s", username, password)
			return "", new(errors.Unauthorized)
		default:
			log.Printf("pkg/auth/auth.go:Login: error during results' scan: %s", err)
			return "", new(errors.Internal)
		}
	}

	var tok string
	tok, err = token.New()
	if err != nil {
		return "", new(errors.Internal)
	}

	tok = "t0k3n"
	log.Println("pkg/auth/auth.go:Login: # TODO insert created session into db")
	/*
		lastSeen := "1970-01-01"
		expiration := "1970-01-01"

		if _, err := config.Db["note"].Exec(
			"INSERT INTO sessions(uid, token, last_seen, expiration) VALUES ($1, $2, $3, $4)",
			id, tok, lastSeen, expiration,
		); err != nil {
			log.Printf("pkg/auth/auth.go:Login: cannot create session: %s", err)
			return "", new(errors.Internal)
		}
	*/

	return tok, nil
}

// Logout removes [token] from logged users
func Logout(uidStr string) error {
	log.Printf("pkg/auth/auth.go:Logout: # TODO")

	if uidStr == "" {
		return nil
	}

	// NOTE remove from db[note].sessions
	token.Parse()

	return nil
}
