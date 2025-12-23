package service

import (
	"errors"
	"golang-register-login/entity"
	"golang-register-login/internal/repository"
	"golang-register-login/model"
	"golang-register-login/pkg/bcyrpt"
	"golang-register-login/pkg/database/mysql"
	"golang-register-login/pkg/jwt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserService interface {
	GetUser(param model.UserParam) (*entity.User, error)
	RegisterUser(param model.UserRegisterParam) (*model.UserRegisterResponse, error)
}

type UserService struct {
	db             *gorm.DB
	UserRepository repository.IUserRepository
	bcyrpt         bcyrpt.Interface
	jwtAuth        jwt.Interface
}

func NewUserService(userRepo repository.IUserRepository, bcrypt bcyrpt.Interface, jwtAuth jwt.Interface) IUserService {
	return &UserService{
		db:             mysql.Connection,
		UserRepository: userRepo,
		bcyrpt:         bcrypt,
		jwtAuth:        jwtAuth,
	}
}

func (u *UserService) GetUser(param model.UserParam) (*entity.User, error) {
	return u.UserRepository.GetUser(param)
}

func (u *UserService) RegisterUser(param model.UserRegisterParam) (*model.UserRegisterResponse, error) {
	tx := u.db.Begin()
	defer tx.Rollback()

	_, err := u.UserRepository.GetUser(model.UserParam{
		Username: param.Username,
	})

	if err == nil {
		return nil, errors.New("username already exists")
	}

	_, err = u.UserRepository.GetUser(model.UserParam{
		Email: param.Email,
	})

	if err == nil {
		return nil, errors.New("email already exists")
	}

	userID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	if param.Password != param.ConfirmPassword {
		return nil, errors.New("Password not match")
	}

	hashPassword, err := u.bcyrpt.GenerateHashPassword(param.Password)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		UserID:   userID,
		RoleID:   2,
		Username: param.Username,
		Email:    param.Email,
		Password: hashPassword,
	}

	err = u.UserRepository.RegisterUser(tx, user)
	if err != nil {
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	response := &model.UserRegisterResponse{
		Username: param.Username,
		Email:    param.Email,
	}

	return response, nil
}
