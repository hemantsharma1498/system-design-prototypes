package server

import (
	"database/sql"
	"net/http"
)

func (m *Members) Routes(db *sql.DB){
  m.Router.HandleFunc("/get-cserver-addresses/{org}", func (w http.ResponseWriter, r *http.Request){
    m.GetCommServerAddress(w, r, db) 
  })
}

