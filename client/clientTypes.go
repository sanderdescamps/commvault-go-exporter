package client

type LoginResp struct {
	AliasName            string             `json:"AliasName"`
	UserGUID             string             `json:"userGUID"`
	LoginAttempts        uint64             `json:"loginAttempts"`
	RemainingLockTime    uint64             `json:"remainingLockTime"`
	SmtpAddress          string             `json:"smtpAddress"`
	UserName             string             `json:"userName"`
	ProviderType         uint64             `json:"providerType"`
	Ccn                  uint64             `json:"ccn"`
	Token                string             `json:"token"`
	Capability           uint64             `json:"capability"`
	ForcePasswordChange  bool               `json:"forcePasswordChange"`
	IsAccountLocked      bool               `json:"isAccountLocked"`
	OwnerOrganization    *LoginOrganization `json:"ownerOrganization"`
	ProviderOrganization *LoginOrganization `json:"providerOrganization"`
	ErrList              []ErrorItem        `json:"errList"`
}

type LoginOrganization struct {
	ProviderId         uint64 `json:"providerId"`
	ProviderDomainName string `json:"providerDomainName"`
}

type LoginError struct {
	ErrLogMessage string `json:"errLogMessage"`
	ErrorCode     uint64 `json:"errorCode"`
}

type ErrorItem struct {
	ErrLogMessage string `json:"errLogMessage"`
	ErrorCode     uint64 `json:"errorCode"`
}
