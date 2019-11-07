package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client interface {
	GetDashboard(int)
}

const ddDashboardAPIUrl = "https://app.datadoghq.com/api/v1/dash"

type DataDogClient struct {
	apiKey string
	appKey string
}

func NewDataDog(apiKey, appKey string) *DataDogClient {
	return &DataDogClient{
		apiKey: apiKey,
		appKey: appKey,
	}
}

func (d *DataDogClient) FetchJSON(id string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/%s?api_key=%s&application_key=%s",
		ddDashboardAPIUrl,
		id,
		d.apiKey,
		d.appKey,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bts, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API Call: %s", bts)
	}

	return bts, nil
}
