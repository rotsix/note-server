package notes

import (
	"database/sql"
	"log"
	"note-server/pkg/config"
	"note-server/pkg/errors"
	"strconv"
	"time"
)

// does user with [uid] exists?
func exists(uid int) bool {
	var id int
	config.Db["note"].QueryRow("SELECT id FROM accounts WHERE id=$1", uid).Scan(&id)
	return id != 0
}

// gives current date
func currentDate() string {
	log.Println("pkg/notes/notes.go:currentDate: # TODO why this date format and not another one?")
	return time.Now().Format("2006-01-02 15:04:05")
}

// New creates a new note into database
func New(uidStr, title, description string) error {
	uid, err := strconv.Atoi(uidStr)
	if title == "" || err != nil || !exists(uid) {
		return new(errors.BadRequest)
	}

	query := `INSERT INTO items (uid, creation_date, edition_date, title, description) VALUES ($1, $2, $3, $4, $5)`
	if _, err := config.Db["note"].Exec(query, uid, currentDate(), currentDate(), title, description); err != nil {
		return new(errors.Internal)
	}

	return nil
}

// Delete given note
func Delete(uidStr, idStr string) error {
	_, err := strconv.Atoi(uidStr)
	_, err2 := strconv.Atoi(uidStr)
	if err != nil || err2 != nil {
		return new(errors.BadRequest)
	}

	query := `DELETE FROM items WHERE id=$1 AND uid=$2`
	if _, err := config.Db["note"].Exec(query, idStr, uidStr); err != nil {
		return new(errors.Internal)
	}

	return nil
}

// Modify given note
func Modify(uidStr, idStr, title, description string) error {
	_, err := strconv.Atoi(uidStr)
	_, err2 := strconv.Atoi(uidStr)
	if err != nil || err2 != nil {
		return new(errors.BadRequest)
	}

	query := `UPDATE items SET edition_date=$3, title=$4, description=$5 WHERE id=$1 AND uid=$2`
	if _, err := config.Db["note"].Exec(query, idStr, uidStr, currentDate(), title, description); err != nil {
		return new(errors.Internal)
	}

	return nil
}

// NoteType provides abstraction on notes
type NoteType struct {
	ID           string `json:"id"`
	UID          string `json:"uid"`
	CreationDate string `json:"creation_date"`
	EditionDate  string `json:"edition_date"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

// Get given note
func Get(uidStr, idStr string) (*NoteType, error) {
	_, err := strconv.Atoi(uidStr)
	_, err2 := strconv.Atoi(uidStr)
	if err != nil || err2 != nil {
		return nil, new(errors.BadRequest)
	}

	query := `SELECT creation_date, edition_date, title, description FROM items WHERE id=$1 AND uid=$2`
	row := config.Db["note"].QueryRow(query, idStr, uidStr)

	// NOTE is there a prettier way to do this?
	note := new(NoteType)
	note.ID = idStr
	note.UID = uidStr
	err = row.Scan(
		&note.CreationDate,
		&note.EditionDate,
		&note.Title,
		&note.Description,
	)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			// note not found
			return nil, new(errors.NotFound)
		default:
			log.Printf("pkg/notes/notes.go:Get: error during results' scan: %s", err)
			return nil, new(errors.Internal)
		}
	}

	return note, nil
}

// All retrieves notes' id from a given id
func All(uidStr string) ([]int, error) {
	if _, err := strconv.Atoi(uidStr); err != nil {
		return nil, new(errors.BadRequest)
	}

	query := `SELECT id FROM items WHERE uid=$1`
	rows, err := config.Db["note"].Query(query, uidStr)
	if err != nil {
		log.Printf("pkg/notes/notes.go:All: error during results' scan: %s", err)
		return nil, new(errors.Internal)
	}
	defer rows.Close()

	var res []int
	for rows.Next() {
		var id int
		if err = rows.Scan(&id); err != nil {
			log.Printf("pkg/notes/notes.go:All: error during results' scan: %s", err)
			return nil, new(errors.Internal)
		}

		res = append(res, id)
	}

	if err = rows.Err(); err != nil {
		log.Printf("pkg/notes/notes.go:All: error during results' scan: %s", err)
		return nil, new(errors.Internal)
	}
	return res, nil
}
