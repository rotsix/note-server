package app

import (
	"log"
	"net/http"
	"server/internal/routes"
	"server/pkg/config"

	"github.com/gorilla/mux"
)

// Run is the app main loop
func Run() {
	conf := config.Parse()
	log.Println(config.ToString(*conf))

	r := mux.NewRouter()
	routes.HandleUsers(r.PathPrefix("/users/").Subrouter())

	log.Println("launching server")
	log.Println("  ------------")
	log.Fatalln(http.ListenAndServe(":8080", r))
}
