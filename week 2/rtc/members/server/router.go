package server

import (
	"database/sql"
	"net/http"
)

func (m *Members) Routes(db *sql.DB){
  m.Router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
    m.Login(w, r, db) 
  })
  m.Router.HandleFunc("/get-connection", func (w http.ResponseWriter, r *http.Request){
    m.GetCommServerAddress(w, r, db) 
  })
}

