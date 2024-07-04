package server

import (
	"connection-balancer/pkg/proto"
	"context"
	"database/sql"
	"flag"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

type ConnectionBalancer struct {
	Router          *http.ServeMux
	GrpcServer      *grpc.Server
	LoadingStatus   bool
	serverAddresses map[string]string
	proto.ConnectionServer
}

func InitServer(d *sql.DB) *ConnectionBalancer {
	s := &ConnectionBalancer{Router: http.NewServeMux(), LoadingStatus: false, serverAddresses: make(map[string]string, 1)}
	go s.LoadCommunicationServers(d)
	s.Routes(d)
	return s
}

func (c *ConnectionBalancer) GetCommServerAddr(ctx context.Context, in *proto.GetCommServerAddrReq) (*proto.GetCommServerAddrReply, error) {
	return &proto.GetCommServerAddrReply{Address: "address"}, nil
}

func (c *ConnectionBalancer) Start(httpAddr string, grpcAddr string) error {
	flag.Parse()
	log.Printf("Starting http server at: %s\n", httpAddr)

	var grpcErr error
	//Start Grpc Server
	go func() {
		lis, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		c.GrpcServer = grpc.NewServer()
		proto.RegisterConnectionServer(c.GrpcServer, &ConnectionBalancer{})
		log.Printf("Grpc server listening at %v", lis.Addr())
		if err := c.GrpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve grpc: %v", err)
			grpcErr = err
		}
	}()
	if grpcErr != nil {
		return grpcErr
	}
	//Start Http server
	if err := http.ListenAndServe(httpAddr, c.Router); err != nil {
		log.Fatalf("failed to serve http: %v", err)
		return err
	}
	return nil
}

func (c *ConnectionBalancer) LoadCommunicationServers(db *sql.DB) {
	if err := c.getServerAddresses(db, 100, 0); err != nil {
		log.Printf("Encountered an error while fetching server addresses: %s\n", err)
	}
	for {
		if c.LoadingStatus == true {
			log.Printf("Loaded addresses for %v servers\n", len(c.serverAddresses))
			break
		}
	}
}

func (c *ConnectionBalancer) getServerAddresses(db *sql.DB, limit int, offset int) error {
	rows, err := db.Query("SELECT org, address FROM communication_servers LIMIT ? OFFSET ?", limit, offset)
	defer rows.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			c.LoadingStatus = true
			return nil
		}
		return err
	}
	foundRows := false
	for rows.Next() {
		foundRows = true
		var (
			org     string
			address string
		)
		if err := rows.Scan(&org, &address); err != nil {
			return err
		}
		if c.serverAddresses[org] == "" {
			c.serverAddresses[org] = address
		}
	}

	if !foundRows {
		c.LoadingStatus = true
		return nil
	}

	c.getServerAddresses(db, limit+100, limit)

	return nil
}
