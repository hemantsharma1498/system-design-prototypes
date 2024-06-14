package store

import (
	"database/sql"
)

var dbPath = "./sqlite3Db.sql"
type Sqlite3 struct {
   Db *sql.DB
}

func GetDbConnection() *Sqlite3{
  return &Sqlite3{}
}

func (s *Sqlite3) Connect() (*Sqlite3, error ) {
  db, err := sql.Open("Sqlite3", dbPath)
  if err != nil{
    return nil,err
  }
  s.Db=db
  
  go func() (*Sqlite3, error) {
    _, err := db.Exec("CREATE TABLE IF NOT EXISTS ec2_status(id INTEGER PRIMARY KEY AUTOINCREMENT, status TEXT);")
    if err != nil {
      return nil, err
    }
    return s, err
  } ()

  return s, nil
}


