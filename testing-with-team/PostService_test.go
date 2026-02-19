package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostService(t *testing.T) {
	t.Run("PostService.GetBy", func(t *testing.T) {
		id := 2
		response := fmt.Sprintf(`{
						"userId": 1,
						"id": %d,
						"title": "title",
						"body": "body"
					}`, id)
		expected := Post{
			UserId: 1,
			Id:     id,
			Title:  "title",
			Body:   "body",
		}
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s", r.Method, r.URL)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}))
		defer server.Close()

		ps := NewPostService(server.URL)

		actual := ps.GetBy(id)

		assert.Equal(t, expected, actual)
	})

	t.Run("PostService.Create", func(t *testing.T) {
		id := 2
		post := Post{
			UserId: 1,
			//Id:     id, will be generated
			Title: "title",
			Body:  "body",
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s", r.Method, r.URL)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(fmt.Sprintf(`{
				"userId": 1,
				"id": %d,
				"title": "title",
				"body": "body"
				}`, id),
			))
		}))
		defer server.Close()

		ps := NewPostService(server.URL)
		actual, err := ps.Create(post)

		assert.Nil(t, err)
		assert.Equal(t, id, actual.Id)
		// assert.Equal(t, post.Title, "title")...
	})
}
