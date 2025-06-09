package cdt

type UserType string

const (
	UserTypeClient UserType = "client"
	UserTypeAgent  UserType = "agent"
)

func (ut UserType) IsValid() bool {
	for _, userType := range []UserType{UserTypeClient, UserTypeAgent} {
		if userType == ut {
			return true
		}
	}

	return false
}
