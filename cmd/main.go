package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	runServer()
}

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warnln("Could not load environment variables")
	}
}

func runServer() {
	e := echo.New()
	fmt.Println("Server Running!")

	apiPathsWithoutTokenValidation := e.Group("/")
	apiPathsWithoutTokenValidation.GET("/healthCheck")

	apiPathsWithTokenValidation := e.Group("/v1")
	apiPathsWithTokenValidation.POST("/create")
	apiPathsWithTokenValidation.GET("/read")
	apiPathsWithTokenValidation.PATCH("/updateAnElement")
	apiPathsWithTokenValidation.DELETE("/deleteElement")
	apiPathsWithTokenValidation.PUT("/update")

	err := e.Start(":" + os.Getenv("CPORT"))
	if err != nil {
		logrus.Fatal(err)
	}
}
