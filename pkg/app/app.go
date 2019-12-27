package app

import (
	"log"
	"note-server/internal/routes"
	"note-server/pkg/config"
	"os"

	"github.com/gorilla/mux"
	"github.com/lucas-clemente/quic-go/http3"
)

// Run is the app main loop
func Run() {
	if err := config.Parse(); err != nil {
		log.Println("config parsing: ", err)
		panic(err)
	}

	if err := config.InitDb(""); err != nil {
		log.Printf("db init: %s", err)
		panic(err)
	}

	r := mux.NewRouter()
	routes.HandleAuth(r.PathPrefix("/auth/").Subrouter())
	routes.HandleNotes(r.PathPrefix("/notes/").Subrouter())

	cert, key := os.Getenv("CERT_LOCATION"), os.Getenv("KEY_LOCATION")
	if cert == "" || key == "" {
		log.Fatalln("couldn't read CERT_LOCATION or KEY_LOCATION in env")
	}

	log.Println("launching server")
	log.Println("  ------------")
	log.Fatalln(http3.ListenAndServe(":8080", cert, key, r))
}
