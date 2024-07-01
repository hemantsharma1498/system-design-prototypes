package server

import (
	"database/sql"
	"encoding/json"
	"log"
	"members/server/dto"
	"net/http"
	"strings"
)



func (m *Members) GetCommServerAddress(w http.ResponseWriter, r *http.Request, db *sql.DB){


}


func (m *Members) Login (w http.ResponseWriter, r *http.Request, db *sql.DB) {
  data := r.Body
}
