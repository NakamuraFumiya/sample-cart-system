package create

import (
	api "github.com/fumiyanakamura/sample-cart-system/openapigen"
	"github.com/labstack/echo/v4"
)

type Request struct {
	api.PostV1UsersJSONBody
}

func toRequest(ctx echo.Context) (*Request, error) {
	req, err := Bind(ctx)
	if err != nil {
		return nil, err
	}
	if err := Validate(ctx, req); err != nil {
		return nil, err
	}
	return req, nil
}

// TODO: 共通処理に移動させる
func Bind(ctx echo.Context) (*Request, error) {
	req := new(Request)
	if err := ctx.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func Validate(ctx echo.Context, req *Request) error {
	return ctx.Validate(req)
}
