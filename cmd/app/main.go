package main

import (
	echo "github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "education/docs"
)

// @title Post app with authorization
// @version 1.0
// @description Educational project

// @contact.email evgen.myroshnykov@gmail.com

// @host localhost:8080
// @BasePath /
// @schemes http

func main() {

	e := echo.New()

	routs(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
