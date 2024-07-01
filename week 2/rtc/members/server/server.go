package server

import (
	"database/sql"
	"log"
	"net/http"
)


type Members struct {
  Router *http.ServeMux
  LoadingStatus bool
  serverAddresses map[string]string
}

func InitServer(d *sql.DB) *Members {
  s := &Members{ Router: http.NewServeMux(), LoadingStatus: false, serverAddresses: make(map[string]string, 1)}
  go s.LoadCommunicationServers(d)
  s.Routes(d)
  return s
}

func (m *Members) Start(address string) error {
  log.Printf("Starting connection balancer server at address: %s\n", address)
  if err := http.ListenAndServe(address, m.Router) ; err != nil{
    return err
  }
  return nil
}


func (m *Members) LoadCommunicationServers(db *sql.DB){
  if err := m.getServerAddresses(db, 100, 0); err != nil {
    log.Printf("Encountered an error while fetching server addresses: %s\n", err)
  }
  for {
    if m.LoadingStatus==true {
      log.Printf("Loaded addresses for %v servers\n", len(m.serverAddresses))
      break
    }
  }
}

func (m *Members) getServerAddresses(db *sql.DB, limit int, offset int) error {
  rows, err := db.Query("SELECT org, address FROM communication_servers LIMIT ? OFFSET ?", limit, offset)
  defer rows.Close()
  if err != nil {
    if err == sql.ErrNoRows {
      m.LoadingStatus = true 
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
    if m.serverAddresses[org] == "" {
      m.serverAddresses[org]=address
    }
  }
  
  if !foundRows{
    m.LoadingStatus = true 
    return nil
  }

  m.getServerAddresses(db, limit+100, limit)

  return nil
}

