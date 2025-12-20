package repository

import "gorm.io/gorm"

type Repository struct {
	RoleRepository IRoleRepository
	UserRepository IUserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		RoleRepository: NewRoleRepository(db),
		UserRepository: NewUserRepository(db),
	}
}
