package request

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Wandering-Digital/anthropos/cdt"
	"github.com/Wandering-Digital/anthropos/internal/customerror"
)

const emailMaxLength = 254

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
		ve.Add("email", fmt.Sprint("must be within %d characters", emailMaxLength))
	} else if !isValidEmailFormat(cu.Email) {
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

// TO-DO: Expand email validation
func isValidEmailFormat(email string) bool {
	re := `^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}$`
	return regexp.MustCompile(re).MatchString(email)
}
