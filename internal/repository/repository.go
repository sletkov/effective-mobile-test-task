package repository

import (
	"context"

	"github.com/sletkov/effective-mobile-test-task/internal/model"
)

type UserRepository interface {
	Get(ctx context.Context, userFilter *model.UserFilter) ([]model.User, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, id uint, u *model.User) error
	Create(ctx context.Context, u *model.User) (uint, error)
	GetUserById(ctx context.Context, id uint) (*model.User, error)
}
