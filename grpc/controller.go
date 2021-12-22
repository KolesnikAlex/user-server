package grpc

import (
	"context"
	"log"
	"user-server/app"
	"user-server/app/service"
	"user-server/config"
	"user-server/internal/database"

)

// userServiceController implements the gRPC UserServiceServer interface.
type userServiceController struct {
	userService service.UserService
}

// NewUserServiceController instantiates a new UserServiceServer.
func NewUserServiceController(userService service.UserService) GrpcUserServiceServer {
	return &userServiceController{
		userService: userService,
	}
}

// GetUsers calls the core service's GetUsers method and maps the result to a grpc service response.
func (ctlr *userServiceController) GetUser(req *Id) (resp *User, err error) {
	result, err := ctlr.userService.GetUser(int64(req.GetId()))
	if err != nil {
		return
	}

	resp = &User{}
	resp = marshalUser(&result)
	log.Printf("handled GetUser(%v)\n", req.GetId())
	return
}

// marshalUser marshals a business object User into a gRPC layer User.
func marshalUser(u *service.User) *User {
	return &User{
		Id: int32(u.ID),
		Name:     u.Name,
		Login:    u.Login,
		Password: u.Password,
	}
}

