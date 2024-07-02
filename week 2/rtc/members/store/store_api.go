package store

import (
	"members/store/types"
)


type Storage interface {
  CreateAccount(*types.Account) error
  GetMemberByEmail(string) (*types.Account, error)
}


type Connecter interface {
	Connect() (Storage, error)
}
