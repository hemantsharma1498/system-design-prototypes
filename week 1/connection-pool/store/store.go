package store

import (
	"math/rand/v2"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)


type Connection struct {
  Conn *sql.DB
  ID int
}

type Pool struct {
  ConnectionChannel chan *Connection
  MaxConnections int32
}
var dbFilePath="./_sqliteDb.db"


func initConnection() (*sql.DB, error){

  _, err := os.Stat(dbFilePath)
  if err != nil{
    if os.IsNotExist(err){
      os.Create("_sqliteDb.db")
    }
  }
  db, err := sql.Open("sqlite3", dbFilePath)
  if err != nil {
    return nil, err
  }
  return db, nil
}

func (p *Pool) CreatePool() error {
  fmt.Println("Hello from CreatePool")
  p.ConnectionChannel = make(chan *Connection, p.MaxConnections)
  for i :=1;i<int(p.MaxConnections);i++{

    conn, err := initConnection()
    if err != nil {
      return fmt.Errorf("failed to create connection with db: %v\n", err)
    }
    
    p.ConnectionChannel<-&Connection{Conn: conn, ID: rand.IntN(int(p.MaxConnections))}
  }
  return nil
}

func (p *Pool) GiveConnection() *Connection {
  return <-p.ConnectionChannel
}

func (p *Pool) ReturnConnection(conn *Connection){
  p.ConnectionChannel<-conn
}

func (p *Pool) QueryDb(conn *sql.DB) error {
  
  _, err := conn.Query("SELECT 1;")
  if err != nil{
    return fmt.Errorf("query failed with error: %v\n", err)
  } 
  return nil
}
