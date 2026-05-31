package main

import "fmt"

func main() {
	//creating map

	m := make(map[string]string)

	//setting an element
	m["name"] = "golang"

	//get an element
	fmt.Println(m["name"])

	//if key doesn't exist in map then it return zero value for it

	//delete an element
	delete(m, "name")

	//clear a map
	clear(m)

	k := map[string]int{"one": 1, "two": 2}
	fmt.Println(k)

	_, ok := m["name"] //check if key exist in map

	if ok {
		fmt.Println("key exist")
	} else {
		fmt.Println("no exist")
	}
}
