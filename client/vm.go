package client

import (
	"encoding/json"
	"fmt"
)

type VmService struct {
	client *CommvaultClient
}

type VmStatus int64

const (
	VmStatus_All               VmStatus = 0
	VmStatus_Protected         VmStatus = 1
	VmStatus_NotProtected      VmStatus = 2
	VmStatus_Pending           VmStatus = 3
	VmStatus_BackedUpWithError VmStatus = 4
	VmStatus_Discovered        VmStatus = 5
)

func (c *VmService) GetVmStatus(status VmStatus) (*[]VmStatusInfo, error) {

	result := []VmStatusInfo{}
	start := 0
	step := 1000
	end := step
	for {
		req, err := c.client.NewRequest("GET", "webconsole/api/VM", map[string]any{"status": status}, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Add("pagingInfo", fmt.Sprintf("%d,%d", start, end))
		req.Header.Add("sortingInfo", "asc:2")

		resp := &GetVmResp{}
		_, err = c.client.Do(req, resp)
		if err != nil {
			return nil, err
		}

		if resp.TotalRecords > 0 {
			result = append(result, *(resp.VmStatusInfoList)...)
		}

		if int(resp.TotalRecords) > end {
			end = end + step
			start = start + step
		} else {
			break
		}
	}
	return &result, nil
}

func (c *VmService) GetVmStatusDebug(status VmStatus) (*[]VmStatusInfo, error) {

	_, err := c.GetVmStatus(status)
	if err != nil {
		return nil, err
	}
	resp := GetVmResp{}
	err = json.Unmarshal([]byte(fake_vmstatus_response()), &resp)
	if err != nil {
		return nil, err
	}
	return resp.VmStatusInfoList, nil

}
