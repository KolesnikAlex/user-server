package main

import (
	"fmt"

	"github.com/KolesnikAlex/user-server/grpc"
	"github.com/phuslu/log"
)

func main() {
	// connect to grpcServer
	grpcServ, err := grpc.NewGRPCService("localhost:9000")
	if err != nil {
		log.Fatal().Err(err).Msg("error grpc connection:")
		return
	}
	log.Printf("connection is: %v", grpcServ)
	user1, err := grpcServ.GetUser(1)
	if err != nil {
		log.Fatal().Err(err).Msg("error GetUser:")
		return
	}
	fmt.Println(user1)
}
