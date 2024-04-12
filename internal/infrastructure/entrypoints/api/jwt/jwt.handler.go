package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/entrypoints/utils"
	"github.com/labstack/echo/v4"
)

type JWTHandler struct {
	Secret string
}

func NewJWTHandler(secret string) *JWTHandler {
	return &JWTHandler{
		Secret: secret,
	}
}

func (a *JWTHandler) ErrorHandler(c echo.Context, err error) error {
	var errorSlice string
	if strings.Contains(err.Error(), "code=") {
		errorSlice = strings.Split(err.Error(), "message=")[1]
	} else {
		errorSlice = err.Error()
	}
	response := utils.ResponseError{
		Success:      false,
		Timestamp:    time.Now(),
		ErrorMessage: errorSlice,
	}
	return c.JSON(http.StatusUnauthorized, response)
}
