package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int `json:"user_id,string`
	Username string
	Adress   string `json:",omitempty"`
	Company  string `json:"-"`
}

func main() {
	u := &User{
		ID:       42,
		Username: "lexkane",
		Adress:   "",
		Company:  "Home",
	}

	result, _ := json.Marshal(u)
}
