package server

import (
  "database/sql"
  "net/http"
)

func (s *HttpServer) Routes(d *sql.DB){
  s.Router.HandleFunc("/book-seats", func(w http.ResponseWriter, r *http.Request) {
    s.BookSeats(w, r, d)
  })
}
