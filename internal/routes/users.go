package routes

import (
	"net/http"
	"server/pkg/users"
	"strings"

	"github.com/gorilla/mux"
)

// HandleUsers handles route for /users prefix
func HandleUsers(r *mux.Router) {
	r.HandleFunc("/login", login)
	//r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/logout", logout).Methods("POST")
	r.HandleFunc("/signin", signIn).Methods("POST")
}

func login(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	user := r.Form.Get("user")
	password := r.Form.Get("password")

	token := users.Login(user, password)
	if token == "" {
		http.NotFound(rw, r)
		return
	}

	var buff strings.Builder
	buff.WriteString(`{"token":"`)
	buff.WriteString(token)
	buff.WriteString(`"}`)
	res := []byte(buff.String())
	WriteJSON(rw, res)
}

func logout(rw http.ResponseWriter, r *http.Request) {}

func signIn(rw http.ResponseWriter, r *http.Request) {}
