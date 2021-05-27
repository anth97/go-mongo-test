package main

import (
	"fmt"
	"net/http"

	hand "queries/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Hola mundo")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hola mundo")
	})

	e.POST("/", hand.CreateHandler)
	// e.GET("/get", controller.ListHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
