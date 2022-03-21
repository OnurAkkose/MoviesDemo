package controller

import (
	"authentication/entity"
	"authentication/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UserController interface {
	GetByFullName(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Save(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userService service.UserService
}

func NewUserController(_userService service.UserService) UserController {
	return &userController{
		userService: _userService,
	}
}

func (userController *userController) Save(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return
	}
	result, err := userController.userService.Save(&user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
func (userController *userController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	users, err := userController.userService.GetAll()
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}
func (userController *userController) GetByFullName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	w.Header().Set("content-type", "application/json")
	user, err := userController.userService.GetByFullName(username)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}
