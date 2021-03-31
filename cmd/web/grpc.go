package main

import (
	"log"
	"net"

	"github.com/tkircsi/vault/api/grpc/server"
	"github.com/tkircsi/vault/api/grpc/vaultpb"
	"google.golang.org/grpc"
)

func (app *application) RunGRPC() {
	l, err := net.Listen("tcp", "0.0.0.0"+app.GRPCPort)
	if err != nil {
		log.Fatalf("grpc: failed listen: %v\n", err)
	}

	s := grpc.NewServer()
	h := server.NewGRPCHandler(app.vault)
	vaultpb.RegisterVaultServiceServer(s, h)

	log.Println("vault gRPC service started...")
	log.Fatal(s.Serve(l))
}
