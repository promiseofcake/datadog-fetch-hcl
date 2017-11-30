package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/promiseofcake/datadog-fetch-hcl/convert"
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

func (d *DataDogClient) GetDashboard(id int) (*convert.Dashboard, error) {
	url := fmt.Sprintf(
		"%s/%d?api_key=%s&application_key=%s",
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

	if res.StatusCode != http.StatusOK {
		bts, _ := ioutil.ReadAll(res.Body)
		return nil, fmt.Errorf("API Call: %s", bts)
	}

	dash := &convert.Dashboard{}
	err = json.NewDecoder(res.Body).Decode(dash)
	if err != nil {
		return nil, err
	}

	return dash, err
}
