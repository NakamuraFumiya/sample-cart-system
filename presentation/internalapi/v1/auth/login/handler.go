package login

import (
	"fmt"
	"net/http"

	"github.com/fumiyanakamura/sample-cart-system/infrastructure/db/gorm"
	"github.com/labstack/echo/v4"
)

func Handler(c echo.Context) error {
	h, err := gorm.GetInstance()
	if err != nil {
		return err
	}

	type User struct {
		ID uint64
	}
	var t User
	h.Take(&t)

	return c.String(http.StatusOK, fmt.Sprintf("ID: %d", t.ID))
}
