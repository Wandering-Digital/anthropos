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

	Models = append(Models, &model.User{}, &model.Traveller{}, &model.Organization{}, &model.Agent{}, &model.TemporarySignupInfo{})
}
