package convert

import "github.com/rodaine/hclencoder"

func EncodeToHCL(d *Dashboard) ([]byte, error) {
	return hclencoder.Encode(d)
}
