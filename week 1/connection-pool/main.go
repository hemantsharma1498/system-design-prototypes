package main

import (
	"connection-pool/store"
	"fmt"
	"time"
)



func main(){
  
  fmt.Println("Hello from connection-pooling")
  var maxConnections int
  /*
  fmt.Println("Enter the maximum number of connections you want in the pool: (1-20)")
  for {
    fmt.Scanln(&maxConnections)
    if maxConnections>20{
      fmt.Println("Please enter a value between 1 and 20")
    }
    fmt.Scanln(&maxConnections)
  }
  fmt.Println("Enter the number of queries for test: ")
  */
  maxConnections=10

  pool:=&store.Pool{
    MaxConnections:int32(maxConnections),
    ConnectionChannel: make(chan *store.Connection, maxConnections),
  }

  fmt.Println("Creating pool...") 
  pool.CreatePool()

  fmt.Printf("Connection pool created with %v connections\n", maxConnections) 
  fmt.Println("Beginning test...") 


  for i:=0; i<500;i++{
    conn := pool.GiveConnection()
    err := pool.QueryDb(conn.Conn)
    if err != nil{
      fmt.Printf("encountered error: %s\n", err)
      continue
    }
    fmt.Printf("Current iteration: %v | Db connection ID: %v\n", i+1, conn.ID)
    time.Sleep(time.Millisecond*200)
    pool.ReturnConnection(conn)
  }

}

