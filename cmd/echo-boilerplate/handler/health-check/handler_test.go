package healthcheck_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	healthcheck "github.com/N-HAL-2020/echo-boilerplate/cmd/echo-boilerplate/handler/health-check"
	"github.com/labstack/echo/v4"
)

func TestHandler_HealthCheck(t *testing.T) {
	t.Parallel()

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/health-check", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/health-check")

	handler := healthcheck.NewHandler()

	if err := handler.HealthCheck(c); err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected %d, got %d", http.StatusOK, rec.Code)
	}
}
