package model

import (
	"time"

	"github.com/Wandering-Digital/anthropos/cdt"

	"gorm.io/gorm"
)

type (
	User struct {
		ID       uint         `gorm:"primaryKey" json:"-"`
		UserName string       `gorm:"not null;type:varchar(100)" json:"user_name"`
		Password string       `gorm:"not null;type:varchar(100)" json:"-"`
		Type     cdt.UserType `gorm:"not null;type:varchar(20)" json:"user_type"`

		CreatedAt time.Time `gorm:"not null;type:timestamptz" json:"-"`
		UpdatedAt time.Time `gorm:"not null;type:timestamptz" json:"-"`
		DeletedAt gorm.DeletedAt
	}
)
