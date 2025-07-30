package delete

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Handler(c echo.Context) error {
	return c.String(http.StatusOK, "Cart Delete Success")
}
