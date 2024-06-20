package main

import (
	"airline-checkin/server"
	"airline-checkin/store"
	"log"
)



func main(){

  log.Printf("Welcome to airline-checkin")  

  log.Printf("Connecting to database...")
  db, err := store.MysqlDbConnector().Connect()
  if err != nil {
    log.Panicf("Encountered error while initialising database: %s\n", err)
  }
  s := server.InitServer(db.Db)

  if err := s.Start(":8080"); err != nil {
    log.Panicf("Encountered error while initialising server: %s\n", err)
  }


}








