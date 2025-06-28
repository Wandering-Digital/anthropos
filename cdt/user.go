package cdt

// AccountType represents a user's account type
type AccountType string

const (
	AccountTypeClient AccountType = "client"
	AccountTypeAgent  AccountType = "agent"
)

// AccountStatus represents the status of the user's account
type AccountStatus string

const (
	AccountStatusActive    AccountStatus = "active"
	AccountStatusInactive  AccountStatus = "inactive"
	AccountStatusSuspended AccountStatus = "suspended"
	AccountStatusBlocked   AccountStatus = "blocked"
)

// AgentOrganizationRole represents the role of an agent within their organization
type AgentOrganizationRole string

const (
	AgentOrganizationRoleOwner          AgentOrganizationRole = "owner"
	AgentOrganizationRoleAdmin          AgentOrganizationRole = "admin"
	AgentOrganizationRoleRepresentative AgentOrganizationRole = "representative"
)

func (at AccountType) IsValid() bool {
	for _, userType := range []AccountType{AccountTypeClient, AccountTypeAgent} {
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

func (aor AgentOrganizationRole) IsValid() bool {
	for _, agentOrganizationRole := range []AgentOrganizationRole{AgentOrganizationRoleOwner, AgentOrganizationRoleAdmin, AgentOrganizationRoleRepresentative} {
		if agentOrganizationRole == aor {
			return true
		}
	}

	return false
}
