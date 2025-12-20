package mysql

import (
	"golang-register-login/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.Role{},
		&entity.User{},
	)

	if err != nil {
		return err
	}

	return nil
}
