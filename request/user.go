package request

import (
	"fmt"
	"strings"

	"github.com/Wandering-Digital/anthropos/cdt"
	"github.com/Wandering-Digital/anthropos/internal/customerror"
	"github.com/Wandering-Digital/anthropos/internal/utils/email"
)

const emailMaxLength = 100

type CreateUser struct {
	Email       string          `json:"email"`
	Password    string          `json:"password"`
	AccountType cdt.AccountType `json:"account_type"`
}

func (cu *CreateUser) Validate() *customerror.ValidationError {
	ve := customerror.NewValidationError()

	cu.Email = strings.TrimSpace(cu.Email)

	if cu.Email == "" {
		ve.Add("email", "is required")
	} else if len(cu.Email) > emailMaxLength {
		ve.Add("email", fmt.Sprintf("must be within %d characters", emailMaxLength))
	} else if !email.IsValidFormat(cu.Email) {
		ve.Add("email", "has an invalid format")
	}

	if cu.Password == "" {
		ve.Add("password", "is required")
	}

	if cu.AccountType == "" {
		ve.Add("account_type", "is required")
	} else if !cu.AccountType.IsValid() {
		ve.Add("account_type", "is invalid")
	}

	if !ve.IsNil() {
		return ve
	}

	return nil
}
