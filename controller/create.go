package controller

import (
	"context"
	"fmt"
	"net/http"
	"queries/db"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateHandler(c echo.Context) error {

	name := c.FormValue("name")
	lastname := c.FormValue("lastname")

	collection, err := db.GetBDCollection("prueba")
	if err != nil {
		return fmt.Errorf("err")
	}

	res, err := collection.InsertOne(context.TODO(), bson.D{{"nombre", name}, {"apellidos", lastname}})
	fmt.Println(res)

	return c.String(http.StatusOK, "ok")

}
