package convert

import (
	"encoding/json"
)

func BuildDashbard(bts []byte) (*Dashboard, error) {
	dash := &Dashboard{}

	err := json.Unmarshal(bts, dash)
	if err != nil {
		return nil, err
	} else {
		return dash, nil
	}
}
