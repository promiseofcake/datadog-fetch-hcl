package convert

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeFromAPI(t *testing.T) {
	bts, err := ioutil.ReadFile("../fixtures/input.json")
	assert.NoError(t, err)

	dash, err := DecodeFromAPI(bts)

	assert.NoError(t, err)
	assert.True(t, len(dash.Graphs) > 0)
	assert.True(t, len(dash.Templates) > 0)
}
