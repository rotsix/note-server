package auth

import (
	"database/sql"
	"log"
	"strconv"

	"note-server/pkg/config"
	"note-server/pkg/errors"
	"note-server/pkg/token"
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
	tok, err = token.New(id)
	if err != nil {
		return "", new(errors.Internal)
	}

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
	return nil
}

// SignIn creates a new user in database
func SignIn(username, password string) error {
	if username == "" || password == "" {
		return new(errors.BadRequest)
	}

	log.Println("pkg/auth/auth.go:SignIn: # TODO")

	return nil
}

// DeleteAccount and associated notes
func DeleteAccount(uidStr string) error {
	if _, err := strconv.Atoi(uidStr); err != nil {
		return new(errors.BadRequest)
	}

	log.Println("pkg/auth/auth.go:DeleteAccount: # TODO")
	whatisthat, err := config.Db["note"].Exec("")
	_ = whatisthat

	if err != nil {
		switch err.(type) {
		default:
			return new(errors.Internal)
		}
	}

	return nil
}
