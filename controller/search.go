package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mikamikuh/goannotator/model/mongo"
)

type SearchRequest struct {
}

func SearchHandler(c echo.Context) error {
	s, err := mongo.NewSession("database", "annotator", "annotator")

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	_, err = s.GetAnnotations("someuri")

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "success")
}
