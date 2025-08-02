package create

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Do(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "user create success")
}
