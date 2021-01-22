package main

import "math/rand"

func rand7() int {
	return rand.Intn(7)
}

func rand5() int {
	num := rand7()
	if num <= 5 {
		return num
	}
	return rand5()
}
