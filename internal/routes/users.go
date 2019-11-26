package routes

import (
	"fmt"
	"log"
	"net/http"
	"server/pkg/errors"
	"server/pkg/users"

	"github.com/gorilla/mux"
)

// HandleUsers handles route for /users prefix
func HandleUsers(r *mux.Router) {
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/logout", logout).Methods("POST")
	r.HandleFunc("/signin", signIn).Methods("POST")
}

func login(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("routes:login: ", err)
	}
	defer r.Body.Close()

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	token, err := users.Login(username, password)
	if err != nil {
		errors.Manage(rw, err)
		return
	}

	res := fmt.Sprintf(`{"token":"%s"}`, token)
	WriteJSON(rw, res)
}

func logout(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("routes:logout: ", err)
	}
	defer r.Body.Close()

	token := r.Form.Get("token")
	if err := users.Logout(token); err != nil {
		errors.Manage(rw, err)
		return
	}
}

func signIn(rw http.ResponseWriter, r *http.Request) {
	log.Println("internal/routes/users.go:signIn: # TODO")
}
