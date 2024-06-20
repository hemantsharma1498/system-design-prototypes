package server

import (
	"database/sql"
	"log"
	"net/http"
)



type HttpServer struct {
  Router *http.ServeMux
}




func InitServer(d *sql.DB) *HttpServer {
  s := &HttpServer{ Router: http.NewServeMux(), }
  s.Routes(d)
  return s
}


func (s *HttpServer) Start(address string) error {
  log.Printf("Starting airline checkin server at address: %s\n", address)
  if err := http.ListenAndServe(address, s.Router); err != nil {
    return err
  }
  return nil
}




