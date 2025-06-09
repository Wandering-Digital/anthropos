package postgres

import (
	"context"

	"github.com/Wandering-Digital/anthropos/domain"
	"github.com/Wandering-Digital/anthropos/internal/conn"
	"github.com/Wandering-Digital/anthropos/model"
)

type User struct {
	*conn.DB
}

func NewUser(db *conn.DB) domain.UserRepository {
	return &User{
		DB: db,
	}
}

func (u *User) Create(ctx context.Context, user *model.User) error {
	if err := u.GormDB.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}

	return nil
}
