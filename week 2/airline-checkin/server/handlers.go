package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type users struct {
  Users []user `json:"users"`
}
type user struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Seat string `json:"seat"`
}

func (s *HttpServer) BookSeats(w http.ResponseWriter, r *http.Request, d *sql.DB) {
  rows, err := d.Query("SELECT * FROM customers LIMIT 5;")
  if err != nil {
    w.Write([]byte("Encountered an error"))
    w.WriteHeader(http.StatusInternalServerError)
  }
  users := users{Users: make([]user, 0)}
  for rows.Next(){
    var (
      id int
      name string
    ) 
    if err := rows.Scan(&id, &name); err != nil {
      fmt.Println(err)
      w.Write([]byte("Encountered an error"))
    }
    user := user{Name: name, Id: id}
    users.Users = append(users.Users, user)
  }
  var wg sync.WaitGroup 
  wg.Add(len(users.Users))
  for i := 0;i<len(users.Users);i++{
    go func(user *user){
      seat, err := book(user, d)
      if err != nil {
        fmt.Println("Encountered error while booking seat", err)
      } else {
        user.Seat=seat
        fmt.Printf("%s got %s\n", user.Name, seat)
      }
      wg.Done()
    }(&users.Users[i])
  }
  wg.Wait()


  data, err := json.Marshal(users)
  if err != nil {
    w.Write([]byte("Encountered some error"))
  }
  w.Write([]byte(data))
}

func book(user *user, d *sql.DB) (string, error){
  tx, err := d.Begin()
  if err != nil {
    return "", err
  }
  row := tx.QueryRow("SELECT seat_number FROM seats WHERE customer_id IS NULL ORDER BY seat_number LIMIT 1 FOR UPDATE SKIP LOCKED;") 
  if row.Err() != nil {
    return "", row.Err()
  }
  var seatNumber string
  row.Scan(&seatNumber)
  
  _, err = tx.Exec("UPDATE seats SET customer_id = ?, customer_name = ? WHERE seat_number = ?", user.Id, user.Name, seatNumber)
  if err != nil {
    return "", err 
  }
  err = tx.Commit()
  if err != nil {
    return "", err 
  }
  return seatNumber, nil
}
