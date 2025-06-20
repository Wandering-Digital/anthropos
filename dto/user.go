package dto

import "github.com/Wandering-Digital/anthropos/cdt"

type GetUserResp struct {
	Email    string          `json:"email"`
	UserType cdt.AccountType `json:"user_type"`
}
