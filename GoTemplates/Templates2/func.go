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

func IsUserOdd(u *User) bool {
	return u.ID%2 != 0
}

func main() {
	tmplFuncs := template.FuncMap{
		"OddUser": IsUserOdd,
	}

	temp, err := template.
		New("").
		Funcs(tmplFuncs).
		ParseFiles("func.html")

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
