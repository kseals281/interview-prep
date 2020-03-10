package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	nums := generateSlice()
	fmt.Println(mergeSort(nums))
}

func generateSlice() (nums []int) {
	for i := 0; i < 10; i++ {
		nums = append(nums, rand.Int()%100)
	}
	return
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	mid := int(len(nums) / 2)
	left := nums[:mid]
	right := nums[mid:]

	sortedLeft := mergeSort(left)
	sortedRight := mergeSort(right)

	return merge(sortedLeft, sortedRight)
}

func merge(left, right []int) (sorted []int) {
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			sorted = append(sorted, left[0])
			left = left[1:]
		} else {
			sorted = append(sorted, right[0])
			right = right[1:]
		}
	}

	if len(left) > 0 {
		sorted = append(sorted, left...)
	} else {
		sorted = append(sorted, right...)
	}
	return
}
