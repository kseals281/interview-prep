package main

import "math"

//Write a function that rotates characters in a string, in both directions:
//
//if n is positive move characters from beginning to end, e.g.: rotate('hello', 2) would return llohe
//if n is negative move characters to the start of the string, e.g.: rotate('hello', -2) would return lohel
//See tests for more info. Have fun!

func rotate(s string, n int) string {
	// Expects string and number of spots to rotate
	for i := 0; i < int(math.Abs(float64(n))); i++ {
		if n < 0 {
			c := s[:1]
			s = s[1:]
			s += c
		} else {
			c := s[len(s)-1:]
			s = s[:len(s)-1]
			s = c + s
		}
	}
	return s
}
