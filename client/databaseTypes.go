package client

type GetDatabasesInstansesResp struct {
	DbInstance []DbInstance `json:"dbInstance"`
	Error      struct {
		ErrorMessage string `json:"errorMessage"`
		ErrorCode    int    `json:"errorCode"`
	} `json:"error"`
}

type DbInstance struct {
	CloudVendorType        int    `json:"cloudVendorType"`
	NotReadyReason         string `json:"notReadyReason"`
	SLACategoryDescription string `json:"slaCategoryDescription"`
	Flags                  int    `json:"flags"`
	SLA                    string `json:"sla"`
	Version                string `json:"version"`
	BackupSize             int    `json:"backupSize"`
	SLAStatus              int    `json:"slaStatus"`
	IsCloudDB              bool   `json:"isCloudDB"`
	ClientType             int    `json:"clientType"`
	IsSnapEnabled          int    `json:"isSnapEnabled"`
	NoDBs                  int    `json:"noDBs"`
	ApplicationSize        int    `json:"applicationSize"`
	BackupTime             int    `json:"backupTime"`
	Status                 string `json:"status"`
	Instance               struct {
		ClientID      int    `json:"clientId"`
		ClientName    string `json:"clientName"`
		InstanceName  string `json:"instanceName"`
		DisplayName   string `json:"displayName"`
		AppName       string `json:"appName"`
		InstanceID    int    `json:"instanceId"`
		ApplicationID int    `json:"applicationId"`
		EntityInfo    struct {
			CompanyID       int    `json:"companyId"`
			CompanyName     string `json:"companyName"`
			MultiCommcellID int    `json:"multiCommcellId"`
		} `json:"entityInfo"`
	} `json:"instance"`
	LastBackupJobInfo struct {
		JobID                       int    `json:"jobID"`
		FailureReasonIds            string `json:"failureReasonIds"`
		CommCellID                  int    `json:"commCellID"`
		FailureReasonMessageEnglish string `json:"failureReasonMessageEnglish"`
		Status                      int    `json:"status"`
		StartTime                   struct {
			Time       int `json:"time"`
			EntityInfo struct {
				CompanyID       int    `json:"companyId"`
				CompanyName     string `json:"companyName"`
				MultiCommcellID int    `json:"multiCommcellId"`
			} `json:"entityInfo"`
		} `json:"startTime"`
		EndTime struct {
			Time       int `json:"time"`
			EntityInfo struct {
				CompanyID       int    `json:"companyId"`
				CompanyName     string `json:"companyName"`
				MultiCommcellID int    `json:"multiCommcellId"`
			} `json:"entityInfo"`
		} `json:"endTime"`
	} `json:"lastBackupJobInfo,omitempty"`
	Plan struct {
		PlanName   string `json:"planName"`
		PlanID     int    `json:"planId"`
		EntityInfo struct {
			CompanyID       int    `json:"companyId"`
			CompanyName     string `json:"companyName"`
			MultiCommcellID int    `json:"multiCommcellId"`
		} `json:"entityInfo"`
	} `json:"plan"`
}
