package db

import (
	"errors"

	"github.com/sharmarajdaksh/basic-auth-microservice/models"
)

func performAutomigrations() error {

	if e := DB.AutoMigrate(&models.User{}); e != nil {
		return errors.New(e.Error())
	}

	return nil
}
