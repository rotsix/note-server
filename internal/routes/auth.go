package routes

import (
	"fmt"
	"log"
	"net/http"
	"server/pkg/auth"
	"server/pkg/errors"

	"github.com/gorilla/mux"
)

// HandleAuth handles routes for /auth prefix
func HandleAuth(r *mux.Router) {
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/logout", logout).Methods("POST")
	r.HandleFunc("/signin", signIn).Methods("POST")
}

func login(rw http.ResponseWriter, r *http.Request) {
	ParseForm(r)
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	token, err := auth.Login(username, password)
	if err != nil {
		errors.Manage(rw, err)
		return
	}

	res := fmt.Sprintf(`{"token":"%s"}`, token)
	WriteJSON(rw, res)
}

func logout(rw http.ResponseWriter, r *http.Request) {
	token := GetToken(r)
	uid := token["uid"]
	if err := auth.Logout(uid); err != nil {
		errors.Manage(rw, err)
		return
	}
}

func signIn(rw http.ResponseWriter, r *http.Request) {
	log.Println("internal/routes/auth.go:signIn: # TODO")
}
