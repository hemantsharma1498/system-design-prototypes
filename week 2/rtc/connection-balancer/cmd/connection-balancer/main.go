package main

import (
	"connection-balancer/pkg/proto"
	"connection-balancer/server"
	"connection-balancer/store"
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const httpAddress = ":8080"
const grpcAddress = ":8081"

func main() {
	log.Printf("Initialising connection balancer")
	log.Printf("Connecting to database...")
	ctx := context.Background()
	store, err := store.NewConnBalConnector().Connect(ctx)
	if err != nil {
		log.Panicf("Unable to connect to db, error: %s\n", err)
	}
	log.Printf("Db connection established")
	s := server.InitServer(store.Db)
	fmt.Println("s: ", s.InstanceID)
	/*
		g := grpc.NewServer()
		proto.RegisterConnectionServer(g, s)
		log.Printf("Listening at %v", grpcAddress)
		if err := g.Serve(grpcAddress); err != nil {
			log.Fatalf("failed to serve :%v", err)
		}
	*/
	go sampleClientCall()
	if err = s.Start(httpAddress, grpcAddress); err != nil {
		log.Panicf("Failed to initialise server at %s, error: %s\n", httpAddress, err)
	}

}

func sampleClientCall() {
	time.Sleep(time.Second * 10)
	// Set up a connection to the server.
	var addr = flag.String("addr", "localhost:50051", "the address to connect to")
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewConnectionClient(conn)
	res, err := c.GetCommServerAddr(context.TODO(), &proto.GetCommServerAddrReq{Org: "Org_99"})
	if err != nil {
		log.Fatalf("Encountered and error:  %v\n", err)
	}
	log.Println(res)
	//log.Printf("Response: %v\n", res.Address)
}
