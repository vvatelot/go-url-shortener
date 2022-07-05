package config

import (
	"os"

	"github.com/vvatelot/url-shortener/api/entities"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = os.Getenv("DATABASE_URI")

func Connect() error {
	var err error
	ormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	}

	switch os.Getenv("ENV") {
	case "dev":
		Database, err = gorm.Open(sqlite.Open(DATABASE_URI), ormConfig)
	case "prod":
		Database, err = gorm.Open(postgres.Open(DATABASE_URI), ormConfig)
	}

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entities.Link{})
	Database.AutoMigrate(&entities.Click{})

	return nil
}
