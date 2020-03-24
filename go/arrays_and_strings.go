package main

import "strconv"

// Implement an algorithm to determine if a string has all unique characters.
func isUnique(s string) bool {
	letters := make(map[rune]bool)
	for _, r := range s {
		if _, ok := letters[r]; ok {
			return false
		}
		letters[r] = true
	}
	return true
}

// Given two strings, write a method to decide if one is a permutation of the other
func checkPermutation(s1, s2 string) bool {
	seen := make(map[rune]int)
	for _, r := range s1 {
		if _, ok := seen[r]; ok {
			seen[r]++
		} else {
			seen[r] = 1
		}
	}
	for _, r := range s2 {
		if _, ok := seen[r]; !ok {
			return false
		}
		seen[r]--
		if seen[r] == 0 {
			delete(seen, r)
		}
	}
	if len(seen) != 0 {
		return false
	}
	return true
}

// Implement a method to perform basic string compression using the counts of repeated characters. For example the
// string aabcccccaaa would become a2b1c5a3. If the compressed string would not become smaller than the original string,
// your method should return the original string. You can assume the string has only uppercase and lowercase letters (a-z)
func stringCompression(s string) string {
	seen := ""
	var curr rune
	prev := 0
	for i, r := range s {
		if i == 0 {
			curr = r
		}
		if r != curr {
			seen += string(curr) + strconv.Itoa(i-prev)
			curr, prev = r, i
		}
	}
	seen += string(curr) + strconv.Itoa(len(s)-prev)
	return seen
}
