package main

import "testing"

func Test_rotate(t *testing.T) {
	tests := []struct {
		name string
		s    string
		n    int
		want string
	}{
		{"Simple Positive Case", "foo", 1, "ofo"},
		{"Simple Negative Case", "foo", -1, "oof"},
		{"Zero Rotate", "foo", 0, "foo"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := rotate(tc.s, tc.n)
			if got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}
