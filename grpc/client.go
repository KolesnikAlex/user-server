package grpc

import (
	"context"
	"encoding/json"
	"log"
	"time"

	grpcUserService "github.com/KolesnikAlex/user-service-proto/grpc"
	localUserService "github.com/KolesnikAlex/user-server/app/service"
	"google.golang.org/grpc"
)


var defaultRequestTimeout = time.Second * 10

type GrpcService struct {
	grpcClient grpcUserService.GrpcUserServiceClient
}

// NewGRPCService creates a new gRPC user service connection using the specified connection string.
func NewGRPCService(connString string) (*GrpcService, error) {
	conn, err := grpc.Dial(connString, grpc.WithInsecure())
	if err != nil {
		log.Printf("error connection to user-server")
		return nil, err
	}
	return &GrpcService{grpcClient: grpcUserService.NewGrpcUserServiceClient(conn)}, nil
}

func (g GrpcService) GetUser(id int64) (user localUserService.User, err error) {
	user = localUserService.User{}
	req := &grpcUserService.Id{
		Id: id,
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	resp, err := g.grpcClient.GetUser(ctx, req)
	if err != nil {
		log.Printf("error get User from user-server")
		return user, err
	}
	user = unMarshalUser(resp)
	log.Printf("success GetUser(%v)\n", id)
	return user, err
}

func (g GrpcService) AddUser(user *localUserService.User) (err error) {
	req := marshalUser(user)
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	_, err = g.grpcClient.AddUser(ctx, req)
	if err != nil {
		log.Printf("error add User to user-server")
		return err
	}
	userJson, _ := json.Marshal(user)
	log.Printf("success AddUser(%v)\n", userJson)
	return nil
}

func (g GrpcService) RemoveUser(id int64) (err error) {
	req := &grpcUserService.Id{
		Id: id,
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	_, err = g.grpcClient.RemoveUser(ctx, req)
	if err != nil {
		log.Printf("error remove User from user-server")
		return err
	}
	log.Printf("success RemoveUser(%v)\n", id)
	return nil
}

func (g GrpcService) UpdateUser(user *localUserService.User) (err error) {
	req := marshalUser(user)
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	_, err = g.grpcClient.UpdateUser(ctx, req)
	if err != nil {
		log.Printf("error update User to user-server")
		return err
	}
	userJson, _ := json.Marshal(user)
	log.Printf("success UpdateUser(%v)\n", userJson)
	return nil
}

