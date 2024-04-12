package echoapirest

import (
	"strings"

	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/global"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type APIRestServer struct {
	echoInstance   *echo.Echo
	globalGroup    *echo.Group
	globalJWTGroup func(*echo.Group)
	port           string
	routes         func(*echo.Group)
}

func NewAPIRestServer(port, globalPrefix string, routes func(*echo.Group), jwtGroup func(*echo.Group)) *APIRestServer {
	e := echo.New()
	return &APIRestServer{
		echoInstance:   e,
		globalGroup:    e.Group(globalPrefix),
		port:           port,
		routes:         routes,
		globalJWTGroup: jwtGroup,
	}
}

func (api *APIRestServer) RunServer() {

	api.echoInstance.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		CustomTimeFormat: global.DateAndHourFormat,
		Format:           "[${time_custom}] - method=${method}, uri=${uri}, status=${status}\n",
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Path(), "/v1/healthz") || c.Path() == "" {
				return true
			}
			return false
		},
	}))

	api.routes(api.globalGroup)
	api.echoInstance.Logger.Fatal(api.echoInstance.Start(":" + api.port))
}
