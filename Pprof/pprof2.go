package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

//goroutine leakage
type Post struct {
	ID       int
	Text     string
	Author   string
	Comments int
	Time     time.Time
}

func getPost(out chan []Post) {
	posts := []Post{}
	for i := 1; i < 10; i++ {
		post := Post{ID: 1, Text: "text"}
		posts = appends(posts, post)
	}
	out <- posts
}

func handleLeak(w http.ResponseWriter, req *http.Request) {
	res := make(chan []Post)
	go getPost(res)
}

func main() {
	http.HandleFunc("/", handleLeak)
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
