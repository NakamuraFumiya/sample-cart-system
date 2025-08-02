package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/fumiyanakamura/sample-cart-system/infrastructure/config/viper"
	"github.com/fumiyanakamura/sample-cart-system/infrastructure/db/gorm"
	"github.com/fumiyanakamura/sample-cart-system/presentation/router"
	"github.com/labstack/echo/v4"
)

func main() {
	// Echoインスタンスの生成
	e := echo.New()

	// ルートの登録
	router.RegisterRoutes(e)

	// DB接続確認(初期化)
	if _, err := gorm.GetInstance(); err != nil {
		panic(fmt.Sprintf("failed to initialize database connection: %v", err))
	}

	// APIサーバーの起動
	serverConfig := viper.LoadConfig().ServerConfig
	if err := e.Start(fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}
