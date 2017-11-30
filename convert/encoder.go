package convert

import "github.com/promiseofcake/hclencoder"

func EncodeToHCL(d *Dashboard) ([]byte, error) {
	return hclencoder.Encode(d)
}
