package grpc

import (
	"context"
	grpcUserService "github.com/KolesnikAlex/user-service-proto/grpc"
	//"encoding/json"
	"log"
	"github.com/KolesnikAlex/user-server/app/service"
)

// userServiceController implements the gRPC UserServiceServer interface.
type userServiceController struct {
	userService service.UserService
}

// NewUserServiceController instantiates a new UserServiceServer.
func NewUserServiceController(userService service.UserService) grpcUserService.GrpcUserServiceServer {
	return &userServiceController{
		userService: userService,
	}
}

func (ctlr *userServiceController) GetUser(ctx context.Context, id *grpcUserService.Id) (resp *grpcUserService.User, err error) {
	result, err := ctlr.userService.GetUser(id.GetId())
	if err != nil {
		log.Printf("error handl GetUser(%v)\n", id.GetId())
		return &grpcUserService.User{}, err
	}
	resp = &grpcUserService.User{}
	resp = marshalUser(&result)
	log.Printf("handled GetUser(%v)\n", id.GetId())
	return
}

func (ctlr *userServiceController) AddUser(ctx context.Context, grpcUser *grpcUserService.User) (req *grpcUserService.Request, err error) {
	user := unMarshalUser(grpcUser)
	err = ctlr.userService.AddUser(user)
	if err != nil {
		log.Printf("error handl AddUser(%s)\n", user)
		return &grpcUserService.Request{}, err
	}
	log.Printf("handled AddUser(%s)\n", user)
	return &grpcUserService.Request{}, err
}

func (ctlr *userServiceController) RemoveUser(ctx context.Context, id *grpcUserService.Id) (req *grpcUserService.Request, err error) {
	err = ctlr.userService.RemoveUser(id.GetId())
	if err != nil {
		log.Printf("error handl RemoveUser(%v)\n", id.GetId())
		return &grpcUserService.Request{}, err
	}
	log.Printf("handled RemoveUser(%v)\n", id.GetId())
	return &grpcUserService.Request{}, err
}

func (ctlr *userServiceController) UpdateUser(ctx context.Context, grpcUser *grpcUserService.User) (req *grpcUserService.Request, err error) {
	user := unMarshalUser(grpcUser)
	err = ctlr.userService.UpdateUser(user)
	if err != nil {
		log.Printf("error handl UpdateUser(%s)\n", user)
		return &grpcUserService.Request{}, err
	}
	log.Printf("handled UpdateUser(%s)\n", user)
	return &grpcUserService.Request{}, err
}

//func (ctlr *userServiceController) mustEmbedUnimplementedGrpcUserServiceServer() {
//	panic("implement me")
//}


// marshalUser marshals a business object User into a gRPC layer User.
func marshalUser(localUser *service.User) *grpcUserService.User {
	return &grpcUserService.User{
		Id:       localUser.ID,
		Name:     localUser.Name,
		Login:    localUser.Login,
		Password: localUser.Password,
	}
}

// unMarshalUser unmarshals a gRPC layer User into a business object User.
func unMarshalUser(grpcUser *grpcUserService.User) service.User {
	return service.User{
		ID:       grpcUser.Id,
		Name:     grpcUser.Name,
		Login:    grpcUser.Login,
		Password: grpcUser.Password,
	}
}

