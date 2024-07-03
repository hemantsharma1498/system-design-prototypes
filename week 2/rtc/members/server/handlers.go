package server

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func (m *Members) GetCommServerAddress(w http.ResponseWriter, r *http.Request, db *sql.DB) {

}

func (m *Members) Login(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	d := &LoginReq{}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		writeResponse(w, err, "Encountered an error. Please try again", http.StatusInternalServerError)
		return
	}
	row := db.QueryRow("SELECT name, email FROM members WHERE email_id = ?", d.Email)
	if row.Err() == sql.ErrNoRows {
		writeResponse(w, nil, "No user found for the given email", http.StatusBadRequest)
		return
	}
}

func writeResponse(w http.ResponseWriter, err error, msg any, httpStatus int) error {
	if err != nil {
		log.Printf("Error occured while decoding req json body: %s\n", err)
	}
	w.WriteHeader(httpStatus)
	return json.NewEncoder(w).Encode(msg)
}
