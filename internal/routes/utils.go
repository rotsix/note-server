package routes

import "net/http"

// WriteJSON sets [rw] header's content-type to JSON, and sends [msg] (a JSON)
func WriteJSON(rw http.ResponseWriter, msg string) {
	rw.Header().Set("Content-Type", "application/json")
	m := []byte(msg)
	rw.Write(m)
}
