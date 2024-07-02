package types

//@API level representation of data stored in Mysql db

type Account struct {
  Id int64 
  Email string
  PasswordHash string
  Salt string
}

