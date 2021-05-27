package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"queries/db"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func ListHandler(c echo.Context) error {
	collection, err := db.GetBDCollection("user")
	if err != nil {
		return fmt.Errorf("err")
	}

	list, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(list)

	return c.String(http.StatusOK, "ok")

}
