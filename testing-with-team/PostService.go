package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type PostService struct {
	baseUrl string
}

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func NewPostService(baseUrl string) *PostService {
	return &PostService{baseUrl: baseUrl}
}

func (p *PostService) GetBy(id int) Post {
	resp, err := http.Get(p.baseUrl + "/posts/" + strconv.Itoa(id))
	if err != nil {
		log.Fatal("PostService.GetBy: ", err)
	}
	defer resp.Body.Close()

	post := Post{}
	err = json.NewDecoder(resp.Body).Decode(&post)
	if err != nil {
		log.Fatal("Error decoding response body", err)
	}

	return post
}

func (p *PostService) Create(post Post) (Post, error) {
	body, _ := json.Marshal(post)

	resp, _ := http.Post(p.baseUrl+"/posts", "application/json", bytes.NewBuffer(body))
	defer resp.Body.Close()

	created := Post{}
	_ = json.NewDecoder(resp.Body).Decode(&created)

	return created, nil
}
