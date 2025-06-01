package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/fumiyanakamura/sample-cart-system/infrastructure/config/viper"
	"github.com/fumiyanakamura/sample-cart-system/infrastructure/db/gorm"
	v1 "github.com/fumiyanakamura/sample-cart-system/userinterface/externalapi/v1"
	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	v1.RegisterRoutes(e)

	// DB接続確認(初期化)
	if _, err := gorm.GetInstance(); err != nil {
		panic(fmt.Sprintf("failed to initialize database connection: %v", err))
	}

	// Start server
	serverConfig := viper.LoadConfig().ServerConfig
	if err := e.Start(fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}
