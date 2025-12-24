package service

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"golang-register-login/entity"
	"golang-register-login/internal/repository"
	"golang-register-login/model"
	"golang-register-login/pkg/bcyrpt"
	"golang-register-login/pkg/database/mysql"
	"golang-register-login/pkg/email"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IResetPasswordService interface {
	RequestResetPassword(param model.UserForgotPasswordParam) error
	VerifyResetToken(token string) (bool, error)
	ResetPassword(token, newPassword, ConfirmPassword string) error
}

type ResetPasswordService struct {
	db             *gorm.DB
	UserRepository repository.IUserRepository
	repo           repository.IResetPasswordRepository
	bycrypt        bcyrpt.Interface
	email          email.Interface
}

func NewResetPasswordService(userRepo repository.IUserRepository, repo repository.IResetPasswordRepository, bycrypt bcyrpt.Interface, emailService email.Interface) IResetPasswordService {
	return &ResetPasswordService{
		db:             mysql.Connection,
		UserRepository: userRepo,
		repo:           repo,
		bycrypt:        bycrypt,
		email:          emailService,
	}
}

func generateRandomToken(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (s *ResetPasswordService) RequestResetPassword(param model.UserForgotPasswordParam) error {
	tx := s.db.Begin()
	defer tx.Rollback()

	user, err := s.UserRepository.GetUser(model.UserParam{
		Email: param.Email,
	})

	if err != nil {
		return nil
	}

	// Generate random token
	token, err := generateRandomToken(32)
	if err != nil {
		return errors.New("failed to generate token")
	}

	// Make reset token entity
	tokenID, _ := uuid.NewUUID()
	resetToken := &entity.PasswordResetToken{
		ID:        tokenID,
		UserID:    user.UserID,
		Token:     token,
		ExpiredAt: time.Now().Add(1 * time.Hour),
		Used:      false,
	}

	// Save to database
	err = s.repo.CreatePasswordResetToken(tx, resetToken)
	if err != nil {
		return errors.New("failed to create reset token")
	}

	// Send email
	err = s.email.SendRestPasswordEmail(param.Email, token)
	if err != nil {
		return errors.New("failed to send email")
	}

	// Commit
	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil

}

func (s *ResetPasswordService) VerifyResetToken(token string) (bool, error) {
	resetToken, err := s.repo.GetResetTokenByToken(token)
	if err != nil {
		return false, errors.New("invalid token")
	}

	if resetToken.Used {
		return false, errors.New("token already used")
	}

	if time.Now().After(resetToken.ExpiredAt) {
		return false, errors.New("token already expired")
	}

	return true, nil
}

func (s *ResetPasswordService) ResetPassword(token, newPassword, ConfirmPassword string) error {
	tx := s.db.Begin()
	defer tx.Rollback()

	// validation password
	if newPassword != ConfirmPassword {
		return errors.New("password not match")
	}

	// find token from database
	resetToken, err := s.repo.GetResetTokenByToken(token)
	if err != nil {
		return errors.New("invalid token")
	}

	// validation (is token already used)
	if resetToken.Used {
		return errors.New("token already used")
	}

	// validation (is token expired)
	if time.Now().After(resetToken.ExpiredAt) {
		return errors.New("token already expired")
	}

	user, err := s.UserRepository.GetUser(model.UserParam{
		UserID: resetToken.UserID,
	})
	if err != nil {
		return errors.New("user not found")
	}

	// Compare password (must different from last password)
	err = s.bycrypt.CompareHashPassword(user.Password, newPassword)
	if err == nil {
		return errors.New("password must different from last password")
	}

	// Hash new password
	newHashPassword, err := s.bycrypt.GenerateHashPassword(newPassword)
	if err != nil {
		return errors.New("failed to hash password")
	}

	// Update user password
	err = s.repo.UpdateUserPassword(tx, user.UserID, newHashPassword)
	if err != nil {
		return errors.New("failed to update password")
	}

	// Mark token to used token
	err = s.repo.MarkTokenUsed(tx, resetToken.ID)
	if err != nil {
		return errors.New("failed to update token")
	}

	// Commit
	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}
