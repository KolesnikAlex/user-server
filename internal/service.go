package internal

import (
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

func (s Service) GetUser(id int64) (service.User, error) {
	res, err := s.userRepo.GetUser(id)
	if err != nil {
		log.Error().Err(err).Msg("errService GetUser")
		return service.User{}, err
	}
	return res, err
}

func (s Service) AddUser(user service.User) error {
	err := s.userRepo.AddUser(user)
	if err != nil {
		log.Error().Err(err).Msg("errService AddUser")
		return err
	}
	return err
}

func (s Service) RemoveUser(id int64) error {
	err := s.userRepo.RemoveUser(id)
	if err != nil {
		log.Error().Err(err).Msg("errService RemoveUser")
		return err
	}
	return err
}

func (s Service) UpdateUser(user service.User) error {
	err := s.userRepo.UpdateUser(user)
	if err != nil {
		log.Error().Err(err).Msg("errService UpdateUser")
		return err
	}
	return err
}

