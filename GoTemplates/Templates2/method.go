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

func (u *User) PrintActive() string {
	if !u.Active {
		return ""
	}
	return "method says user" + u.Name + "active"
}

func main() {
	temp, err := template.New("").ParseFiles("method.html")
	if err != nil {
		panic(err)
	}
	users := []User{
		User{1, "Ivan", true},
		User{2, "Alex", false},
		User{3, "Dave", true},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "method.html",
			struct {
				Users []User
			}{
				users,
			})
	})

	fmt.Println("starting server at:8080")
	http.ListenAndServe(":8080", nil)
}
