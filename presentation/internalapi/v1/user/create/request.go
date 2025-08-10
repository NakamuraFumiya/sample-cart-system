package create

import (
	api "github.com/fumiyanakamura/sample-cart-system/openapigen"
	"github.com/fumiyanakamura/sample-cart-system/presentation/util"
	"github.com/labstack/echo/v4"
)

type Request struct {
	api.PostV1UsersJSONBody
}

func toRequest(ctx echo.Context) (*Request, error) {
	req, err := util.Bind(ctx, Request{})
	if err != nil {
		return nil, err
	}
	if err := util.Validate(ctx, *req); err != nil {
		return nil, err
	}
	return req, nil
}
