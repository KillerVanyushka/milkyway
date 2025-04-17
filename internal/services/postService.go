package services

import (
	"errors"
	"milkyway/internal/models"
)

type PostService struct{}

var PId = 3

var posts = []models.Post{
	{Id: 1, UserId: 1, Content: "Meet me on the rood"},
	{Id: 2, UserId: 1, Content: "Lorem ipsum dolor sit amet"},
}

func (p *PostService) GetAllPosts() []models.Post {
	return posts
}

//func GetAllPosts(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	err := json.NewEncoder(w).Encode(posts)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func (p *PostService) GetPostById(id int) (models.Post, error) {
	for _, post := range posts {
		if post.Id == id {
			return post, nil
		}
	}
	return models.Post{}, errors.New("Пос не найден")
}

//
//func GetPostById(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	postId, err := strconv.Atoi(r.URL.Query().Get("id"))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, post := range posts {
//		if post.Id == postId {
//			err := json.NewEncoder(w).Encode(post)
//			if err != nil {
//				log.Fatal(err)
//			}
//			break
//		}
//	}
//
//}

func (p *PostService) CreatePost(postEdit models.PostEdit) models.Post {
	newPost := models.Post{
		Id:      PId,
		UserId:  postEdit.UserId,
		Content: postEdit.Content,
	}

	PId++
	posts = append(posts, newPost)
	return newPost

}

//func CreatePost(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	var postCreate models.PostEdit
//
//	body, err := io.ReadAll(r.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = json.Unmarshal(body, &postCreate)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	newPost := models.Post{
//		Id:      PId,
//		UserId:  2,
//		Content: postCreate.Content,
//	}
//
//	PId += 1
//
//	posts = append(posts, newPost)
//
//	err = json.NewEncoder(w).Encode(newPost)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func (p *PostService) UpdatePost(id int, postEdit models.PostEdit) (models.Post, error) {
	for i, post := range posts {
		if post.Id == id {
			updatedPost := models.Post{
				Id:      post.Id,
				UserId:  postEdit.UserId,
				Content: postEdit.Content,
			}

			posts[i] = updatedPost
			return updatedPost, nil
		}

	}
	return models.Post{}, errors.New("Пост не найден")
}

//func UpdatePost(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	var postEdit models.PostEdit
//	postId, err := strconv.Atoi(r.URL.Query().Get("id"))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	body, err := io.ReadAll(r.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = json.Unmarshal(body, &postEdit)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	updatedPost := models.Post{
//		Content: postEdit.Content,
//	}
//
//	for i, post := range posts {
//		if post.Id == postId {
//			updatedPost.Id = post.Id
//			posts = append(posts[:i], posts[i+1:]...)
//			posts = append(posts, updatedPost)
//			break
//		}
//	}
//	err = json.NewEncoder(w).Encode(posts)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func (p *PostService) DeletePost(id int) error {
	for i, post := range posts {
		if post.Id == id {
			posts = append(posts[:i], posts[i+1:]...)
			return nil
		}
	}
	return errors.New("Пост не найден")
}

//func DeletePost(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	postId, err := strconv.Atoi(r.URL.Query().Get("id"))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for i, post := range posts {
//		if post.Id == postId {
//			posts = append(posts[:i], posts[i+1:]...)
//			break
//		}
//	}
//
//	err = json.NewEncoder(w).Encode(posts)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
