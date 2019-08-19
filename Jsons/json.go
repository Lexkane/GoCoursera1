package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int
	Username string
	phone    string
}

var jsonStr = `{"id":42,"username":"lexkane","phone":"123"}`

func main() {
	data := []byte(jsonStr)
	u := &User{}
	json.Unmarshal(data, u)
	fmt.Printf("struct:\n\t%#v\n\n", u)
}
