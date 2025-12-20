package entity

import "github.com/google/uuid"

type User struct {
	UserID   uuid.UUID `json:"user_id" gorm:"type:char(36);primaryKey"`
	Username string    `json:"username" gorm:"type:varchar(255);not null"`
	Email    string    `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password string    `json:"password" gorm:"type:varchar(255);not null"`
	RoleID   int       `json:"role_id" gorm:"type:int;not null"`
	// Role     Role      `json:"role" gorm:"foreignKey:RoleID"`
}
