package convert

import "encoding/json"

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
	Viz       string    `json:"viz" hcl:"viz"`
	AutoScale bool      `json:"autoscale,omitempty" hcl:"autoscale" hcle:"omitempty"`
	Precision string    `json:"precision,omitempty" hcl:"precision" hcle:"omitempty"`
	Requests  []Request `json:"requests" hcl:"request"`
}

type GraphStyle struct {
	Palette     string `json:"palette,omitempty" hcl:"palette" hcle:"omitempty"`
	PaletteFlip bool   `json:"palette_flip,omitempty" hcl:"palette_flip" hcle:"omitempty"`
}

type Request struct {
	Q string `json:"q" hcl:"q"`
	//Aggregator string       `json:"aggregator,omitempty" hcl:"aggregator" hcle:"omitempty"`
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
