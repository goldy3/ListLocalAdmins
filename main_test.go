package main

import (
	"reflect"
	"testing"
)

// TestParsePowerShellOutput - testing the parsePowerShellOutput function.
func TestParsePowerShellOutput(t *testing.T) {
	// Define test cases
	tests := []struct {
		name   string
		output string
		want   [][]string
	}{
		// Add test cases here
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parsePowerShellOutput(tt.output); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePowerShellOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}
