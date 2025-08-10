package create

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Do(ctx echo.Context) error {
	req, err := toRequest(ctx)
	if err != nil {
		return err
	}
	fmt.Println(req)

	return ctx.String(http.StatusOK, "user create success")
}
