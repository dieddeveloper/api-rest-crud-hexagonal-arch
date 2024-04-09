package entrypoints

import (
	"net/http"
	"time"

	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/entrypoints/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (handler *UseCasesStruc) GetAllPersonInformationHandler(c echo.Context) error {
	result, err := handler.usecasesStruct.GetAllPersonInformationUseCase()
	if err != nil {
		logrus.Error("error getting all people data ", err)
		return c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(false, time.Now(), err.Error()))
	}
	return c.JSON(http.StatusOK, utils.BuildSuccessfulResponse(true, time.Now(), result, nil))
}
