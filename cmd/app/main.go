package main

import (
	"golang-register-login/internal/handler/rest"
	"golang-register-login/internal/repository"
	"golang-register-login/internal/service"
	"golang-register-login/pkg/bcyrpt"
	"golang-register-login/pkg/config"
	"golang-register-login/pkg/database/mysql"
	"golang-register-login/pkg/jwt"
	"golang-register-login/pkg/middleware"
	"log"
)

func main() {
	config.LoadEnv()

	db, err := mysql.ConnectionDatabase()
	if err != nil {
		log.Fatal(err)
	}

	err = mysql.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	bcrypt := bcyrpt.Init()
	jwtAuth := jwt.Init()
	svc := service.NewService(repo, bcrypt, jwtAuth)
	middleware := middleware.Init(svc, jwtAuth)
	r := rest.NewRest(svc, middleware)
	r.MountEndPoint()
	r.Run()
}
