package factorymethod

import (
	"os"

	echoapirest "github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/entrypoints"
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/entrypoints/api/interfaces"
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/entrypoints/api/routes"
)

func NewEchoAPIRestAdapter() interfaces.IAPIRest {
	return echoapirest.NewAPIRestServer(
		os.Getenv("CPORT"),
		os.Getenv("API_REST_BASE_URL"),
		routes.ApiRoutes,
		routes.ApiJWTRoutes,
	)
}
