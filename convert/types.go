package convert

import (
	"encoding/json"

	"github.com/zorkian/go-datadog-api"
)

type Dashboard struct {
	Dash `json:"dash" hcl:"resource"`
}

type Dash struct {
	Title       string     `json:"title" hcl:"title"`
	Description string     `json:"description" hcl:"description"`
	ReadOnly    bool       `json:"read_only,omitempty" hcl:"read_only"`
	Templates   []Template `json:"template_variables,omitempty" hcl:"template_variable" hcle:"omitempty"`
	Graphs      []Graph    `json:"graphs" hcl:"graph"`
}

type Template struct {
	Name    string `json:"name" hcl:"name"`
	Default string `json:"default" hcl:"default"`
	Prefix  string `json:"prefix,omitempty" hcl:"prefix" hcle:"omitempty"`
}

type Graph struct {
	Title           string `json:"title" hcl:"title"`
	GraphDefinition `json:"definition" hcl:",squash"`
}

type GraphDefinition struct {
	Viz       string             `json:"viz" hcl:"viz"`
	AutoScale bool               `json:"autoscale,omitempty" hcl:"autoscale" hcle:"omitempty"`
	Precision datadog.PrecisionT `json:"precision,omitempty" hcl:"precision" hcle:"omitempty"`
	Requests  []Request          `json:"requests" hcl:"request"`
	Events    []Event            `json:"events" hcl:"events" hcle:"omitempty"`
	Yaxis     Yaxis              `json:"yaxis,omitempty" hcl:"yaxis" hcle:"omitempty"`
}

// copied from https://github.com/zorkian/go-datadog-api due to custom unmarshalling
// we need to duplicate due to custom HCL marshalling
// TODO: use zorkian/go-datadog-api entirely and build a "to our struct" function
type Yaxis struct {
	Min   *float64 `json:"min,omitempty" hcl:"min" hcle:"omitempty"`
	Max   *float64 `json:"max,omitempty" hcl:"max" hcle:"omitempty"`
	Scale *string  `json:"scale,omitempty" hcl:"scale" hcle:"omitempty"`
}

func (y *Yaxis) UnmarshalJSON(data []byte) error {
	type Alias Yaxis
	wrapper := &struct {
		Min *json.Number `json:"min,omitempty"`
		Max *json.Number `json:"max,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(y),
	}

	if err := json.Unmarshal(data, &wrapper); err != nil {
		return err
	}

	if wrapper.Min != nil {
		if *wrapper.Min == "auto" {
			y.Min = nil
		} else {
			f, err := wrapper.Min.Float64()
			if err != nil {
				return err
			}
			y.Min = &f
		}
	}

	if wrapper.Max != nil {
		if *wrapper.Max == "auto" {
			y.Max = nil
		} else {
			f, err := wrapper.Max.Float64()
			if err != nil {
				return err
			}
			y.Max = &f
		}
	}
	return nil
}

type GraphStyle struct {
	Palette     string `json:"palette,omitempty" hcl:"palette" hcle:"omitempty"`
	PaletteFlip bool   `json:"palette_flip,omitempty" hcl:"palette_flip" hcle:"omitempty"`
}

type Request struct {
	Q                  string              `json:"q" hcl:"q"`
	Aggregator         string              `json:"aggregator,omitempty" hcl:"aggregator" hcle:"omitempty"`
	Stacked            bool                `json:"stacked,omitempty" hcl:"stacked" hcle:"omitempty"`
	Type               string              `json:"type,omitempty" hcl:"type" hcle:"omitempty"`
	Style              RequestStyle        `json:"style,omitempty" hcl:"style" hcle:"omitempty"`
	ConditionalFormats []ConditionalFormat `json:"conditional_formats,omitempty" hcl:"conditional_format" hcle:"omitempty"`
}

type RequestStyle struct {
	Palette string `json:"palette,omitempty" hcl:"palette" hcle:"omitempty"`
	Width   string `json:"width,omitempty" hcl:"width" hcle:"omitempty"`
	Type    string `json:"type,omitempty" hcl:"type" hcle:"omitempty"`
}

type ConditionalFormat struct {
	Palette    string      `json:"palette,omitempty" hcl:"palette" hcle:"omitempty"`
	Comparator string      `json:"comparator,omitempty" hcl:"comparator" hcle:"omitempty"`
	Value      json.Number `json:"value,omitempty" hcl:"value" hcle:"omitempty"`
}

type Event struct {
	Q string `json:"q" hcl:"q" hcle:"omitempty"`
}
