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
)

const DICTIONARY = "/home/kseals/go/src/interview-prep/setup/palindrome_dictionary.txt"

func errCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %e", msg, err)
		panic(err)
	}
}

func loadDictionary(c chan string) {
	//	Load dictionary and return via channel
	f, err := ioutil.ReadFile(DICTIONARY)
	errCheck("Unable to open file", err)
	dat := string(f)
	for _, word := range s.Split(dat, "\n") {
		c <- word
	}
	close(c)
}

func isPalindrome(word string) bool {
	//	Return if word is palindrome, 'madam' would be one.
	//	Case insensitive, so Madam is valid too.
	//	It should work for phrases too so strip all but alphanumeric chars.
	//	So "No 'x' in 'Nixon'" should pass (see tests for more)
	return true
}

func longest(words string) {
	//	Given a list of words return the longest palindrome
	//	If called without argument use the load_dictionary helper
	//	to populate the words list

}

func main() {
	c := make(chan string)
	go loadDictionary(c)

	for word := range c {
		go isPalindrome(word)
	}
}
