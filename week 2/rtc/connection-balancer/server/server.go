package server

import (
	"database/sql"
	"log"
	"net/http"
)


type ConnectionBalancer struct {
  Router *http.ServeMux
}



func InitServer(d *sql.DB) *ConnectionBalancer {
  s := &ConnectionBalancer{ Router: http.NewServeMux(), }
  s.LoadCommunicationServers(d)
  return s
}

func (c *ConnectionBalancer) Start(address string) error {
  log.Printf("Starting connection balancer server at address: %s\n", address)
  if err := http.ListenAndServe(address, c.Router) ; err != nil{
    return err
  }
  return nil
}


func (c *ConnectionBalancer) LoadCommunicationServers(db *sql.DB){
  
}
