package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	ID     int
	Name   string
	Active bool
}

func main() {
	tmpl := template.Must(template.ParseFiles("users.html"))

	users := []User{
		User{1, "Ivan", true},
		User{2, "<i>Alex</i>", false},
		User{3, "Dave", true},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, struct{ Users []User }{users})
	})

	fmt.Println("starting server at:8080")
	http.ListenAndServe(":8080", nil)
}
