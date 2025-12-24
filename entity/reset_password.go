package entity

import (
	"time"

	"github.com/google/uuid"
)

type PasswordResetToken struct {
	ID        uuid.UUID `gorm:"type:char(36);primayKey"`
	UserID    uuid.UUID `gorm:"type:char(36);not null"`
	Token     string    `gorm:"type:varchar(255);uniqe;not null"`
	ExpiredAt time.Time `gorm:"not null"`
	Used      bool      `gorm:"default:false"`
	CreatedAt time.Time
}
