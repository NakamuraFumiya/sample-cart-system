package repositoryimpl

import (
	"context"

	"github.com/fumiyanakamura/sample-cart-system/core/domain/model"
	"github.com/fumiyanakamura/sample-cart-system/core/domain/repository"
	gormHandler "github.com/fumiyanakamura/sample-cart-system/infrastructure/db/gorm"
	"gorm.io/gorm"
)

var _ repository.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	handler *gorm.DB
}

func NewUserRepository() *UserRepository {
	h, _ := gormHandler.GetInstance()
	return &UserRepository{
		handler: h,
	}
}

// TODO: 実装とテスト追加
func (r *UserRepository) Create(context.Context, model.User) (*model.User, error) {
	return nil, nil
}
