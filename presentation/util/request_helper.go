package util

import "github.com/labstack/echo/v4"

func Bind[T any](ctx echo.Context, req T) (*T, error) {
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func Validate[T any](ctx echo.Context, req T) error {
	return ctx.Validate(&req)
}
