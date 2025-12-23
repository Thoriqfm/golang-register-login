package rest

import (
	"fmt"
	"golang-register-login/model"
	"golang-register-login/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (r *Rest) RegisterUser(c *gin.Context) {
	var param model.UserRegisterParam

	err := c.ShouldBindJSON(&param)
	if err != nil {
		if validationError, ok := err.(validator.ValidationErrors); ok {
			errorsMessages := make([]string, 0)
			for _, e := range validationError {
				msg := fmt.Sprintf("Field '%s' error: %s", e.Field(), e.Tag())
				errorsMessages = append(errorsMessages, msg)
			}
			response.Error(c, http.StatusBadRequest, "invalid email format", fmt.Errorf("%v", errorsMessages))
			return
		}
		response.Error(c, http.StatusBadRequest, "failed to bind json", err)
		return
	}

	reps, err := r.service.UserService.RegisterUser(param)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to register user", err)
		return
	}

	response.Success(c, http.StatusOK, "user registered successfully", reps)
}
