package entrypoints

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/dtos"
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/entrypoints/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (handler *UseCasesStruc) GetAllPersonInformationHandler(c echo.Context) error {
	limit, err := strconv.ParseInt(c.QueryParam("limit"), 10, 64)
	if err != nil {
		logrus.Error("there is an error parsing limit", err)
		return c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(false, time.Now(), err.Error()))
	}
	offset, err := strconv.ParseInt(c.QueryParam("offset"), 10, 64)
	if err != nil {
		logrus.Error("there is an error parsing offset", err)
		return c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(false, time.Now(), err.Error()))
	}
	requestMetadata := dtos.RequestInformationMetadata{
		Limit:  limit,
		Offset: offset,
	}
	result, pagination, err := handler.usecasesStruct.GetAllPersonInformationUseCase(requestMetadata)
	if err != nil {
		logrus.Error("error getting all people data ", err)
		return c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(false, time.Now(), err.Error()))
	}
	return c.JSON(http.StatusOK, utils.BuildSuccessfulResponse(true, time.Now(), result, pagination))
}
