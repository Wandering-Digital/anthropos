package dto

import "github.com/Wandering-Digital/anthropos/cdt"

type GetUserResp struct {
	UserName string       `json:"user_name"`
	UserType cdt.UserType `json:"user_type"`
}
