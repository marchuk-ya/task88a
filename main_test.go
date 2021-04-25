package main

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		n    string
		want bool
	}{
		{"12345", true},
		{"0", false},
		{"111", true},
		{"-1", true},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf(tt.n)
		t.Run(testName, func(t *testing.T) {
			ans := task88a(tt.n)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
