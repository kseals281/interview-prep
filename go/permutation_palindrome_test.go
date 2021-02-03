package main

import "testing"

func Test_permPalindrome(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			"palindrome",
			"civic",
			true,
		}, {
			"permutation palindrome",
			"cciiv",
			true,
		}, {
			"non palindrome",
			"cdcdssfg",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permPalindrome(tt.str); got != tt.want {
				t.Errorf("permPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_altPermPalindrome(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			"palindrome",
			"civic",
			true,
		}, {
			"permutation palindrome",
			"cciiv",
			true,
		}, {
			"non palindrome",
			"cdcdssfg",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := altPermPalindrome(tt.str); got != tt.want {
				t.Errorf("altPermPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
