package client

type GetVmResp struct {
	TotalRecords     uint64          `json:"totalRecords"`
	PageNo           uint64          `json:"pageNo"`
	ErrorMessage     string          `json:"errorMessage"`
	ErrorCode        uint64          `json:"errorCode"`
	PageSize         uint64          `json:"pageSize"`
	VmStatusInfoList *[]VmStatusInfo `json:"vmStatusInfoList"`
}

type VmStatusInfo struct {
	VmHost              string                   `json:"vmHost"`
	VmGuestSpace        uint64                   `json:"vmGuestSpace"`
	BkpStartTime        uint64                   `json:"bkpStartTime"`
	IsIndexingV2        bool                     `json:"isIndexingV2"`
	Type                uint64                   `json:"type"`
	VmStatus            uint64                   `json:"vmStatus"`
	LatestRecoveryTime  uint64                   `json:"LatestRecoveryTime"`
	VmBackupJob         uint64                   `json:"vmBackupJob"`
	StrOSName           string                   `json:"strOSName"`
	IsDeleted           bool                     `json:"isDeleted"`
	Vendor              uint64                   `json:"vendor"`
	VmSize              uint64                   `json:"vmSize"`
	VmUsedSpace         uint64                   `json:"vmUsedSpace"`
	SubclientId         uint64                   `json:"subclientId"`
	BkpEndTime          uint64                   `json:"bkpEndTime"`
	VmAgent             string                   `json:"vmAgent"`
	OldestRecoveryPoint uint64                   `json:"OldestRecoveryPoint"`
	IsContentIndexded   bool                     `json:"isContentIndexded"`
	Name                string                   `json:"name"`
	VmHardwareVer       string                   `json:"vmHardwareVer"`
	StrGUID             string                   `json:"strGUID"`
	SubclientName       string                   `json:"subclientName"`
	Client              VmStatusInfoClient       `json:"client"`
	ProxyClient         VmStatusInfoProxyClient  `json:"proxyClient"`
	Plan                VmStatusInfoPlan         `json:"plan"`
	PseudoClient        VmStatusInfoPseudoClient `json:"pseudoClient"`
}

type VmStatusInfoClient struct {
	ClientId   uint64 `json:"clientId"`
	ClientName string `json:"clientName"`
}

type VmStatusInfoProxyClient struct {
	ClientId   uint64 `json:"clientId"`
	ClientName string `json:"clientName"`
}

type VmStatusInfoPlan struct {
	PlanId   uint64 `json:"planId"`
	PlanName string `json:"planName"`
}

type VmStatusInfoPseudoClient struct {
	ClientId   uint64                        `json:"clientId"`
	ClientName string                        `json:"clientName"`
	Flags      VmStatusInfoPseudoClientFlags `json:"flags"`
}

type VmStatusInfoPseudoClientFlags struct {
	Disabled bool `json:"disabled"`
}

// type SiteInfo struct {
// 	BkpJobsToSync            string                           `json:"BkpJobsToSync"`
// 	Flags                    uint64                           `json:"flags"`
// 	LastRestoreTime          uint64                           `json:"lastRestoreTime"`
// 	ReplicationId            uint64                           `json:"replicationId"`
// 	VMSyncedTillTime         uint64                           `json:"VMSyncedTillTime"`
// 	DestinationName          string                           `json:"destinationName"`
// 	ReplicationGuid          string                           `json:"replicationGuid"`
// 	Options                  string                           `json:"options"`
// 	DestinationGuid          string                           `json:"destinationGuid"`
// 	ValidationFailCount      uint64                           `json:"ValidationFailCount"`
// 	CopyPrecedence           uint64                           `json:"copyPrecedence"`
// 	PowerOn                  uint64                           `json:"powerOn"`
// 	LastSyncedBkpJob         uint64                           `json:"lastSyncedBkpJob"`
// 	Spare3                   string                           `json:"spare3"`
// 	SLAStatus                uint64                           `json:"SLAStatus"`
// 	Spare2                   string                           `json:"spare2"`
// 	LastBackupTime           uint64                           `json:"lastBackupTime"`
// 	Spare1                   string                           `json:"spare1"`
// 	FailoverStatus           uint64                           `json:"FailoverStatus"`
// 	SourceName               string                           `json:"sourceName"`
// 	SourceVendorType         uint64                           `json:"SourceVendorType"`
// 	SourceGuid               string                           `json:"sourceGuid"`
// 	CopyPrecedenceApplicable bool                             `json:"copyPrecedenceApplicable"`
// 	Status                   uint64                           `json:"status"`
// 	SubTask                  SiteInfoSubTask              `json:"subTask"`
// 	ReplicationTarget        []SiteInfoVMReplInfoProperty `json:"replicationTarget"`
// 	DestProxy                SiteInfoDestProxy            `json:"destProxy"`
// 	DestinationInstance      SiteInfoDestinationInstance  `json:"destinationInstance"`
// 	VMReplInfoProperties     []SiteInfoVMReplInfoProperty `json:"VMReplInfoProperties"`
// 	Entity                   SiteInfoEntity               `json:"entity"`
// }

// type SiteInfoSubTask struct {
// 	SubtaskId   uint64 `json:"subtaskId"`
// 	SubtaskName string `json:"subtaskName"`
// 	TaskId      uint64 `json:"taskId"`
// }

// type SiteInfoReplicationTarget struct {
// 	VmAllocPolicyName string `json:"vmAllocPolicyName"`
// 	VmAllocPolicyId   uint64 `json:"vmAllocPolicyId"`
// }

// type SiteInfoDestProxy struct {
// 	ClientId   uint64 `json:"clientId"`
// 	ClientName string `json:"clientName"`
// }

// type SiteInfoDestinationInstance struct {
// 	ClientId     uint64 `json:"clientId"`
// 	ClientName   uint64 `json:"clientName"`
// 	InstanceName string `json:"instanceName"`
// 	InstanceId   uint64 `json:"instanceId"`
// }
// type SiteInfoVMReplInfoProperty struct {
// 	PropertyValue string `json:"propertyValue"`
// 	PropertyId    uint64 `json:"propertyId"`
// }
// type SiteInfoEntity struct {
// 	ClientName    uint64 `json:"clientName"`
// 	InstanceName  string `json:"instanceName"`
// 	BackupsetId   uint64 `json:"backupsetId"`
// 	InstanceId    uint64 `json:"instanceId"`
// 	CommCellId    uint64 `json:"commCellId"`
// 	SubclientId   uint64 `json:"subclientId"`
// 	ClientId      uint64 `json:"clientId"`
// 	AppName       uint64 `json:"appName"`
// 	BackupsetName string `json:"backupsetName"`
// 	ApplicationId uint64 `json:"applicationId"`
// 	SubclientName string `json:"subclientName"`
// }
