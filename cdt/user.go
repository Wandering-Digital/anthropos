package cdt

type AccountType string

const (
	UserTypeClient AccountType = "client"
	UserTypeAgent  AccountType = "agent"
)

func (ut AccountType) IsValid() bool {
	for _, userType := range []AccountType{UserTypeClient, UserTypeAgent} {
		if userType == ut {
			return true
		}
	}

	return false
}
