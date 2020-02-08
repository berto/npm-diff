package cmd

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

var versions = []struct {
	V1     string
	V2     string
	Result bool
}{
	{V1: "~1.0.4", V2: "~1.0.2", Result: false},
	{V1: "~1.2.3", V2: "~1.0.2", Result: true},
	{V1: "^1.2.3", V2: "^1.2.2", Result: false},
	{V1: "^1.2.3", V2: "^1.0.2", Result: false},
	{V1: "^2.3", V2: "^2.50", Result: false},
	{V1: "^2.3.3", V2: "^2.0", Result: false},
	{V1: "^2.3.0", V2: "^1.5", Result: true},
	{V1: "2.3.0", V2: "1.5.3", Result: true},
	{V1: "2.3.0", V2: "2.3.1", Result: true},
	{V1: "2.3.0", V2: "2.3.0", Result: false},
	{V1: "2", V2: "2.3", Result: true},
	{V1: "2", V2: "2", Result: false},
}

func TestDiff(t *testing.T) {
	for _, tt := range versions {
		assert.Equal(t, diff(tt.V1, tt.V2), tt.Result)
	}
}
