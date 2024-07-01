package store

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const dsn = "hemant:1@Million@tcp(localhost)/connection_balancer"
const migrationDir = "./store/migrations"
type ConnBal struct {
  Db *sql.DB
  host string
  port string
}

func NewConnBalConnector() *ConnBal {
  return &ConnBal{}
}


func (c *ConnBal) Connect(ctx context.Context) (*ConnBal, error) {
  if c.Db == nil {
    var err error
    c.Db, err = initDb(ctx)
    if err != nil {
      return nil, err
    }
    return c, nil
  }
  return c, nil
}

func initDb(ctx context.Context) (*sql.DB, error) {
  
  db, err := sql.Open("mysql", dsn)
  if err != nil {
    return nil, err
  }

  if err = db.Ping() ; err != nil {
    return nil, err
  }

  return db, nil
}

