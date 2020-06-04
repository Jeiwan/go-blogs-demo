package web

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func (s Server) badRequest(c echo.Context, err error) error {
	logrus.Errorf("bad request: %+v", err)
	return c.JSON(http.StatusBadGateway, map[string]string{"error": "bad request"})
}

func (s Server) notFound(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
}

func (s Server) internalServerError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
}
