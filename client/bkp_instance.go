package client

import "fmt"

func (c *BkpService) GetBkpInstancesByClientId(clientId uint64) ([]InstanceProperty, error) {

	params := map[string]any{"clientId": clientId}
	req, err := c.client.NewRequest("GET", "/webconsole/api/instance", params, nil)
	if err != nil {
		return nil, err
	}

	resp := &GetBkpInstanceResp{}
	_, err = c.client.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp.InstanceProperties, nil
}

// func (c *StorageService) GetBkpInstancesByAgentId() (*GetBkpClientResp, error) {

// 	req, err := c.client.NewRequest("GET", "/webconsole/api/Client", nil, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp := &GetBkpClientResp{}
// 	_, err = c.client.Do(req, resp)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return resp, nil
// }

func (c *StorageService) GetBkpInstance(instanceId uint64) ([]InstanceProperty, error) {

	req, err := c.client.NewRequest("GET", fmt.Sprintf("/webconsole/api/instance/%d", instanceId), nil, nil)
	if err != nil {
		return nil, err
	}

	resp := &GetBkpInstanceResp{}
	_, err = c.client.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp.InstanceProperties, nil
}
