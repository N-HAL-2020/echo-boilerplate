package validator_test

import (
	"net/http/httptest"
	"testing"

	"github.com/N-HAL-2020/echo-boilerplate/cmd/echo-boilerplate/validator"
	"github.com/labstack/echo/v4"
)

type testRequest struct {
	Name string `json:"name" validate:"required"`
}

func TestApplyValidator(t *testing.T) {
	t.Parallel()

	e := echo.New()
	validator.ApplyValidator(e)

	req := httptest.NewRequest(echo.POST, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")

	cases := []struct {
		name        string
		req         interface{}
		expectedErr bool
	}{
		{
			name: "valid",
			req: testRequest{
				Name: "test",
			},
			expectedErr: false,
		},
		{
			name:        "invalid",
			req:         testRequest{},
			expectedErr: true,
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if err := c.Bind(&tc.req); err != nil {
				t.Fatal(err)
			}
			if err := c.Validate(tc.req); (err != nil) != tc.expectedErr {
				t.Fatal(err)
			}
		})
	}
}
