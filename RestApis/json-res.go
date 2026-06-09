package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getUser(
	w http.ResponseWriter,
	r *http.Request,
) {

	user := User{
		Name: "Prakhar",
		Age:  23,
	}

	json.NewEncoder(w).Encode(user)
}
