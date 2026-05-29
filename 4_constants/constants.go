package main

import "fmt"

const age = 30 //const can declare outside the main func

func main() {
	const name = "Prakhar"

	//we can't assign a const a new value

	fmt.Println(age)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}
