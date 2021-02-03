package main

import "fmt"

func permPalindrome(str string) bool {
	letterCount := make(map[rune]int)
	oddSeen := false
	for _, r := range str {
		if _, ok := letterCount[r]; !ok {
			letterCount[r] = 1
		} else {
			letterCount[r] += 1
		}
	}
	for _, i := range letterCount {
		if i%2 == 1 {
			if oddSeen {
				return false
			}
			oddSeen = true
		}
	}
	return true
}

func altPermPalindrome(str string) bool {
	unpaired := make(map[rune]bool)
	for _, r := range str {
		if _, ok := unpaired[r]; !ok {
			unpaired[r] = true
		} else {
			delete(unpaired, r)
		}
		fmt.Println(unpaired)
	}
	if len(unpaired) > 1 {
		return false
	}
	return true
}
