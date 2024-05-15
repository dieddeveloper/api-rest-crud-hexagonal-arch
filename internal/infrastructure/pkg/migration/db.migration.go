package migration

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/pkg"
	"github.com/sirupsen/logrus"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunDBMigration() error {

	_, err := pkg.NewConn(&pkg.DBConnDto{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	})
	if err != nil {
		logrus.Errorln("RunDBMigration", err)
		return err
	}

	connect := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("postgres", connect)

	if err != nil {
		logrus.Errorln("Error while opening connection", err)
		panic(err.Error() + " Error while opening connection")
	}

	defer db.Close()
	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		logrus.Errorln("Error while opening connection", err)
		panic(err.Error() + " Error while opening connection")
	}

	migrationFile, err := migrate.NewWithDatabaseInstance("file://db/migration/DDL", os.Getenv("DB_NAME"), driver)

	if err != nil {
		logrus.Errorln("Error while opening connection", err)
		panic(err.Error() + " Error while opening connection")
	}

	err = migrationFile.Up()
	if err != nil {
		if err.Error() == "no change" {
			logrus.Errorln("RunDBMigration", "There were no changes in the database migration")
		} else {
			panic(err.Error() + " Error while opening connection")
		}
	}

	defer driver.Close()
	logrus.Errorln("RunDBMigration", "Database migration carried out successfully")
	return nil
}
