package model

import (
	"errors"
	"regexp"

	"github.com/fumiyanakamura/sample-cart-system/core/domain/identifier"
)

const UserUserCodeRegex = "^[a-za-z0-9_]+$"
const UserPasswordRegex = "^[\x21-\x7e]{8,64}$"

type User struct {
	id       identifier.UserID
	userCode string
	password string
}

func NewUser(
	userCode string,
	password string,
) (*User, error) {
	// ユーザーコードの検証
	userCodeMatched, err := regexp.MatchString(UserUserCodeRegex, userCode)
	if err != nil {
		return nil, err
	}
	if !userCodeMatched {
		return nil, errors.New("invalid userCode")
	}
	// パスワードの検証
	passwordMatched, err := regexp.MatchString(UserPasswordRegex, password)
	if err != nil {
		return nil, err
	}
	if !passwordMatched {
		return nil, errors.New("invalid password")
	}

	return &User{
		userCode: userCode,
		password: password,
	}, nil
}

func (u *User) ID() identifier.UserID { return u.id }
func (u *User) UserCode() string      { return u.userCode }

func (u *User) SetUserCode(userCode string) { u.userCode = userCode }
