package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("person is an adult")
	} else if age >= 12 {
		fmt.Println("person is a teenager")
	} else {
		fmt.Println("person is a kid")
	}

	var role = "admin"
	var hasPermission = true

	if role == "admin" || hasPermission {
		fmt.Println("Access granted")
	}

	//we can directly declare a variable in the if statement

	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 50 {
		fmt.Println("Grade: C")
	}

	// go doesn't have ternary operator so we have to use if else
}
