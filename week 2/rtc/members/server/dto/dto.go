package dto



type Login struct {
  Email string `json:"email"` 
  Password string `json:"password"`
}

type SignUp struct {
  Email string `json:"email"` 
  FirstName string `json:"firstName"`
  LastName string `json:"lastName"`
  Organisation string `json:"org"`
}
