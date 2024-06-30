package server

import (
	"database/sql"
	"log"
	"net/http"
)


type Members struct {
  Router *http.ServeMux
  LoadingStatus bool
}

func InitServer(d *sql.DB) *Members {
  s := &Members{ Router: http.NewServeMux(), LoadingStatus: false, }
  s.Routes(d)
  return s
}

func (c *Members) Start(address string) error {
  log.Printf("Starting members server at address: %s\n", address)
  if err := http.ListenAndServe(address, c.Router) ; err != nil{
    return err
  }
  return nil
}


