package services

import (
	"milkyway/internal/models"
)

type UserRepo interface {
	GetAll() ([]models.User, error)
	GetById(id int) (*models.User, error)
	Create(user *models.User) error
	Update(id int, user *models.UserEdit) error
	Delete(id int) error
}

type UserService struct {
	repo UserRepo
}

func NewUserService(userRepo UserRepo) *UserService {
	return &UserService{repo: userRepo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetUserById(id int) (*models.User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) CreateUser(username string, phonenumber string, age int, gender string, email string, password string) (*models.User, error) {
	user := &models.User{
		Username:    username,
		Password:    password,
		Gender:      gender,
		Age:         age,
		PhoneNumber: phonenumber,
		Email:       email,
	}
	err := s.repo.Create(user)
	return user, err
}

func (s *UserService) UpdateUser(id int, userEdit *models.UserEdit) (*models.User, error) {
	err := s.repo.Update(id, userEdit)
	if err != nil {
		return nil, err
	}
	return s.repo.GetById(id)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}

//
//var UId = 3
//
//var users = []models.User{
//	{Id: 1, Username: "Zoro", PhoneNumber: "8-777-777-77-77", Age: 19, Gender: "Male", Email: "zoro@gmail.com", Password: "123456"},
//	{Id: 2, Username: "Robbin", PhoneNumber: "8-776-777-77-77", Age: 19, Gender: "Female", Email: "robbin@gmail.com", Password: "123456"},
//}

//func GetAllUsers(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	err := json.NewEncoder(w).Encode(users)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//func (u *UserService) GetAllUsers() []models.User {
//	return users
//}

//func GetUserById(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	id, err := strconv.Atoi(r.URL.Query().Get("id"))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, user := range users {
//		if user.Id == id {
//			err := json.NewEncoder(w).Encode(user)
//			if err != nil {
//				log.Fatal(err)
//			}
//			break
//		}
//	}
//}
//
//func (u *UserService) GetUserById(id int) (models.User, error) {
//	for _, user := range users {
//		if user.Id == id {
//			return user, nil
//		}
//	}
//	return models.User{}, errors.New("Пользователь не найден")
//}
//
//
//func CreateUser(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	var userCreate models.UserEdit
//
//	body, err := io.ReadAll(r.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = json.Unmarshal(body, &userCreate)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	newUser := models.User{
//		Id:          UId,
//		Username:    userCreate.Username,
//		Password:    userCreate.Password,
//		Age:         userCreate.Age,
//		Email:       userCreate.Email,
//		Gender:      userCreate.Gender,
//		PhoneNumber: userCreate.PhoneNumber,
//	}
//
//	UId += 1
//
//	users = append(users, newUser)
//
//	err = json.NewEncoder(w).Encode(users)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//}
//
//func (u *UserService) CreateUser(userEdit models.UserEdit) models.User {
//	newUser := models.User{
//		Id:          UId,
//		Username:    userEdit.Username,
//		PhoneNumber: userEdit.PhoneNumber,
//		Age:         userEdit.Age,
//		Gender:      userEdit.Gender,
//		Email:       userEdit.Email,
//		Password:    userEdit.Password,
//	}
//
//	UId++
//	users = append(users, newUser)
//
//	return newUser
//}
//
////
////func UpdateUser(w http.ResponseWriter, r *http.Request) {
////	w.Header().Set("Content-Type", "application/json")
////
////	var userEdit models.UserEdit
////	userId, err := strconv.Atoi(r.URL.Query().Get("id"))
////	if err != nil {
////		log.Fatal(err)
////	}
////
////	body, err := io.ReadAll(r.Body)
////	if err != nil {
////		log.Fatal(err)
////	}
////
////	err = json.Unmarshal(body, &userEdit)
////	if err != nil {
////		log.Fatal(err)
////	}
////
////	updatedUser := models.User{
////		Username:    userEdit.Username,
////		Password:    userEdit.Password,
////		Age:         userEdit.Age,
////		Email:       userEdit.Email,
////		Gender:      userEdit.Gender,
////		PhoneNumber: userEdit.PhoneNumber,
////	}
////
////	for i, user := range users {
////		if user.Id == userId {
////
////			updatedUser.Id = user.Id
////			users = append(users[:i], users[i+1:]...)
////			users = append(users, updatedUser)
////			break
////		}
////	}
////
////	err = json.NewEncoder(w).Encode(users)
////	if err != nil {
////		log.Fatal(err)
////	}
////
////}
//
//func (u *UserService) UpdateUser(userId int, userEdit models.UserEdit) (models.User, error) {
//	for i, user := range users {
//		if user.Id == userId {
//			updatedUser := models.User{
//				Id:          user.Id,
//				Username:    userEdit.Username,
//				PhoneNumber: userEdit.PhoneNumber,
//				Age:         userEdit.Age,
//				Gender:      userEdit.Gender,
//				Email:       userEdit.Email,
//				Password:    userEdit.Password,
//			}
//
//			users[i] = updatedUser
//			return updatedUser, nil
//		}
//	}
//	return models.User{}, errors.New("Пользователь не найден")
//}
//
//func (u *UserService) DeleteUser(userId int) error {
//	for i, user := range users {
//		if user.Id == userId {
//			users = append(users[:i], users[i+1:]...)
//			return nil
//		}
//	}
//	return errors.New("Пользователь не найден")
//}
//
////
////func DeleteUser(w http.ResponseWriter, r *http.Request) {
////	w.Header().Set("Content-Type", "application/json")
////
////	userId, err := strconv.Atoi(r.URL.Query().Get("id"))
////	if err != nil {
////		log.Fatal(err)
////	}
////
////	for i, user := range users {
////		if user.Id == userId {
////			users = append(users[:i], users[i+1:]...)
////			break
////		}
////	}
////	err = json.NewEncoder(w).Encode(users)
////	if err != nil {
////		log.Fatal(err)
////	}
////}
