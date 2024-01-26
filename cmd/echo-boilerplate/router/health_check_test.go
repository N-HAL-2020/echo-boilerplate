package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/N-HAL-2020/echo-boilerplate/cmd/echo-boilerplate/handler/health-check/mock"
	"github.com/N-HAL-2020/echo-boilerplate/cmd/echo-boilerplate/router"
	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
)

func TestHealthCheck(t *testing.T) {
	t.Parallel()

	e := echo.New()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	handler := mock.NewMockHandler(mockCtl)
	handler.EXPECT().HealthCheck(gomock.Any()).DoAndReturn(func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	router := router.NewHealthCheckRouter(
		handler,
	)
	router.Register(e)

	req := httptest.NewRequest(http.MethodGet, "/health-check", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected %d, got %d", http.StatusOK, rec.Code)
	}
}
