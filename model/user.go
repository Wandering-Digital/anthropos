package model

import (
	"time"

	"github.com/Wandering-Digital/anthropos/cdt"

	"gorm.io/gorm"
)

type (
	User struct {
		ID          uint              `gorm:"column:id;primaryKey" json:"id"`
		Email       string            `gorm:"column:email;not null;type:varchar(100)" json:"email"`
		Password    string            `gorm:"column:password;not null;type:varchar(100)" json:"-"`
		AccountType cdt.AccountType   `gorm:"column:account_type;not null;type:varchar(20)" json:"account_type"`
		Status      cdt.AccountStatus `gorm:"column:status;not null;type:varchar(30)" json:"status"`

		CreatedAt time.Time `gorm:"column:created_at;not null;type:timestamptz" json:"-"`
		UpdatedAt time.Time `gorm:"column:updated_at;not null;type:timestamptz" json:"-"`
		DeletedAt gorm.DeletedAt
	}

	Customer struct {
		ID      uint   `gorm:"column:id;primaryKey" json:"-"`
		UserID  uint   `gorm:"column:user_id;not null;type:int8" json:"user_id"`
		Name    string `gorm:"column:name;not null;type:varchar(150)" json:"name"`
		Photo   string `gorm:"column:photo;not null; type:varchar(100)" json:"photo"`
		Address string `gorm:"column:address;not null;type:varchar(250)" json:"address"`

		CreatedAt time.Time `gorm:"column:created_at;not null;type:timestamptz" json:"-"`
		UpdatedAt time.Time `gorm:"column:updated_at;not null;type:timestamptz" json:"-"`
		DeletedAt gorm.DeletedAt

		User User `gorm:"foreignKey:user_id" json:"-"`
	}

	Agency struct {
		ID          uint   `gorm:"column:id;primaryKey" json:"-"`
		Name        string `gorm:"column:name;not null;type:varchar(150)" json:"name"`
		Email       string `gorm:"column:email;not null;type:varchar(100)" json:"email"`
		Description string `gorm:"column:description;not null;type:varchar(300)" json:"description"`
		Address     string `gorm:"column:address;not null;type:varchar(250)" json:"address"`
		Logo        string `gorm:"column:logo;type:varchar(100)" json:"logo"`

		CreatedAt time.Time `gorm:"column:created_at;not null;type:timestamptz" json:"-"`
		UpdatedAt time.Time `gorm:"column:updated_at;not null;type:timestamptz" json:"-"`
		DeletedAt gorm.DeletedAt

		Staff []Agent `gorm:"foreignKey:agency_id;constraint:OnDelete:CASCADE" json:"staff,omitempty"`
	}

	Agent struct {
		ID       uint           `gorm:"column:id;primaryKey" json:"-"`
		UserID   uint           `gorm:"column:user_id;not null;type:int8" json:"user_id"`
		AgencyID uint           `gorm:"column:agency_id;not null;type:int8" json:"agency_id"`
		Name     string         `gorm:"column:name;not null;type:varchar(150)" json:"name"`
		Photo    string         `gorm:"column:photo;not null; type:varchar(100)" json:"photo"`
		NID      string         `gorm:"column:nid;not null;type:varchar(100)" json:"nid"`
		Address  string         `gorm:"column:address;not null;type:varchar(250)" json:"address"`
		Role     cdt.AgencyRole `gorm:"column:role;not null;type:varchar(30)" json:"role"`

		CreatedAt time.Time `gorm:"column:created_at;not null;type:timestamptz" json:"-"`
		UpdatedAt time.Time `gorm:"column:updated_at;not null;type:timestamptz" json:"-"`
		DeletedAt gorm.DeletedAt

		User User `gorm:"foreignKey:user_id" json:"-"`
	}

	TemporarySignupInfo struct {
		ID     uint `gorm:"column:id;primaryKey" json:"-"`
		UserID uint `gorm:"column:user_id;not null" json:"user_id"`

		// Agency Staff Info
		AgentPhoto   string `gorm:"column:agent_photo;not null; type:varchar(100)" json:"agent_photo"`
		AgentNID     string `gorm:"column:agent_nid;not null;type:varchar(100)" json:"agent_nid"`
		AgentAddress string `gorm:"column:agent_address;not null;type:varchar(250)" json:"agent_address"`

		// Agency Info
		AgencyName        string `gorm:"column:agency_name;not null;type:varchar(150)" json:"agency_name"`
		AgencyEmail       string `gorm:"column:agency_email;not null;type:varchar(100)" json:"agency_email"`
		AgencyDescription string `gorm:"column:agency_description;not null;type:varchar(300)" json:"agency_description"`
		AgencyAddress     string `gorm:"column:agency_address;not null;type:varchar(250)" json:"agency_address"`
		AgencyLogo        string `gorm:"column:agency_logo;type:varchar(100)" json:"agency_logo"`

		CreatedAt time.Time `gorm:"column:created_at;not null;type:timestamptz" json:"-"`
		UpdatedAt time.Time `gorm:"column:updated_at;not null;type:timestamptz" json:"-"`

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

func (as *Agent) TableName() string {
	return "agents"
}

func (tsi *TemporarySignupInfo) TableName() string {
	return "temporary_signup_info"
}
