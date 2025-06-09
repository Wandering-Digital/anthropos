package usecase

import (
	"context"

	"github.com/Wandering-Digital/anthropos/criteria"
	"github.com/Wandering-Digital/anthropos/domain"
	"github.com/Wandering-Digital/anthropos/model"
)

type User struct {
	userRepo domain.UserRepository
}

func NewUser(userRepo domain.UserRepository) domain.UserUseCase {
	return &User{
		userRepo: userRepo,
	}
}

func (u *User) Create(ctx context.Context, ctr *criteria.CreateUser) (*model.User, error) {
	user := model.User{
		UserName: ctr.UserName,
	}

	if err := u.userRepo.Create(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
