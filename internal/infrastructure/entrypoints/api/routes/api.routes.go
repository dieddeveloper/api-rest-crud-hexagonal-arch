package routes

import (
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/application/usecase"
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/service"
	adapters "github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/adapters/database/person"
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/designPatterns/factoryMethod/singleton"
	entrypoints "github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/entrypoints/api/handlers"
	"github.com/labstack/echo/v4"
)

func ApiRoutes(group *echo.Group) {

	dbAdapter, err := singleton.NewSingletonPostgresAdapter()
	if err != nil {
		panic(err)
	}

	//adapters
	personAdapter := adapters.NewPersonAdapter(dbAdapter)

	//services
	personService := service.NewServicesConstructor(personAdapter)

	//usecases
	personUsecases := usecase.NewUseCaseConstructor(personService)

	//handlers
	personHandler := entrypoints.NewHandlerConstructor(personUsecases)

	group.GET("v1/getAllInformation", personHandler.GetAllPersonInformationHandler)

	/*
		group.GET("/healthCheck")

		apiPathsWithTokenValidation := e.Group("/v1")
		apiPathsWithTokenValidation.POST("/create")
		apiPathsWithTokenValidation.GET("/read")
		apiPathsWithTokenValidation.PATCH("/updateAnElement")
		apiPathsWithTokenValidation.DELETE("/deleteElement")
		apiPathsWithTokenValidation.PUT("/update")*/

}