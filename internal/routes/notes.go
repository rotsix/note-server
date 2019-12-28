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
	r.HandleFunc("/delete", delete).Methods("POST")
	r.HandleFunc("/modify", modify).Methods("POST")
	r.HandleFunc("/get", get).Methods("POST")
	r.HandleFunc("/all", all).Methods("POST")
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
	WriteJSON(rw, strings.Join(notesStr, ","))
}
