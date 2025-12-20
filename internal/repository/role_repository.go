package repository

import "gorm.io/gorm"

type IRoleRepository interface {
}

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) IRoleRepository {
	return &RoleRepository{db: db}
}
