package repository

import (
	"golang-register-login/entity"
	"golang-register-login/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IResetPasswordRepository interface {
	CreatePasswordResetToken(tx *gorm.DB, token *entity.PasswordResetToken) error
	GetResetToken(param model.UserForgotPasswordParam) (*entity.PasswordResetToken, error)
	GetResetTokenByToken(token string) (*entity.PasswordResetToken, error)
	MarkTokenUsed(tx *gorm.DB, tokenID uuid.UUID) error
	UpdateUserPassword(tx *gorm.DB, userID uuid.UUID, newPassword string) error
}

type ResetPasswordRepository struct {
	db *gorm.DB
}

func NewResetPasswordRepository(db *gorm.DB) IResetPasswordRepository {
	return &ResetPasswordRepository{db: db}
}

func (r *ResetPasswordRepository) CreatePasswordResetToken(tx *gorm.DB, token *entity.PasswordResetToken) error {
	err := tx.Create(&token).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ResetPasswordRepository) GetResetToken(param model.UserForgotPasswordParam) (*entity.PasswordResetToken, error) {
	var resetToken entity.PasswordResetToken

	err := r.db.Where(&param).First(&resetToken).Error
	if err != nil {
		return nil, err
	}
	return &resetToken, nil
}

func (r *ResetPasswordRepository) GetResetTokenByToken(token string) (*entity.PasswordResetToken, error) {
	var resetToken entity.PasswordResetToken

	err := r.db.Where("token = ?", token).First(&resetToken).Error
	if err != nil {
		return nil, err
	}
	return &resetToken, nil
}

func (r *ResetPasswordRepository) MarkTokenUsed(tx *gorm.DB, tokenID uuid.UUID) error {
	err := tx.Model(&entity.PasswordResetToken{}).Where("id = ?", tokenID).Update("used", true).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ResetPasswordRepository) UpdateUserPassword(tx *gorm.DB, userID uuid.UUID, newPassword string) error {
	err := tx.Model(&entity.User{}).Where("user_id = ?", userID).Update("password", newPassword).Error
	if err != nil {
		return err
	}
	return nil
}
