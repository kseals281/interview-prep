package main

import (
	"fmt"
	"testing"
)

func Test_rand5(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test 1"},
		{"test 2"},
		{"test 3"},
		{"test 4"},
		{"test 5"},
		{"test 6"},
		{"test 7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num := rand5()
			fmt.Println(fmt.Sprintf("%s: %d", tt.name, num))
			if num < 0 || num > 5 {
				t.Errorf("num not within range: %d", num)
			}
		})
	}
}
