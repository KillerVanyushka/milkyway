package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"milkyway/internal/models"
	"net/http"
	"strconv"
)

var UId = 3

var users = []models.User{
	{Id: 1, Username: "Zoro", PhoneNumber: "8-777-777-77-77", Age: 19, Gender: "Male", Email: "zoro@gmail.com", Password: "123456"},
	{Id: 2, Username: "Robbin", PhoneNumber: "8-776-777-77-77", Age: 19, Gender: "Female", Email: "robbin@gmail.com", Password: "123456"},
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		if user.Id == id {
			err := json.NewEncoder(w).Encode(user)
			if err != nil {
				log.Fatal(err)
			}
			break
		}
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userCreate models.UserEdit

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &userCreate)
	if err != nil {
		log.Fatal(err)
	}

	newUser := models.User{
		Id:          UId,
		Username:    userCreate.Username,
		Password:    userCreate.Password,
		Age:         userCreate.Age,
		Email:       userCreate.Email,
		Gender:      userCreate.Gender,
		PhoneNumber: userCreate.PhoneNumber,
	}

	UId += 1

	users = append(users, newUser)

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Fatal(err)
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userEdit models.UserEdit
	userId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &userEdit)
	if err != nil {
		log.Fatal(err)
	}

	updatedUser := models.User{
		Username:    userEdit.Username,
		Password:    userEdit.Password,
		Age:         userEdit.Age,
		Email:       userEdit.Email,
		Gender:      userEdit.Gender,
		PhoneNumber: userEdit.PhoneNumber,
	}

	for i, user := range users {
		if user.Id == userId {

			updatedUser.Id = user.Id
			users = append(users[:i], users[i+1:]...)
			users = append(users, updatedUser)
			break
		}
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Fatal(err)
	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	for i, user := range users {
		if user.Id == userId {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Fatal(err)
	}
}
