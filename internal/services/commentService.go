package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"milkyway/internal/models"
	"net/http"
	"strconv"
)

var CId = 3

var comments = []models.Comment{
	{Id: 1, PostId: 1, UserId: 2, Text: "Why cant we just hug it out, pinche radiant"},
	{Id: 2, PostId: 2, UserId: 2, Text: "ancara ancara ancara, espada bankai"},
}

func GetCommentsByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId, err := strconv.Atoi(r.URL.Query().Get("userId"))
	if err != nil {
		log.Fatal(err)
	}

	for _, comment := range comments {
		if comment.UserId == userId {
			err := json.NewEncoder(w).Encode(comment)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func GetCommentsByPostId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	postId, err := strconv.Atoi(r.URL.Query().Get("postId"))
	if err != nil {
		log.Fatal(err)
	}

	for _, comment := range comments {
		if comment.PostId == postId {
			err := json.NewEncoder(w).Encode(comment)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var commentEdit models.CommentEdit
	commentId, err := strconv.Atoi(r.URL.Query().Get("commentId"))
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &commentEdit)
	if err != nil {
		log.Fatal(err)
	}

	updatedComment := models.Comment{
		Text:   commentEdit.Text,
		UserId: commentEdit.UserId,
		PostId: commentEdit.PostId,
	}

	for i, comment := range comments {
		if comment.Id == commentId {
			updatedComment.Id = commentId
			comments = append(comments[:i], comments[i+1:]...)
			comments = append(comments, updatedComment)
			break
		}
	}
	err = json.NewEncoder(w).Encode(updatedComment)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var commentEdit models.CommentEdit

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &commentEdit)
	if err != nil {
		log.Fatal(err)
	}

	newComment := models.Comment{
		Id:     CId,
		UserId: 2,
		PostId: 1,
		Text:   commentEdit.Text,
	}

	PId += 1

	comments = append(comments, newComment)

	err = json.NewEncoder(w).Encode(newComment)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "apllication/json")

	commentId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	for i, comment := range comments {
		if comment.Id == commentId {
			comments = append(comments[:i], comments[i+1:]...)
			break
		}
	}

	err = json.NewEncoder(w).Encode(comments)
}
