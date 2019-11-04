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
	if err := config.Parse(); err != nil {
		log.Println("config parsing: ", err)
		panic(err)
	}

	for dbName := range config.Config.Databases {
		if err := config.InitDb(dbName); err != nil {
			log.Printf("db init '%s': %s", dbName, err)
			panic(err)

		}
	}

	r := mux.NewRouter()
	routes.HandleUsers(r.PathPrefix("/users/").Subrouter())

	log.Println("launching server")
	log.Println("  ------------")
	log.Fatalln(http.ListenAndServe(":8080", r))
}
