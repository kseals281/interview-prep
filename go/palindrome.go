//	Write a function to determine if a word (or phrase) is a palindrome.
//
//	Then write a second function to receive a word (or phrase) list and determine which word is the longest palindrome.
//
//	See the template code / docstrings below for further requirements as well as the tests
//
//	A palindrome is a word, phrase, number, or other sequence of characters
//	which reads the same backward as forward
package main

import (
	"fmt"
	"io/ioutil"
	s "strings"
	"unicode"
)

const DICTIONARY = "setup/palindrome_dictionary.txt"

func errCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %e", msg, err)
		panic(err)
	}
}

func loadDictionary(dct string) <-chan string {
	//	Load dictionary and return via channel
	c := make(chan string)
	f, err := ioutil.ReadFile(dct)
	errCheck("Unable to open file", err)
	dat := string(f)
	go func() {
		defer close(c)
		for _, word := range s.Split(dat, "\n") {
			c <- word
		}
	}()

	return c
}

func isPalindrome(word string) bool {
	//	Return if word is palindrome, 'madam' would be one.
	//	Case insensitive, so Madam is valid too.
	//	It should work for phrases too so strip all but alphanumeric chars.
	//	So "No 'x' in 'Nixon'" should pass (see tests for more)
	var lettersOnly string
	for _, letter := range word {
		if unicode.IsLetter(letter) {
			lettersOnly += s.ToLower(string(letter))
		}
	}
	var firstHalf string
	var secondHalf string
	if len(lettersOnly)%2 == 0 {
		firstHalf = lettersOnly[:len(lettersOnly)/2]
		secondHalf = lettersOnly[len(lettersOnly)/2:]
	} else {
		firstHalf = lettersOnly[:(len(lettersOnly) / 2)]
		secondHalf = lettersOnly[len(lettersOnly)/2+1:]
	}
	return firstHalf == reversed(secondHalf)
}

func longest(words []string) string {
	//	Given a list of words return the longest palindrome
	//	If called without argument use the load_dictionary helper
	//	to populate the words list
	if len(words) == 0 {
		loadDictionary(DICTIONARY)
	}
	longest := ""
	for _, word := range words {
		if isPalindrome(word) {
			if len(word) > len(longest) {
				longest = word
			}
		}
	}
	return longest
}

func reversed(s string) string {
	var rev string
	for _, r := range s {
		rev = string(r) + rev
	}
	return rev
}
