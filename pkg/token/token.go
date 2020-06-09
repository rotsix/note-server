package token

import (
	"log"
	"net/http"
)

// New generates a new JWT token
func New(id int) (string, error) {
	log.Println("pkg/token/token.go:New: # TODO")
	return "t0k3n", nil

	/*
		claims := &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
		}

		tok := jwt.NewWithClaims(nil, claims)
		res, err := tok.SignedString(nil)
		if err != nil {
			return "", err
		}
		return res, nil
	*/
}

// Parse a JWT token
func Parse(jwtRaw *http.Cookie) map[string]string {
	log.Println("pkg/token/token.go:Parse: # TODO")
	res := map[string]string{"uid": "1"}
	return res
}
