package main

import (
	factorymethod "github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/designPatterns/factoryMethod"
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/pkg/migration"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {

	err := migration.RunDBMigration()
	if err != nil {
		panic(err)
	}

	factorymethod.NewEchoAPIRestAdapter().RunServer()
}

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warnln("Could not load environment variables")
	}
}
