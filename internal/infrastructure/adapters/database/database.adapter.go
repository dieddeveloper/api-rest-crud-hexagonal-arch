package database

import "gorm.io/gorm"

func NewDatabasePostgresConnection(dbConnection *gorm.DB) *gorm.DB {
	return dbConnection
}
