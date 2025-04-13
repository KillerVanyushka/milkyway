package services

import (
	"encoding/json"
	"io"
	"log"
	"milkyway/internal/models"
	"net/http"
	"strconv"
)

var PId = 3

var posts = []models.Post{
	{Id: 1, UserId: 1, Content: "Meet me on the rood"},
	{Id: 2, UserId: 1, Content: "Lorem ipsum dolor sit amet"},
}

func GetAllPosts(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(posts)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	postId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	for _, post := range posts {
		if post.Id == postId {
			err := json.NewEncoder(w).Encode(post)
			if err != nil {
				log.Fatal(err)
			}
			break
		}
	}

}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var postCreate models.PostEdit

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &postCreate)
	if err != nil {
		log.Fatal(err)
	}

	newPost := models.Post{
		Id:      PId,
		UserId:  2,
		Content: postCreate.Content,
	}

	PId += 1

	posts = append(posts, newPost)

	err = json.NewEncoder(w).Encode(newPost)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var postEdit models.PostEdit
	postId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &postEdit)
	if err != nil {
		log.Fatal(err)
	}

	updatedPost := models.Post{
		Content: postEdit.Content,
	}

	for i, post := range posts {
		if post.Id == postId {
			updatedPost.Id = post.Id
			posts = append(posts[:i], posts[i+1:]...)
			posts = append(posts, updatedPost)
			break
		}
	}
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		log.Fatal(err)
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	postId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	for i, post := range posts {
		if post.Id == postId {
			posts = append(posts[:i], posts[i+1:]...)
			break
		}
	}

	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		log.Fatal(err)
	}
}
