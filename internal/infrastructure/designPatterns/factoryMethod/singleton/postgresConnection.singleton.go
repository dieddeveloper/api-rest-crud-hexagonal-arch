package singleton

import (
	"os"

	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/adapters/database"
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/pkg"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB
	dbError bool
)

func NewSingletonPostgresAdapter() (*gorm.DB, error) {
	var err error
	if db == nil {
		err = connectToDB()
		if err != nil {
			dbError = true
			return nil, err
		}
	}

	if dbError {
		err = connectToDB()
		if err != nil {
			dbError = true
			return nil, err
		}
	}
	dbError = false
	campaignAdapter := database.NewDatabasePostgresConnection(db)
	return campaignAdapter, nil
}

func connectToDB() error {
	var err error
	db, err = pkg.NewConn(&pkg.DBConnDto{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	})

	return err
}
