package store

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
  "log"
)

const dbPath="hemant:1@Million@tcp(localhost)/airline"


type MysqlDb struct {
  Db *sql.DB
}

func MysqlDbConnector() *MysqlDb {
  return &MysqlDb{}
}

func (d *MysqlDb) Connect() (*MysqlDb, error) {
  db, err := initDb()
  if err != nil {
    return nil, err
  }
  d.Db = db
  log.Println("Connected to sql database")
  return d, nil
}

func initDb() (*sql.DB, error) {
  db, err := sql.Open("mysql", dbPath)
  if err != nil {
    return nil, err
  }

  if err = db.Ping(); err != nil {
    return nil, err
  }
  return db, nil
}
