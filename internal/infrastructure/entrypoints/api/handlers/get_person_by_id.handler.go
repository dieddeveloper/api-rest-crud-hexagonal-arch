package entrypoints

import (
	"net/http"
	"time"

	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/constants"
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/entrypoints/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (handler *UseCasesStruc) GetPersonByIDHandler(c echo.Context) error {
	cardNumber := c.Param("cardNumber")
	if cardNumber != "" {
		logrus.Error("there is an error parsing id", cardNumber)
		return c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(false, time.Now(), `cardNumber`+constants.FieldRequiredMessage))
	}

	result, err := handler.usecasesStruct.GetPersonByIDUseCase(cardNumber)
	if err != nil {
		logrus.Error("error getting person data ", err)
		return c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(false, time.Now(), err.Error()))
	}
	return c.JSON(http.StatusOK, utils.BuildSuccessfulResponse(true, time.Now(), result, nil))
}
