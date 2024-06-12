package main

import (
	"log"
	"net/http"
	"strings"
  "time"
  "math/rand"
  v2 "math/rand/v2"
)


var serverSideStatus string
var pollPath string = "/short-poll/"

func main(){
  log.Println("Listening and serving on port 3000")
  initServer()
}


func initServer(){
  http.HandleFunc(pollPath, shortPolling)
  http.ListenAndServe(":3000", nil)
}


func shortPolling(w http.ResponseWriter, r *http.Request){
  
  status := strings.TrimPrefix(r.URL.Path, pollPath) 
  
  serverSideStatus=statusChanger()

  if status != serverSideStatus {
    w.Write([]byte("Status hasn't changed")) 
    return
  }

    w.Write([]byte("Status has changed")) 
}


func statusChanger() string {
  numbers := []int{12,13,14,15,16,17,18,19,20}
  rand.Seed(time.Now().Unix())
  randomIndex := v2.IntN(9)
  log.Printf("number selected: %v\n", numbers[randomIndex])
  if numbers[randomIndex]%7 != 0 {
    return "true"
  }
  return "false"
}
