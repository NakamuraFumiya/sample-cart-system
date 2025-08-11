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
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	fmt.Println(req.Usercode)
	fmt.Println(req.Password)

	return ctx.String(http.StatusOK, "user create success")
}
