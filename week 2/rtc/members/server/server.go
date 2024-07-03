package server

import (
	"log"
	"members/store"
	"net/http"
)

type Members struct {
	listenAddress string
	Router        *http.ServeMux
	store         store.Storage
}

func InitServer(listenAddress string, store store.Storage) *Members {
	s := &Members{listenAddress: listenAddress, Router: http.NewServeMux(), store: store}
    s.Routes()
	return s
}

func (m *Members) Start() error {
	log.Printf("Starting members server at address: %s\n", m.listenAddress)
	if err := http.ListenAndServe(m.listenAddress, m.Router); err != nil {
		return err
	}
	return nil
}
