package router

import (
	"github.com/fumiyanakamura/sample-cart-system/presentation/internalapi/v1/auth/login"
	cartCreate "github.com/fumiyanakamura/sample-cart-system/presentation/internalapi/v1/cart/create"
	cartDetail "github.com/fumiyanakamura/sample-cart-system/presentation/internalapi/v1/cart/delete"
	cartDelete "github.com/fumiyanakamura/sample-cart-system/presentation/internalapi/v1/cart/detail"
	productDetail "github.com/fumiyanakamura/sample-cart-system/presentation/internalapi/v1/product/detail"
	purchaseCreate "github.com/fumiyanakamura/sample-cart-system/presentation/internalapi/v1/purchase/create"
	userCreate "github.com/fumiyanakamura/sample-cart-system/presentation/internalapi/v1/user/create"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	v1 := e.Group("/v1")

	// users
	v1.POST("/users", userCreate.NewController().Do)

	// login
	v1.POST("/login", login.Handler)

	// carts
	v1.GET("/carts/:id", cartDetail.Handler)
	v1.POST("/carts", cartCreate.Handler)
	v1.DELETE("/carts", cartDelete.Handler)

	// products
	v1.GET("/products/:id", productDetail.Handler)

	// purchase
	v1.POST("/purchase", purchaseCreate.Handler)
}
