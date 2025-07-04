package cdt

// AccountType represents a user's account type
type AccountType string

const (
	AccountTypePlatformAdmin AccountType = "platform_admin"
	AccountTypeAgent         AccountType = "agent"
	AccountTypeCustomer      AccountType = "customer"
)

// AccountStatus represents the status of the user's account
type AccountStatus string

const (
	AccountStatusActive    AccountStatus = "active"
	AccountStatusInactive  AccountStatus = "inactive"
	AccountStatusSuspended AccountStatus = "suspended"
	AccountStatusBlocked   AccountStatus = "blocked"
)

type AgencyRole string

const (
	AgencyRoleOwner AgencyRole = "owner"
	AgencyRoleStaff AgencyRole = "staff"
)

func (at AccountType) IsValid() bool {
	for _, userType := range []AccountType{
		AccountTypePlatformAdmin,
		AccountTypeAgent,
		AccountTypeCustomer,
	} {
		if userType == at {
			return true
		}
	}

	return false
}

func (as AccountStatus) IsValid() bool {
	for _, accountStatus := range []AccountStatus{AccountStatusActive, AccountStatusInactive, AccountStatusSuspended, AccountStatusBlocked} {
		if accountStatus == as {
			return true
		}
	}

	return false
}

func (ar AgencyRole) IsValid() bool {
	for _, agencyRole := range []AgencyRole{AgencyRoleOwner, AgencyRoleStaff} {
		if agencyRole == ar {
			return true
		}
	}

	return false
}
