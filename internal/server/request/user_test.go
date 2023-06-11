package request

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/tsusowake/go.server/util/echoutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewCreateUser(t *testing.T) {
	e := echo.New()
	echoutil.UseCustomValidator(e)

	t.Run("Valid requests", func(t *testing.T) {
		body := `{"password":"password-1", "email":"test-1@email.com"}`
		req := httptest.NewRequest(
			http.MethodPost,
			"/",
			bytes.NewReader([]byte(body)),
		)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		got, err := NewCreateUser(e.NewContext(req, rec))
		assert.NoError(t, err)
		assert.Equal(t, got, &CreateUser{
			Password: "password-1",
			Email:    "test-1@email.com",
		})
	})
	t.Run("Password is empty", func(t *testing.T) {
		body := `{"password":"", "email":"test-1@email.com"}`
		req := httptest.NewRequest(
			http.MethodPost,
			"/",
			bytes.NewReader([]byte(body)),
		)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		_, err := NewCreateUser(e.NewContext(req, rec))
		assert.Error(t, err)
	})
}
