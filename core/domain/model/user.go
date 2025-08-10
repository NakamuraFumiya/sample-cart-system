package model

import "github.com/fumiyanakamura/sample-cart-system/core/domain/identifier"

type User struct {
	id   identifier.UserID
	name string
}
