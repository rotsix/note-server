package app

import (
	"log"
	"net/http"
	"server/internal/routes"
	"server/pkg/config"
	"server/pkg/users"

	"github.com/gorilla/mux"
)

// Run is the app main loop
func Run() {
	conf := config.Parse()
	if err := users.Init(conf); err != nil {
		log.Println("users' db init: ", err)
		panic(err)
	}

	r := mux.NewRouter()
	routes.HandleUsers(r.PathPrefix("/users/").Subrouter())

	log.Println("launching server")
	log.Println("  ------------")
	log.Fatalln(http.ListenAndServe(":8080", r))
}
