package main

import (
	"golang-register-login/internal/handler/rest"
	"golang-register-login/internal/repository"
	"golang-register-login/internal/service"
	"golang-register-login/pkg/config"
	"golang-register-login/pkg/database/mysql"
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
	svc := service.NewService(repo)
	r := rest.NewRest(svc)
	r.MountEndPoint()
	r.Run()
}
