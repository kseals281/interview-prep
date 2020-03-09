package main

import (
	"fmt"
	"math"
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

	mid := int(math.Floor(float64(len(nums) / 2)))
	var left, right []int

	for i, j := range nums {
		if i >= mid {
			right = append(right, j)
		} else {
			left = append(left, j)
		}
	}

	sortedLeft := mergeSort(left)
	sortedRight := mergeSort(right)

	return merge(sortedLeft, sortedRight)
}

func merge(left, right []int) (sorted []int) {
	count := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			sorted = append(sorted, left[0])
			left = left[1:]
		} else {
			sorted = append(sorted, right[0])
			right = right[1:]
		}
		count++
	}

	if len(left) > 0 {
		sorted = append(sorted, left...)
	} else {
		sorted = append(sorted, right...)
	}
	return
}
