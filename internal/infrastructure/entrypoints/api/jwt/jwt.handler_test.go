package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestErrorHandler(t *testing.T) {
	t.Run("returns token validation errors", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/v1/health", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := NewJWTHandler("test")
		if assert.NoError(t, h.ErrorHandler(c, errors.New(""))) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	t.Run("returns token validation errors, but error has message-code structure", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/v1/health", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := NewJWTHandler("test")
		if assert.NoError(t, h.ErrorHandler(c, errors.New("code=401 message=error"))) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})
}
