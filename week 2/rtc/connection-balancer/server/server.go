package server

import (
	"database/sql"
	"log"
	"net/http"
)


type ConnectionBalancer struct {
  Router *http.ServeMux
  LoadingStatus bool
  serverAddresses map[string]string
}

func InitServer(d *sql.DB) *ConnectionBalancer {
  s := &ConnectionBalancer{ Router: http.NewServeMux(), LoadingStatus: false, serverAddresses: make(map[string]string, 1)}
  go s.LoadCommunicationServers(d)
  s.Routes(d)
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
  if err := c.getServerAddresses(db, 100, 0); err != nil {
    log.Printf("Encountered an error while fetching server addresses: %s\n", err)
  }
  for {
    if c.LoadingStatus==true {
      log.Printf("Loaded addresses for %v servers\n", len(c.serverAddresses))
      break
    }
  }
}

func (c *ConnectionBalancer) getServerAddresses(db *sql.DB, limit int, offset int) error {
  rows, err := db.Query("SELECT org, address FROM communication_servers LIMIT ? OFFSET ?", limit, offset)
  defer rows.Close()
  if err != nil {
    if err == sql.ErrNoRows {
      c.LoadingStatus = true 
      return nil
    }
    return err
  }
  foundRows := false
  for rows.Next(){
    foundRows = true
    var (
      org string
      address string
    )
    if err := rows.Scan(&org, &address); err != nil {
      return err
    }
    if c.serverAddresses[org] == "" {
      c.serverAddresses[org]=address
    }
  }
  
  if !foundRows{
    c.LoadingStatus = true 
    return nil
  }

  c.getServerAddresses(db, limit+100, limit)

  return nil
}

