package routes

import (
	"log"
	"net/http"
	"note-server/pkg/token"
)

// WriteJSON sets [rw] header's content-type to JSON, and sends [msg] (a JSON)
func WriteJSON(rw http.ResponseWriter, msg string) {
	rw.Header().Set("Content-Type", "application/json")
	m := []byte(msg)
	rw.Write(m)
}

// ParseForm simply.. parses [r] form
func ParseForm(r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("routes: parsing form: ", err)
	}
	defer r.Body.Close()
}

// GetToken from request
func GetToken(r *http.Request) map[string]string {
	tokRaw, err := r.Cookie("token")
	if err != nil {
		return nil
	}

	log.Println("internal/routes/utils.go:GetToken: # TODO parse token")
	tok := token.Parse(tokRaw)

	return tok
}
