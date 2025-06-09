package request

import (
	"fmt"
	"strings"

	"github.com/Wandering-Digital/anthropos/cdt"
	"github.com/Wandering-Digital/anthropos/internal/customerror"
)

const userNameMaxLength = 50

type CreateUser struct {
	UserName string       `json:"user_name"`
	Password string       `json:"password"`
	UserType cdt.UserType `json:"user_type"`
}

func (cu *CreateUser) Validate() *customerror.ValidationError {
	ve := customerror.NewValidationError()

	cu.UserName = strings.TrimSpace(cu.UserName)

	if cu.UserName == "" {
		ve.Add("user_name", "is required")
	} else if len(cu.UserName) > userNameMaxLength {
		ve.Add("user_name", fmt.Sprint("must be withing %d characters", userNameMaxLength))
	}

	if cu.Password == "" {
		ve.Add("password", "is required")
	}

	if cu.UserType == "" {
		ve.Add("user_type", "is required")
	} else if !cu.UserType.IsValid() {
		ve.Add("user_type", "is invalid")
	}

	if !ve.IsNil() {
		return ve
	}

	return nil
}
