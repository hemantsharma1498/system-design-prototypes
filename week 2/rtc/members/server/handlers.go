package server

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)


type ServerAddressResponse struct {
  Address string `json:"Address"`
}


func (m *Members) GetCommServerAddress(w http.ResponseWriter, r *http.Request, db *sql.DB){
//  org := r.PathValue("org")
  orgArr := strings.Split(r.URL.Path, "/")
  org := orgArr[len(orgArr)-1] 
  if len(m.serverAddresses) == 0 {
    m.LoadingStatus=false
    m.LoadCommunicationServers(db)
  }
  var address string = ""
  address = m.serverAddresses[org] 
  if address == ""{
    log.Printf("No address found for %s\n", org)
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("No address found for the given org name"))
    return 
  }
  
  resp := ServerAddressResponse{Address: address}
  data, err := json.Marshal(resp)
  if err != nil {
    log.Printf("Encountered an error while marshalling address for %s\n", org)
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("Server encountered an error. Please try again."))
    return
  }
  w.WriteHeader(http.StatusOK)
  w.Write(data)
}

