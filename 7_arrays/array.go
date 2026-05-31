package main

import "fmt"

//arrays are functional sequence of elements of same type

func main() {
	var nums [4]int //array of 4 integers - we decalre array like this

	nums[0] = 10 //assining values to array

	fmt.Println(nums[0])

	fmt.Println(len(nums)) //find length of array

	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	//2d arrays
	twoD := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(twoD)
}
