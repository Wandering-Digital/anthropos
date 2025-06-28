package model

import (
	"time"

	"github.com/Wandering-Digital/anthropos/cdt"

	"gorm.io/gorm"
)

type (
	User struct {
		ID          uint              `gorm:"primaryKey" json:"id"`
		Email       string            `gorm:"not null;type:varchar(100)" json:"email"`
		Password    string            `gorm:"not null;type:varchar(100)" json:"-"`
		AccountType cdt.AccountType   `gorm:"not null;type:varchar(20)" json:"account_type"`
		Status      cdt.AccountStatus `gorm:"not null;type:varchar(30)" json:"status"`

		CreatedAt time.Time `gorm:"not null;type:timestamptz" json:"-"`
		UpdatedAt time.Time `gorm:"not null;type:timestamptz" json:"-"`
		DeletedAt gorm.DeletedAt
	}

	Customer struct {
		ID      uint   `gorm:"primaryKey" json:"-"`
		UserID  uint   `gorm:"not null;type:int8" json:"user_id"`
		Name    string `gorm:"not null;type:varchar(150)" json:"name"`
		Photo   string `gorm:"not null; type:varchar(100)" json:"photo"`
		Address string `gorm:"not null;type:varchar(250)" json:"address"`

		CreatedAt time.Time `gorm:"not null;type:timestamptz" json:"-"`
		UpdatedAt time.Time `gorm:"not null;type:timestamptz" json:"-"`
		DeletedAt gorm.DeletedAt

		User User `gorm:"foreignKey:user_id" json:"-"`
	}

	Agency struct {
		ID          uint   `gorm:"primaryKey" json:"-"`
		Name        string `gorm:"not null;type:varchar(150)" json:"name"`
		Email       string `gorm:"not null;type:varchar(100)" json:"email"`
		Description string `gorm:"not null;type:varchar(300)" json:"description"`
		Address     string `gorm:"not null;type:varchar(250)" json:"address"`
		Logo        string `gorm:"type:varchar(100)" json:"logo"`

		CreatedAt time.Time `gorm:"not null;type:timestamptz" json:"-"`
		UpdatedAt time.Time `gorm:"not null;type:timestamptz" json:"-"`
		DeletedAt gorm.DeletedAt

		Staff []AgencyStaff `gorm:"foreignKey:agency_id;constraint:OnDelete:CASCADE" json:"staff,omitempty"`
	}

	AgencyStaff struct {
		ID       uint            `gorm:"primaryKey" json:"-"`
		UserID   uint            `gorm:"not null;type:int8" json:"user_id"`
		AgencyID uint            `gorm:"not null;type:int8" json:"agency_id"`
		Name     string          `gorm:"not null;type:varchar(150)" json:"name"`
		Photo    string          `gorm:"not null; type:varchar(100)" json:"photo"`
		NID      string          `gorm:"not null;type:varchar(100)" json:"nid"`
		Address  string          `gorm:"not null;type:varchar(250)" json:"address"`
		Role     cdt.AgencyRole  `gorm:"not null;type:varchar(30)" json:"role"`

		CreatedAt time.Time `gorm:"not null;type:timestamptz" json:"-"`
		UpdatedAt time.Time `gorm:"not null;type:timestamptz" json:"-"`
		DeletedAt gorm.DeletedAt

		User User `gorm:"foreignKey:user_id" json:"-"`
	}

	TemporarySignupInfo struct {
		ID     uint `gorm:"primaryKey" json:"-"`
		UserID uint `gorm:"not null" json:"link_id"`

		// Agency Staff Info
		StaffPhoto   string `gorm:"not null; type:varchar(100)" json:"staff_photo"`
		StaffNID     string `gorm:"not null;type:varchar(100)" json:"staff_nid"`
		StaffAddress string `gorm:"not null;type:varchar(250)" json:"staff_address"`

		// Agency Info
		AgencyName        string `gorm:"not null;type:varchar(150)" json:"agency_name"`
		AgencyEmail       string `gorm:"not null;type:varchar(100)" json:"agency_email"`
		AgencyDescription string `gorm:"not null;type:varchar(300)" json:"agency_description"`
		AgencyAddress     string `gorm:"not null;type:varchar(250)" json:"agency_address"`
		AgencyLogo        string `gorm:"type:varchar(100)" json:"agency_logo"`

		CreatedAt time.Time `gorm:"not null;type:timestamptz" json:"-"`
		UpdatedAt time.Time `gorm:"not null;type:timestamptz" json:"-"`

		User User `gorm:"foreignKey:user_id" json:"-"`
	}
)

func (u *User) TableName() string {
	return "users"
}

func (c *Customer) TableName() string {
	return "customers"
}

func (a *Agency) TableName() string {
	return "agencies"
}

func (as *AgencyStaff) TableName() string {
	return "agency_staff"
}

func (tsi *TemporarySignupInfo) TableName() string {
	return "temporary_signup_info"
}
