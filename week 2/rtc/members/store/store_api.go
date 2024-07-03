package store

import (
	"members/store/types"
)

type Storage interface {
	CreateAccount(*types.MemberAccount) error
	GetMemberByEmail([]string) (*types.MemberAccount, error)
}

type Connecter interface {
	Connect() (Storage, error)
}
