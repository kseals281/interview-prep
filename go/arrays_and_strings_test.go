package main

import "testing"

func Test_isUnique(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			"unique string",
			"ached",
			true,
		}, {
			"non-unique string",
			"ccb",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnique(tt.s); got != tt.want {
				t.Errorf("isUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPermutation(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"simple permutation",
			args{
				s1: "foo",
				s2: "oof",
			},
			true,
		}, {
			"not permutation same length",
			args{
				s1: "foo",
				s2: "bar",
			},
			false,
		}, {
			"not permutation larger first",
			args{
				s1: "abcdefg",
				s2: "abcd",
			},
			false,
		}, {
			"not permutation larger second",
			args{
				s1: "abcd",
				s2: "abcdefg",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPermutation(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringCompression(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			"two letter compression",
			"aaaaaabbbbb",
			"a6b5",
		}, {
			"four sets of compression",
			"aabcccccaaa",
			"a2b1c5a3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringCompression(tt.s); got != tt.want {
				t.Errorf("stringCompression() = %v, want %v", got, tt.want)
			}
		})
	}
}
