package main

import (
	"log"
	"math/rand"
	v2 "math/rand/v2"
	"net/http"
	"strings"
	"time"
)


var serverSideStatus string
var pollPath string = "/long-poll/"
const MAX_RETRIES = 15 
func main(){
  log.Println("Listening and serving on port 3000")
  initServer()
}


func initServer(){
  http.HandleFunc(pollPath, longPolling)
  http.ListenAndServe(":3000", nil)
}


func longPolling(w http.ResponseWriter, r *http.Request){
  
  status := strings.TrimPrefix(r.URL.Path, pollPath) 
    
  tries := 0
  for tries < MAX_RETRIES {
    serverSideStatus=statusChanger()
    if status != serverSideStatus {
      w.Write([]byte("Status has changed")) 
      return
    }
    log.Println(tries)
    time.Sleep(1 * time.Second)
    tries++
  }
  w.Write([]byte("Request timeout, try again")) 
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
