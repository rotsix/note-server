package routes

import (
	"encoding/json"
	"net/http"
	"note-server/pkg/errors"
	"note-server/pkg/notes"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// HandleNotes handles routes for /notes prefix
func HandleNotes(r *mux.Router) {
	r.HandleFunc("/new", new).Methods("POST")
	r.HandleFunc("/delete/{id:[0-9]+}", delete).Methods("POST")
	r.HandleFunc("/modify/{id:[0-9]+}", modify).Methods("PUT")
	r.HandleFunc("/get/{id:[0-9]+}", get).Methods("GET")
	r.HandleFunc("/all", all).Methods("GET")
}

func getID(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["id"]
}

func new(rw http.ResponseWriter, r *http.Request) {
	token := GetToken(r)
	uid := token["uid"]
	ParseForm(r)
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
	id := getID(r)
	if err := notes.Delete(uid, id); err != nil {
		errors.Manage(rw, err)
		return
	}
}

func modify(rw http.ResponseWriter, r *http.Request) {
	token := GetToken(r)
	uid := token["uid"]
	ParseForm(r)
	id := getID(r)
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
	id := getID(r)

	note, err := notes.Get(uid, id)
	if err != nil {
		errors.Manage(rw, err)
		return
	}

	noteJSON, _ := json.Marshal(note)
	WriteJSON(rw, string(noteJSON))
}

func all(rw http.ResponseWriter, r *http.Request) {
	token := GetToken(r)
	uid := token["uid"]

	notesInt, err := notes.All(uid)
	if err != nil {
		errors.Manage(rw, err)
		return
	}

	// [1, 2, 3] -> ["1", "2", "3"]
	notesStr := make([]string, len(notesInt))
	for i, note := range notesInt {
		notesStr[i] = strconv.Itoa(note)
	}

	// "1,2,3"
	rw.Write([]byte(strings.Join(notesStr, ",")))
}
