package client

type BkpService struct {
	client *CommvaultClient
}

func (c *BkpService) GetBkpClients() ([]ClientProperty, error) {

	req, err := c.client.NewRequest("GET", "/webconsole/api/Client", nil, nil)
	if err != nil {
		return nil, err
	}

	resp := &GetBkpClientResp{}
	_, err = c.client.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp.ClientProperty, nil
}
