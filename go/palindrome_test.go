package main

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		word string
		want bool
	}{
		{"Simple Odd Palindrome", "madam", true},
		{"Simple Even Palindrome", "poop", true},
		{"Mixed Case Simple Odd Palindrome", "MAdam", true},
		{"Complex Palindrome", "No 'x' in 'Nixon'", true},
		{"Non-Palindrome", "foo", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isPalindrome(tc.word)
			if got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func Test_loadDictionary(t *testing.T) {
	tests := []struct {
		name string
		c    <-chan string
		dct  string
		want []string
	}{
		{
			"Empty dictionary",
			make(chan string),
			"test_dictionaries/empty.txt",
			[]string{""},
		}, {
			"Non-Empty dictionary",
			make(chan string),
			"test_dictionaries/two_lines.txt",
			[]string{"one", "two"},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var got []string
			tc.c = loadDictionary(tc.dct)
			for word := range tc.c {
				got = append(got, word)
			}
			for i := range got {
				if tc.want[i] != got[i] {
					t.Errorf("got %s, want %s", got, tc.want)
				}
			}
		})
	}
}

func Test_longest(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  string
	}{
		{
			"Simple Case",
			[]string{"racecar"},
			"racecar",
		}, {
			"No Palindrome",
			[]string{"foo"},
			"",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := longest(tc.words)
			if got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}

func Test_reversed(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			"Empty",
			"",
			"",
		}, {
			"All Lower",
			"foo",
			"oof",
		}, {
			"All Upper",
			"FOO",
			"OOF",
		}, {
			"Mixed Cases",
			"FoO",
			"OoF",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := reversed(tc.s); got != tc.want {
				t.Errorf("reversed() = %v, want %v", got, tc.want)
			}
		})
	}
}
