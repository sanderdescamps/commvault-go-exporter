package client

type GetBkpAgentResp struct {
	AgentProperties []AgentProperty `json:"agentProperties"`
}

type AgentProperty struct {
	AgentProperties struct {
		InstallDate     int    `json:"installDate"`
		AgentVersion    string `json:"agentVersion"`
		IsMarkedDeleted bool   `json:"isMarkedDeleted"`
	} `json:"AgentProperties"`
	IdaActivityControl struct {
		ActivityControlOptions []struct {
			ActivityType       int  `json:"activityType"`
			EnableAfterADelay  bool `json:"enableAfterADelay"`
			EnableActivityType bool `json:"enableActivityType"`
		} `json:"activityControlOptions"`
	} `json:"idaActivityControl"`
	IdaEntity struct {
		ClientID      int    `json:"clientId"`
		ClientName    string `json:"clientName"`
		AppName       string `json:"appName"`
		CommCellID    int    `json:"commCellId"`
		ApplicationID int    `json:"applicationId"`
	} `json:"idaEntity"`
}
