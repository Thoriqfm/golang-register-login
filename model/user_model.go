package model

import "github.com/google/uuid"

type UserParam struct {
	UserID   uuid.UUID `json:"-"`
	Username string    `json:"-"`
	Email    string    `json:"-"`
}
type UserRegisterParam struct {
	Username        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type UserRegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserLoginParam struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	Token  string `json:"token"`
	RoleID int    `json:"role_id"`
}

type UserForgotPasswordParam struct {
	Email string `json:"email" binding:"required,email"`
}

type UserForgotPasswordResponse struct {
	Token              string `json:"token"`
	NewPassword        string `json:"new_password" binding:"required"`
	ConfirmNewPassword string `json:"confirm_new_password" binding:"required"`
}
