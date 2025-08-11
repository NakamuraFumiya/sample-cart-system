package create

import (
	"errors"
	"regexp"

	api "github.com/fumiyanakamura/sample-cart-system/openapigen"
	"github.com/fumiyanakamura/sample-cart-system/presentation/util"
	"github.com/labstack/echo/v4"
)

func toRequest(ctx echo.Context) (*api.PostV1UsersJSONBody, error) {
	req, err := util.Bind(ctx, api.PostV1UsersJSONBody{})
	if err != nil {
		return nil, err
	}
	if err := util.Validate(ctx, *req); err != nil {
		return nil, err
	}
	// ユーザーコードの検証
	usercodeMatched, err := regexp.MatchString("^[a-za-z0-9_]+$", req.Usercode)
	if err != nil {
		return nil, err
	}
	if !usercodeMatched {
		return nil, errors.New("invalid usercode")
	}
	// パスワードの検証
	passwordMatched, err := regexp.MatchString("^[\x21-\x7e]{8,64}$", req.Password)
	if err != nil {
		return nil, err
	}
	if !passwordMatched {
		return nil, errors.New("invalid password")
	}

	return req, nil
}
