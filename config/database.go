package config

import (
	"github.com/vvatelot/url-shortener/api/entities"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() error {
	var err error
	ormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	}

	switch ENV.Env {
	case "dev":
		Database, err = gorm.Open(sqlite.Open(ENV.DatabaseUri), ormConfig)
	case "prod":
		Database, err = gorm.Open(postgres.Open(ENV.DatabaseUri), ormConfig)
	}

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entities.Link{})
	Database.AutoMigrate(&entities.Click{})

	return nil
}
