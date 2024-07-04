package server

import (
	"log"
	//"members/pkg/proto"
	"members/store"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Members struct {
	listenAddressHttp string
	Router            *http.ServeMux
	GrpcClient        *grpc.ClientConn
	store             store.Storage
}

func InitServer(listenAddressHttp string, store store.Storage) *Members {
	s := &Members{listenAddressHttp: listenAddressHttp, Router: http.NewServeMux(), store: store}
	s.Routes()
	return s
}

func (m *Members) Start() error {
	log.Printf("Starting members server at address: %s\n", m.listenAddressHttp)

	// Set up a connection to the server. @TODO CHANGE THE HARDCODED VALUE
	m.GrpcClient, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//c := proto.NewGreeterClient(conn)

	if err := http.ListenAndServe(m.listenAddressHttp, m.Router); err != nil {
		return err
	}
	return nil
}
