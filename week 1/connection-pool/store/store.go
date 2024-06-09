package store

import (
	//"context"
	"database/sql"
	"fmt"
	"sync"

	//"os"
	_ "github.com/mattn/go-sqlite3"
)

type Pool struct {
  conn *sql.DB
  mutex sync.Mutex
  connectionChannel chan *sql.DB
  maxConnections int32
}
var dbFilePath="./_sqliteDb.db"

func initConnection() (*sql.DB, error){
  db, err := sql.Open("sqlite3", dbFilePath)
  if err != nil {
    return nil, err
  }
  return db, nil
}

func (p *Pool) CreatePool(maxConnections int32){
  fmt.Println("Hello from CreatePool")

  p.maxConnections=maxConnections
  var err error 
  p.conn, err = initConnection()
  if err != nil {
    fmt.Errorf("failed to create connection with db: %v\n", err)
  }
  p.connectionChannel = make(chan *sql.DB, p.maxConnections)
  
}

func (p *Pool) GiveConnection() *sql.D

}


