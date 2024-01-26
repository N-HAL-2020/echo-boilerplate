//go:generate mockgen -source=$GOFILE -destination=mock/mock_$GOFILE -package=mock
package healthcheck

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	HealthCheck(c echo.Context) error
}

type handlerImpl struct{}

func NewHandler() Handler {
	return &handlerImpl{}
}

func (h *handlerImpl) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{Status: "OK"})
}
