package entrypoints

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/dtos"
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/entrypoints/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (handler *UseCasesStruc) CreatePersonHandler(c echo.Context) error {
	var personDTO *dtos.PersonDTO
	err := c.Bind(&personDTO)
	if err != nil {
		logrus.Error("there is an error binding body", err)
		return c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(false, time.Now(), err.Error()))
	}
	err = personDTO.Validate()
	if err != nil {
		logrus.Error("there is an error parsing id", err)
		return c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(false, time.Now(), err.Error()))
	}

	result, err := handler.usecasesStruct.CreatePersonUseCase(personDTO)
	if err != nil {
		logrus.Error("error creating person ", err)
		return c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(false, time.Now(), err.Error()))
	}
	return c.JSON(http.StatusOK, utils.BuildSuccessfulResponse(true, time.Now(), fmt.Sprint("person created id: ", result), nil))
}
