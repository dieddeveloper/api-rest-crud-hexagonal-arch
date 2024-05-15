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
	adapters := adapters.NewPersonAdapter(dbAdapter)

	//services
	services := service.NewServicesConstructor(adapters)

	//usecases
	usecases := usecase.NewUseCaseConstructor(services)

	//handlers
	handlers := entrypoints.NewHandlerConstructor(usecases)

	group.GET("v1/health", handlers.HealthHandler)

}

func ApiJWTRoutes(group *echo.Group) {

	dbAdapter, err := singleton.NewSingletonPostgresAdapter()
	if err != nil {
		panic(err)
	}

	//adapters
	adapters := adapters.NewPersonAdapter(dbAdapter)

	//services
	services := service.NewServicesConstructor(adapters)

	//usecases
	usecases := usecase.NewUseCaseConstructor(services)

	//handlers
	handlers := entrypoints.NewHandlerConstructor(usecases)

	group.GET("v1/person/all", handlers.GetAllPersonInformationHandler)
	group.GET("v1/person/get/:cardNumber", handlers.GetPersonByIDHandler)
	group.POST("v1/person/create", handlers.GetPersonByIDHandler)
	/*group.PATCH("/updateAnElement", handlers.GetPersonByIDHandler)

	apiPathsWithTokenValidation.DELETE("/deleteElement")
	apiPathsWithTokenValidation.PUT("/update")*/

}
