package repository

import (
	"golang-register-login/entity"
	"golang-register-login/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	RegisterUser(tx *gorm.DB, user *entity.User) error
	GetUser(param model.UserParam) (*entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) RegisterUser(tx *gorm.DB, user *entity.User) error {
	err := tx.Debug().Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUser(param model.UserParam) (*entity.User, error) {
	var user entity.User

	err := r.db.Where(&param).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
