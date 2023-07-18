package client

type DatabaseService struct {
	client *CommvaultClient
}

func (c *DatabaseService) GetDatabaseInstances() ([]DbInstance, error) {

	req, err := c.client.NewRequest("GET", "/webconsole/api/databases/instances", nil, nil)
	if err != nil {
		return nil, err
	}

	resp := &GetDatabasesInstansesResp{}
	_, err = c.client.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp.DbInstance, nil
}
