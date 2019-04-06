package routes

import (
	"fmt"
	"log"
	"net/http"
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
		panic(err)
	}
	defer r.Body.Close()

	user := r.Form.Get("user")
	password := r.Form.Get("password")

	token, err := users.Login(user, password)
	if err != nil {
		http.NotFound(rw, r)
		return
	}

	res := fmt.Sprintf(`{"token":"%s"}`, token)
	WriteJSON(rw, res)
}

func logout(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	defer r.Body.Close()

	token := r.Form.Get("token")
	users.Logout(token)
	Success(rw)
}

func signIn(rw http.ResponseWriter, r *http.Request) {
	log.Println("internal/routes/users.go:signIn: # TODO")
}
