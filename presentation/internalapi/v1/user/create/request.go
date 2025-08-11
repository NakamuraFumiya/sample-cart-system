package create

import (
	"errors"
	"regexp"

	"github.com/fumiyanakamura/sample-cart-system/core/domain/model"
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
	userCodeMatched, err := regexp.MatchString(model.UserUserCodeRegex, req.Usercode)
	if err != nil {
		return nil, err
	}
	if !userCodeMatched {
		return nil, errors.New("invalid userCode")
	}
	// パスワードの検証
	passwordMatched, err := regexp.MatchString(model.UserPasswordRegex, req.Password)
	if err != nil {
		return nil, err
	}
	if !passwordMatched {
		return nil, errors.New("invalid password")
	}

	return req, nil
}
