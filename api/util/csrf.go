package util

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// GetCSRFToken function
func GetCSRFToken(c echo.Context) error {
	token := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
	return c.String(http.StatusOK, token)
}
