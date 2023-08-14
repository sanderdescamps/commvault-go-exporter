package client

import (
	"fmt"
)

type StorageService struct {
	client *CommvaultClient
}

type LibraryType int64

const (
	LibraryType_Tape              VmStatus = 0
	LibraryType_Disk_Cloud        VmStatus = 3
	LibraryType_StandAloneLibrary VmStatus = 5
)

func (c *StorageService) GetLibraries() ([]EntityInfo, error) {

	req, err := c.client.NewRequest("GET", "/webconsole/api/Library", nil, nil)
	if err != nil {
		return nil, err
	}

	resp := &GetLibraryResp{}
	_, err = c.client.Do(req, resp)
	if err != nil {
		return nil, err
	}

	result := []EntityInfo{}
	for _, i := range *resp.Response {
		result = append(result, i.EntityInfo)
	}
	return result, nil
}

func (c *StorageService) GetLibrariesDetails() ([]LibraryDetail, error) {

	result := []LibraryDetail{}
	libs, err := c.GetLibraries()
	if err != nil {
		return nil, err
	}

	for _, lib := range libs {
		req, err := c.client.NewRequest("GET", fmt.Sprintf("/webconsole/api/Library/%d", lib.Id), nil, nil)
		if err != nil {
			return nil, err
		}
		resp := &GetLibraryDetailsResp{}
		_, err = c.client.Do(req, resp)
		if err != nil {
			return nil, err
		}
		if resp.LibraryInfo != nil {
			result = append(result, *resp.LibraryInfo)
		}
	}

	return result, nil
}
