package service

import "golang-register-login/internal/repository"

type Service struct {
	UserService IUserService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		UserService: NewUserService(repository.UserRepository),
	}
}
