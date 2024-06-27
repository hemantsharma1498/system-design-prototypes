package server

import (
	"database/sql"
	"net/http"
)



func (c *ConnectionBalancer) GetCommServerAddress(w http.ResponseWriter, r *http.Request, db *sql.DB){
  org := r.PathValue("org")
  if len(org) == 0 {
      
  }
  
}
