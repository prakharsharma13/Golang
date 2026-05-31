package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most used construct in go
// +useful methods
func main() {
	//uniniialied slice is nil
	// var nums []int

	var nums = make([]int, 3, 5)
	fmt.Println(nums)

	//capacity -> max nums of elements can fit
	fmt.Println(cap(nums))

	nums = append(nums, 3) //append adds element to slice and returns new slice

	//copy func
	var nums1 = make([]int, 0, 5)
	var nums2 = make([]int, len(nums1))

	copy(nums2, nums1)

	//slice operator
	var slicenum = []int{1, 2, 3}

	fmt.Println(slicenum[0:2])

	//slice package
	var num1 = []int{1, 2}
	var num2 = []int{1, 2}

	fmt.Println(slices.Equal(num1, num2))
}
