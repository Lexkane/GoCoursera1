package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
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

func handle(w http.ResponseWriter, req *http.Request) {
	result := ""
	for i := 0; i < 100; i++ {
		currPost := &Post{ID: i, Text: "new post", Time: time.Now()}
		jsonRaw, _ := json.Marshal(currentPost)
		result += string(jsonRaw)
	}
	time.Sleep(3 * time.Millisecond)
	w.Write([]byte(result))
}

func main() {
	runtime.GOMAXPROCS(4)

	http.HandleFunc("/", handle)
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
