package main

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/fumiyanakamura/sample-cart-system/userinterface/externalapi/v1/auth/login"
	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", hello)

	v1 := e.Group("/v1")
	v1.POST("/login", login.Handler)

	// Start server
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world")
}
