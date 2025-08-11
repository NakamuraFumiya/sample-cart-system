package repository

import (
	"context"

	"github.com/fumiyanakamura/sample-cart-system/core/domain/model"
)

type UserRepository interface {
	Create(context.Context, model.User) (*model.User, error)
}
