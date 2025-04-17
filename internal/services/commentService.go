package services

import (
	"errors"
	"milkyway/internal/models"
)

type CommentService struct{}

var CId = 3

var comments = []models.Comment{
	{Id: 1, PostId: 1, UserId: 2, Text: "Why cant we just hug it out, pinche radiant"},
	{Id: 2, PostId: 2, UserId: 2, Text: "ancara ancara ancara, espada bankai"},
}

//	func GetCommentsByUserId(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json")
//
//		userId, err := strconv.Atoi(r.URL.Query().Get("userId"))
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		for _, comment := range comments {
//			if comment.UserId == userId {
//				err := json.NewEncoder(w).Encode(comment)
//				if err != nil {
//					log.Fatal(err)
//				}
//			}
//		}
//	}
//
//	func GetCommentsByPostId(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json")
//
//		postId, err := strconv.Atoi(r.URL.Query().Get("postId"))
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		for _, comment := range comments {
//			if comment.PostId == postId {
//				err := json.NewEncoder(w).Encode(comment)
//				if err != nil {
//					log.Fatal(err)
//				}
//			}
//		}
//	}
func (c *CommentService) UpdateComment(id int, commentEdit models.CommentEdit) (models.Comment, error) {
	for i, comment := range comments {
		if comment.Id == id {
			updatedComment := models.Comment{
				Id:     CId,
				PostId: commentEdit.PostId,
				UserId: commentEdit.UserId,
				Text:   commentEdit.Text,
			}

			comments[i] = updatedComment
			return updatedComment, nil
		}
	}
	return models.Comment{}, errors.New("Комментарий не найден")

}

func (c *CommentService) DeleteComment(id int) error {
	for i, comment := range comments {
		if comment.Id == id {
			comments = append(comments[:i], comments[i+1:]...)
			return nil
		}
	}
	return errors.New("Комментарий не найден")
}

//func UpdateComment(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	var commentEdit models.CommentEdit
//	commentId, err := strconv.Atoi(r.URL.Query().Get("commentId"))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	body, err := io.ReadAll(r.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = json.Unmarshal(body, &commentEdit)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	updatedComment := models.Comment{
//		Text:   commentEdit.Text,
//		UserId: commentEdit.UserId,
//		PostId: commentEdit.PostId,
//	}
//
//	for i, comment := range comments {
//		if comment.Id == commentId {
//			updatedComment.Id = commentId
//			comments = append(comments[:i], comments[i+1:]...)
//			comments = append(comments, updatedComment)
//			break
//		}
//	}
//	err = json.NewEncoder(w).Encode(updatedComment)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func (c *CommentService) CreateComment(commentEdit models.CommentEdit) models.Comment {
	newComment := models.Comment{
		Id:     CId,
		PostId: commentEdit.PostId,
		UserId: commentEdit.UserId,
		Text:   commentEdit.Text,
	}

	CId++
	comments = append(comments, newComment)

	return newComment
}

//
//func CreateComment(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	var commentEdit models.CommentEdit
//
//	body, err := io.ReadAll(r.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = json.Unmarshal(body, &commentEdit)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	newComment := models.Comment{
//		Id:     CId,
//		UserId: 2,
//		PostId: 1,
//		Text:   commentEdit.Text,
//	}
//
//	PId += 1
//
//	comments = append(comments, newComment)
//
//	err = json.NewEncoder(w).Encode(newComment)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//func DeleteComment(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "apllication/json")
//
//	commentId, err := strconv.Atoi(r.URL.Query().Get("id"))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for i, comment := range comments {
//		if comment.Id == commentId {
//			comments = append(comments[:i], comments[i+1:]...)
//			break
//		}
//	}
//
//	err = json.NewEncoder(w).Encode(comments)
//}
