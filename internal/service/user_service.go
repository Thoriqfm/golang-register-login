package service

import (
	"golang-register-login/internal/repository"
	"golang-register-login/pkg/database/mysql"

	"gorm.io/gorm"
)

type IUserService interface {
}

type UserService struct {
	db             *gorm.DB
	UserRepository repository.IUserRepository
}

func NewUserService(userRepo repository.IUserRepository) IUserService {
	return &UserService{
		db:             mysql.Connection,
		UserRepository: userRepo,
	}
}
