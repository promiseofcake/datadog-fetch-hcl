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

func float64Pointer(v float64) *float64 {
	return &v
}

func stringPointer(v string) *string {
	return &v
}

func TestYaxisDecoder(t *testing.T) {
	tests := []struct {
		input    string
		expected *Yaxis
	}{
		{
			input:    `{"min":"auto", "max":"auto"}`,
			expected: &Yaxis{},
		},
		{
			input:    `{"min":"0"}`,
			expected: &Yaxis{Min: float64Pointer(0)},
		},
		{
			input:    `{"min":"1"}`,
			expected: &Yaxis{Min: float64Pointer(1)},
		},
		{
			input:    `{"max":"100"}`,
			expected: &Yaxis{Max: float64Pointer(100)},
		},
		{
			input:    `{"max":"100","scale":"log"}`,
			expected: &Yaxis{Max: float64Pointer(100), Scale: stringPointer("log")},
		},
	}

	for _, test := range tests {
		y := &Yaxis{}
		err := json.Unmarshal([]byte(test.input), y)
		assert.NoError(t, err)
		assert.EqualValues(t, test.expected, y)
	}
}
