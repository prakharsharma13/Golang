package main

import "fmt"

func changeNum(num int) {
	num = 5
	fmt.Println("Inchange num", num)
}

func pointer(num *int) {
	*num = 5
	fmt.Println("pointer num", num)
}
func main() {
	num := 10
	fmt.Println("Before change num", num)
	changeNum(num)
	fmt.Println("After change num", num)
}
