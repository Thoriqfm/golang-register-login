package rest

import (
	"golang-register-login/model"
	"golang-register-login/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST
func (r *Rest) ForgotPassword(c *gin.Context) {
	var param model.UserForgotPasswordParam

	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind json", err)
		return
	}

	err = r.service.ResetPasswordService.RequestResetPassword(param)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to request reset password", err)
		return
	}

	response.Success(c, http.StatusOK, "reset password request successfully", nil)
}

// GET
func (r *Rest) VerifyResetToken(c *gin.Context) {
	token := c.Query("token")

	if token == "" {
		response.Error(c, http.StatusBadRequest, "token is required", nil)
		return
	}

	valid, err := r.service.ResetPasswordService.VerifyResetToken(token)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "token is valid", gin.H{
		"valid": valid,
	})
}

// POST
func (r *Rest) ResetPassword(c *gin.Context) {
	var param model.UserForgotPasswordResponse

	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind json", err)
		return
	}

	err = r.service.ResetPasswordService.ResetPassword(param.Token, param.NewPassword, param.ConfirmNewPassword)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "password reset successfully", nil)
}
