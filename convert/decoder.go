package convert

import "encoding/json"

func DecodeFromAPI(b []byte) (*Dashboard, error) {
	raw := &Dashboard{}
	err := json.Unmarshal(b, raw)
	if err != nil {
		return nil, err
	}

	return raw, err
}
