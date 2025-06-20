package cdt

type AccountType string

const (
	UserTypeClient AccountType = "client"
	UserTypeAgent  AccountType = "agent"
)

func (at AccountType) IsValid() bool {
	for _, userType := range []AccountType{UserTypeClient, UserTypeAgent} {
		if userType == at {
			return true
		}
	}

	return false
}
