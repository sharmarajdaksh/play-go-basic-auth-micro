package db

import (
	"fmt"

	"github.com/sharmarajdaksh/basic-auth-microservice/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global database connection object
var DB *gorm.DB

// InitializeDB initializes the global DB object and make automigrations
func InitializeDB() error {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.C.Database.Postgres.PostgresUsername,
		config.C.Database.Postgres.PostgresPassword,
		config.C.Database.Postgres.PostgresDatabaseName,
		config.C.Database.Postgres.PostgresHost,
		config.C.Database.Postgres.PostgresConnectionPort,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if e := performAutomigrations(); e != nil {
		return e
	}

	return nil
}
