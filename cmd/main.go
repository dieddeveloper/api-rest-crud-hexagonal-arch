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

	err := e.Start(":" + os.Getenv("CPORT"))
	if err != nil {
		logrus.Fatal(err)
	}
}
