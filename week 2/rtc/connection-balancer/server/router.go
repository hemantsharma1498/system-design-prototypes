package server

import (
	"database/sql"
	"net/http"
)

func (c *ConnectionBalancer) Routes(db *sql.DB){
  c.Router.HandleFunc("/get-cserver-addresses/{org}", func (w http.ResponseWriter, r *http.Request){
    c.GetCommServerAddress(w, r, db) 
  })
}

