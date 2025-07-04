package migration

import (
	"github.com/Wandering-Digital/anthropos/model"
)

var (
	// Models describe models list for migration
	Models []interface{}
)

func init() {
	// Register models for migration

	Models = []interface{}{&model.User{}, &model.Customer{}, &model.Agency{}, &model.Agent{}, &model.TemporarySignupInfo{}}
}
