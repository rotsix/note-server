package routes

import (
	"encoding/json"
	"net/http"
	"server/pkg/errors"
	"server/pkg/notes"

	"github.com/gorilla/mux"
)

// HandleNotes handles routes for /notes prefix
func HandleNotes(r *mux.Router) {
	r.HandleFunc("/new", new).Methods("POST")
	r.HandleFunc("/delete", delete).Methods("POST")
	r.HandleFunc("/modify", modify).Methods("POST")
	r.HandleFunc("/get", get).Methods("POST")
}

func new(rw http.ResponseWriter, r *http.Request) {
	token := GetToken(r)
	uid := token["uid"]
	title := r.Form.Get("title")
	description := r.Form.Get("description")

	if err := notes.New(uid, title, description); err != nil {
		errors.Manage(rw, err)
		return
	}
}

func delete(rw http.ResponseWriter, r *http.Request) {
	token := GetToken(r)
	uid := token["uid"]
	id := r.Form.Get("id")
	if err := notes.Delete(uid, id); err != nil {
		errors.Manage(rw, err)
		return
	}
}

func modify(rw http.ResponseWriter, r *http.Request) {
	token := GetToken(r)
	uid := token["uid"]
	id := r.Form.Get("id")
	title := r.Form.Get("title")
	description := r.Form.Get("description")

	if err := notes.Modify(uid, id, title, description); err != nil {
		errors.Manage(rw, err)
		return
	}
}

func get(rw http.ResponseWriter, r *http.Request) {
	token := GetToken(r)
	uid := token["uid"]
	id := r.Form.Get("id")

	note, err := notes.Get(uid, id)
	if err != nil {
		errors.Manage(rw, err)
		return
	}

	noteJSON, _ := json.Marshal(note)
	WriteJSON(rw, string(noteJSON))
}
