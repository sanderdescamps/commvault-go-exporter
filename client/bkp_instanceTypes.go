package client

type GetBkpInstanceResp struct {
	InstanceProperties []InstanceProperty `json:"instanceProperties"`
}

type InstanceProperty struct {
	IsDeleted       bool  `json:"isDeleted"`
	ApplicationSize int64 `json:"applicationSize"`
	ScIdxEnabled    int   `json:"scIdxEnabled"`
	Instance        struct {
		ClientID      int    `json:"clientId"`
		ClientName    string `json:"clientName"`
		InstanceName  string `json:"instanceName"`
		DisplayName   string `json:"displayName"`
		AppName       string `json:"appName"`
		InstanceID    int    `json:"instanceId"`
		ApplicationID int    `json:"applicationId"`
		InstanceGUID  string `json:"instanceGUID"`
		EntityInfo    struct {
			CompanyID       int    `json:"companyId"`
			CompanyName     string `json:"companyName"`
			MultiCommcellID int    `json:"multiCommcellId"`
		} `json:"entityInfo"`
	} `json:"instance"`
	InstanceActivityControl struct {
		ActivityControlOptions []struct {
			ActivityType       int  `json:"activityType"`
			EnableAfterADelay  bool `json:"enableAfterADelay"`
			EnableActivityType bool `json:"enableActivityType"`
		} `json:"activityControlOptions"`
	} `json:"instanceActivityControl"`
	CloudDBInstance struct {
		OverRideAccessNode bool `json:"overRideAccessNode"`
		AccessNodes        struct {
		} `json:"accessNodes"`
	} `json:"cloudDBInstance"`
	MssqlInstance struct {
		AwsSqlInfo struct {
			SqlInstanceCredentials struct {
			} `json:"sqlInstanceCredentials"`
		} `json:"awsSqlInfo"`
		AgProperties struct {
			SQLAvailabilityReplicasList struct {
			} `json:"SQLAvailabilityReplicasList"`
			AvailabilityGroup struct {
			} `json:"availabilityGroup"`
			ProxyClient struct {
				EntityInfo struct {
					CompanyID       int    `json:"companyId"`
					CompanyName     string `json:"companyName"`
					MultiCommcellID int    `json:"multiCommcellId"`
				} `json:"entityInfo"`
			} `json:"proxyClient"`
		} `json:"agProperties"`
		AzureInfo struct {
			AzureCredentials struct {
			} `json:"azureCredentials"`
			AzureSqlInstanceCredentials struct {
			} `json:"azureSqlInstanceCredentials"`
		} `json:"azureInfo"`
		Proxies struct {
		} `json:"proxies"`
	} `json:"mssqlInstance,omitempty"`
}
