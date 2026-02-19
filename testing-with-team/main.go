package main

import "log"

func main() {
	task1()
}

func task1() {
	ps := NewPostService("https://jsonplaceholder.typicode.com")
	posts := ps.GetBy(1)
	log.Println("posts:", posts)

	p := Post{
		UserId: 1,
		Title:  "",
		Body:   "",
	}

	ps.Create(p)
	log.Println("Created post:", p)
}
