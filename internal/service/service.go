package service

import (
	"golang-register-login/internal/repository"
	"golang-register-login/pkg/bcyrpt"
	"golang-register-login/pkg/email"
	"golang-register-login/pkg/jwt"
)

type Service struct {
	UserService          IUserService
	ResetPasswordService IResetPasswordService
}

func NewService(repository *repository.Repository, bcrypt bcyrpt.Interface, jwtAuth jwt.Interface, email email.Interface) *Service {
	return &Service{
		UserService:          NewUserService(repository.UserRepository, bcrypt, jwtAuth),
		ResetPasswordService: NewResetPasswordService(repository.UserRepository, repository.ResetPasswordRepository, bcrypt, email),
	}
}
