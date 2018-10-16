package convert

import (
	"io/ioutil"
	"testing"

	"github.com/rodaine/hclencoder"
	"github.com/stretchr/testify/assert"
)

func TestEncodeToHCL(t *testing.T) {
	bts, err := ioutil.ReadFile("../fixtures/input.json")
	assert.NoError(t, err)

	dash, err := DecodeFromAPI(bts)
	assert.NoError(t, err)

	_, err = hclencoder.Encode(dash)
	assert.NoError(t, err)

	// TODO ensure that we can re-code from the HCL into the appropriate structs
	//d := Dashboard{}
	//fmt.Println(string(out))
	//err = hcl.Decode(&d, string(out))
	//assert.NoError(t, err)
}
