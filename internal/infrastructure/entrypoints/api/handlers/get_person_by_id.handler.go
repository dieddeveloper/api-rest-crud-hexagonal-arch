package entrypoints

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/entrypoints/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (handler *UseCasesStruc) GetPersonByIDHandler(c echo.Context) error {
	personID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logrus.Error("there is an error parsing id", err)
		return c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(false, time.Now(), err.Error()))
	}

	result, err := handler.usecasesStruct.GetPersonByIDUseCase(personID)
	if err != nil {
		logrus.Error("error getting person data ", err)
		return c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(false, time.Now(), err.Error()))
	}
	return c.JSON(http.StatusOK, utils.BuildSuccessfulResponse(true, time.Now(), result, nil))
}
