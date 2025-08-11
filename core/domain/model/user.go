package model

import "github.com/fumiyanakamura/sample-cart-system/core/domain/identifier"

type User struct {
	id       identifier.UserID
	name     string
	password string
}

func NewUser() (*User, error) {
	return nil, nil
}

func (u *User) ID() identifier.UserID { return u.id }
func (u *User) Name() string          { return u.name }

func (u *User) SetName(name string) { u.name = name }
