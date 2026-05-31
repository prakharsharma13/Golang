package main

import "fmt"

func main() {
	nums := []int{6, 7, 8}

	for i, num := range nums {
		fmt.Println(num, i)
	}

	//range with maps

	m := map[string]string{"name": "golang", "type": "programming language"}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
