package entrypoints

import (
	"net/http"
	"time"

	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/entrypoints/utils"
	"github.com/labstack/echo/v4"
)

func (handler *UseCasesStruc) HealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, utils.BuildSuccessfulResponse(true, time.Now(), nil, nil))
}
