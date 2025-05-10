package main

import (
	"errors"
	"log/slog"
	"net/http"

	v1 "github.com/fumiyanakamura/sample-cart-system/userinterface/externalapi/v1"
	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	v1.RegisterRoutes(e)

	// Start server
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}
