package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amonunt   float32
	status    string
	createdAt time.Time
}

func main() {
	myorder := order{
		id:        "1234",
		amonunt:   100.0,
		status:    "pending",
		createdAt: time.Now(),
	}

	fmt.Println(myorder)
}
