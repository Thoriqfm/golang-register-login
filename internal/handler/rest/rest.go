package rest

import (
	"fmt"
	"golang-register-login/internal/service"
	"golang-register-login/pkg/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.Interface
}

func NewRest(service *service.Service, middleware middleware.Interface) *Rest {
	return &Rest{
		router:     gin.Default(),
		service:    service,
		middleware: middleware,
	}
}

func (r *Rest) MountEndPoint() {
	baseURL := r.router.Group("/api")

	// Auth endpoint
	auth := baseURL.Group("/auth")
	auth.POST("/register", r.RegisterUser)
	auth.POST("/login", r.LoginUser)

	// Reset password endpoint
	auth.POST("/forgot-password", r.ForgotPassword)
	auth.GET("/verify-reset-token", r.VerifyResetToken)
	auth.POST("/reset-password", r.ResetPassword)
}

func (r *Rest) Run() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	r.router.Run(fmt.Sprintf("%s:%s", addr, port))
}
