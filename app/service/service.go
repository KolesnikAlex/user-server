package service

import (
	"context"
	"errors"
)

// ErrNotFound signifies that a single requested object was not found.
var ErrNotFound = errors.New("not found")

// User is a user business object.
type User struct {
	ID   int64
	Name string
	Login string
	Password string
}

// Service defines the interface exposed by this package.
type Service interface {
	GetUser(ctx context.Context, id int64) (User, error)
	AddUser(ctx context.Context, user User) (error)
	RemoveUser(ctx context.Context, id int64) (error)
	UpdateUser(ctx context.Context, user User) (error)
}
