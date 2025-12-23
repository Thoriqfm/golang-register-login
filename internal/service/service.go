package service

import (
	"golang-register-login/internal/repository"
	"golang-register-login/pkg/bcyrpt"
	"golang-register-login/pkg/jwt"
)

type Service struct {
	UserService IUserService
}

func NewService(repository *repository.Repository, bcrypt bcyrpt.Interface, jwtAuth jwt.Interface) *Service {
	return &Service{
		UserService: NewUserService(repository.UserRepository, bcrypt, jwtAuth),
	}
}
