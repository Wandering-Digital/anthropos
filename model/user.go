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

	Traveller struct {
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

	Organization struct {
		ID          uint   `gorm:"primaryKey" json:"-"`
		Name        string `gorm:"not null;type:varchar(150)" json:"name"`
		Email       string `gorm:"not null;type:varchar(100)" json:"email"`
		Description string `gorm:"not null;type:varchar(300)" json:"description"`
		Address     string `gorm:"not null;type:varchar(250)" json:"address"`
		Logo        string `gorm:"type:varchar(100)" json:"logo"`

		CreatedAt time.Time `gorm:"not null;type:timestamptz" json:"-"`
		UpdatedAt time.Time `gorm:"not null;type:timestamptz" json:"-"`
		DeletedAt gorm.DeletedAt

		Agents []Agent `gorm:"foreignKey:organization_id;constraint:OnDelete:CASCADE" json:"agents,omitempty"`
	}

	Agent struct {
		ID             uint                      `gorm:"primaryKey" json:"-"`
		UserID         uint                      `gorm:"not null;type:int8" json:"user_id"`
		OrganizationID uint                      `gorm:"not null;type:int8" json:"organization_id"`
		Name           string                    `gorm:"not null;type:varchar(150)" json:"name"`
		Photo          string                    `gorm:"not null; type:varchar(100)" json:"photo"`
		NID            string                    `gorm:"not null;type:varchar(100)" json:"nid"`
		Address        string                    `gorm:"not null;type:varchar(250)" json:"address"`
		Role           cdt.AgentOrganizationRole `gorm:"not null;type:varchar(30)" json:"role"`

		CreatedAt time.Time `gorm:"not null;type:timestamptz" json:"-"`
		UpdatedAt time.Time `gorm:"not null;type:timestamptz" json:"-"`
		DeletedAt gorm.DeletedAt

		User User `gorm:"foreignKey:user_id" json:"-"`
	}

	TemporarySignupInfo struct {
		ID     uint `gorm:"primaryKey" json:"-"`
		UserID uint `gorm:"not null" json:"link_id"`

		// Agent Info
		AgentPhoto   string `gorm:"not null; type:varchar(100)" json:"agent_photo"`
		AgentNID     string `gorm:"not null;type:varchar(100)" json:"agent_nid"`
		AgentAddress string `gorm:"not null;type:varchar(250)" json:"agent_address"`

		// Organization Info
		OrganizationName        string `gorm:"not null;type:varchar(150)" json:"organization_name"`
		OrganizationEmail       string `gorm:"not null;type:varchar(100)" json:"organization_email"`
		OrganizationDescription string `gorm:"not null;type:varchar(300)" json:"organization_description"`
		OrganizationAddress     string `gorm:"not null;type:varchar(250)" json:"organization_address"`
		OrganizationLogo        string `gorm:"type:varchar(100)" json:"organization_photo"`

		CreatedAt time.Time `gorm:"not null;type:timestamptz" json:"-"`
		UpdatedAt time.Time `gorm:"not null;type:timestamptz" json:"-"`

		User User `gorm:"foreignKey:user_id" json:"-"`
	}
)

func (u *User) TableName() string {
	return "users"
}

func (t *Traveller) TableName() string {
	return "travellers"
}

func (o *Organization) TableName() string {
	return "organizations"
}

func (a *Agent) TableName() string {
	return "agents"
}

func (tsi *TemporarySignupInfo) TableName() string {
	return "temporary_signup_info"
}
