package client

type GetBkpClientResp struct {
	ClientProperty []ClientProperty `json:"clientProperties"`
}

type ClientProperty struct {
	ClientProps struct {
		ISCSIPort   int `json:"iSCSIPort"`
		ClusterType int `json:"clusterType"`
	} `json:"clientProps"`
	Client struct {
		EvmgrcPort   int `json:"evmgrcPort"`
		CvdPort      int `json:"cvdPort"`
		ClientEntity struct {
			HostName    string `json:"hostName"`
			ClientID    int    `json:"clientId"`
			ClientName  string `json:"clientName"`
			DisplayName string `json:"displayName"`
			Type        int    `json:"_type_"`
			ClientGUID  string `json:"clientGUID"`
			EntityInfo  struct {
				CompanyID       int    `json:"companyId"`
				CompanyName     string `json:"companyName"`
				MultiCommcellID int    `json:"multiCommcellId"`
			} `json:"entityInfo"`
		} `json:"clientEntity"`
	} `json:"client"`
}
