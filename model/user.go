package model

import (
	"time"

	"github.com/Wandering-Digital/anthropos/cdt"

	"gorm.io/gorm"
)

type (
	User struct {
		ID          uint            `gorm:"primaryKey" json:"-"`
		Email       string          `gorm:"not null;type:varchar(100)" json:"email"`
		Password    string          `gorm:"not null;type:varchar(100)" json:"-"`
		AccountType cdt.AccountType `gorm:"not null;type:varchar(20)" json:"account_type"`

		CreatedAt time.Time `gorm:"not null;type:timestamptz" json:"-"`
		UpdatedAt time.Time `gorm:"not null;type:timestamptz" json:"-"`
		DeletedAt gorm.DeletedAt
	}
)

func (u *User) TableName() string {
	return "users"
}
