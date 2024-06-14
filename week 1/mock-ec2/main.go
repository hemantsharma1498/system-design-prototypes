package main

import (
	"encoding/json"
	"io"
	"log"
	"mock-ec2/store"
	"net/http"
	"strings"
	"sync"
	"time"
)

const MAX_RETRIES = 15


type response struct {
  Status string `json:"status"`
}

/*
Use mutex instead of wait groups. 
1. Wait groups will prevent the outer function from returning without the work of mockEc2Creation fn having been completed
2. Use mutex.
3. General note, always have control over spinning of threads over a network. Good concurrency and security practice
*/

func main(){
  
  log.Println("Spun up mock ec2 creation")


  http.HandleFunc("/create-ec2-instance", createEc2Instance)
  http.HandleFunc("/check-creation-status", checkCreationStatus)
  http.ListenAndServe(":3000", nil)

}

func checkCreationStatus(w http.ResponseWriter, r *http.Request){

  store := store.GetDbConnection()
  
  conn, err := store.Connect()
  if err != nil{
    log.Panic("Unable to connect to database")
  } 

  var currentStatus string
  
  serverId := strings.Trim(r.URL.Path, "/check-creation-status/:")
  prevStatus, err := io.ReadAll(r.Body)
  if err != nil{
    log.Panic("Unable to connect to database")
  } 
  tries := 1
  for tries < MAX_RETRIES {
    row := conn.Db.QueryRow("SELECT status FROM ec2_status WHERE id = ?", serverId) 
    if row.Err() != nil {
      log.Panic("Encountered error while checking status", err)
      w.WriteHeader(http.StatusInternalServerError)
      w.Write([]byte("Encountered an error while checking status, please try again"))
    }
    row.Scan(&currentStatus)
    if currentStatus != string(prevStatus) {
      break
    }
  }
  data := response{Status: currentStatus}
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(data)
}

func createEc2Instance(w http.ResponseWriter, r *http.Request){
  var wg sync.WaitGroup
  wg.Add(1)

  serverId := strings.Trim(r.URL.Path, "/create-ec2-instance/:")
  go mockEc2Creation(&wg, serverId)


  data := response{ Status: "Creating" } 
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)
  json.NewEncoder(w).Encode(data)
}

func mockEc2Creation(wg *sync.WaitGroup, serverId string){
  defer wg.Done()
  store := store.GetDbConnection()
  
  conn, err := store.Connect()
  if err != nil{
    log.Panic("Unable to connect to database")
  } 
  
  _, err = conn.Db.Exec("INSERT INTO ec2_status(id, status) VALUES(?, ?)", serverId, "Queued")
  if err != nil {
    log.Panic("Encountered error while creating instance")
  }
  time.Sleep(5 * time.Second) 
  _, err = conn.Db.Exec("UPDATE ec2_status SET status=? WHERE id = ?", "Initialising",serverId)
  if err != nil {
    log.Panic("Encountered error while creating instance")
  }
  time.Sleep(5 * time.Second) 

  _, err = conn.Db.Exec("UPDATE ec2_status SET status=? WHERE id = ?", "Created",serverId)
  if err != nil {
    log.Panic("Encountered error while creating instance")
  }
  

  
}



