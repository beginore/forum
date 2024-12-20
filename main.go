package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type post struct {
	ID   int
	Name string
	Body string
	Date string
	User string
}

var posts = []post{
	{
		ID:   1,
		Name: "Welcome to My Blog",
		Body: "This is the first post on my new blog. I will share my thoughts on various topics.",
		Date: "2024-10-19",
		User: "John Doe",
	},
	{
		ID:   2,
		Name: "Understanding Go Structs",
		Body: "In this post, we will dive into how structs work in Go and how to use them efficiently.",
		Date: "2024-10-20",
		User: "Jane Smith",
	},
	{
		ID:   3,
		Name: "Working with HTTP in Go",
		Body: "Let's explore how to build a simple web server in Go using the net/http package.",
		Date: "2024-10-21",
		User: "DevGuru",
	},
}

func main() {
	http.HandleFunc("GET /", homeHandler)
	http.HandleFunc("GET /post/", postHandler)
	log.Println("The server listened at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "sdsds")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// if
	tmp, err := template.ParseFiles("./templates/home.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tmp.Execute(w, posts)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func postHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/post/")
	num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Fprint(w, posts[num-1])
	tmp, err := template.ParseFiles("./templates/post.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tmp.Execute(w, posts[num-1])
	if err != nil {
		fmt.Println(err)
		return
	}
}
