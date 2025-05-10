package v1

import (
	"github.com/fumiyanakamura/sample-cart-system/userinterface/externalapi/v1/auth/login"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	v1 := e.Group("/v1")
	v1.POST("/login", login.Handler)
}
