package convert

import (
	"io/ioutil"
	"testing"

	"encoding/json"

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

func TestGraphDefinitionDecoder(t *testing.T) {
	tests := []struct {
		input    string
		expected *GraphDefinition
	}{
		{
			input:    `{"precision":"1"}`,
			expected: &GraphDefinition{Precision: "1"},
		},
		{
			input:    `{"precision":1}`,
			expected: &GraphDefinition{Precision: "1"},
		},
		{
			input:    `{"precision":"100%"}`,
			expected: &GraphDefinition{Precision: "100%"},
		},
		{
			input:    `{"precision":"*"}`,
			expected: &GraphDefinition{Precision: "*"},
		},
	}

	for _, test := range tests {
		gd := &GraphDefinition{}
		err := json.Unmarshal([]byte(test.input), gd)
		assert.NoError(t, err)
		assert.EqualValues(t, test.expected, gd)
	}
}
