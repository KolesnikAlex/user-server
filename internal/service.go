package internal

import (
	"context"
	"github.com/phuslu/log"
	"user-server/app/service"
)

type Service struct {
	userRepo service.Service
}

func newService(service service.Service) *Service {
	return &Service{
		userRepo: service,
	}
}

func (s Service) GetUser(ctx context.Context, id int64) (service.User, error) {
	res, err := s.userRepo.GetUser(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("errService GetUser")
		return service.User{}, err
	}
	return res, err
}

func (s Service) AddUser(ctx context.Context, user service.User) error {
	err := s.userRepo.AddUser(ctx, user)
	if err != nil {
		log.Error().Err(err).Msg("errService AddUser")
		return err
	}
	return err
}

func (s Service) RemoveUser(ctx context.Context, id int64) error {
	err := s.userRepo.RemoveUser(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("errService RemoveUser")
		return err
	}
	return err
}

func (s Service) UpdateUser(ctx context.Context, user service.User) error {
	err := s.userRepo.UpdateUser(ctx, user)
	if err != nil {
		log.Error().Err(err).Msg("errService UpdateUser")
		return err
	}
	return err
}

