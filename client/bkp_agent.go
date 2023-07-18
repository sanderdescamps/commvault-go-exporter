package client

func (c *BkpService) GetBkpAgents(clientId uint64) ([]AgentProperty, error) {

	params := map[string]any{"clientId": clientId}
	req, err := c.client.NewRequest("GET", "/webconsole/api/Agent", params, nil)
	if err != nil {
		return nil, err
	}

	resp := &GetBkpAgentResp{}
	_, err = c.client.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp.AgentProperties, nil
}
