package domain

import (
	"context"

	"github.com/Wandering-Digital/anthropos/criteria"
	"github.com/Wandering-Digital/anthropos/model"
)

type UserUseCase interface {
	Create(ctx context.Context, ctr *criteria.CreateUser) (*model.User, error)
}

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
}
