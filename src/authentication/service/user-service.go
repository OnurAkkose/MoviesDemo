package service

import (
	"authentication/entity"
	"authentication/repository"
)

type UserService interface {
	Save(user *entity.User) (*entity.User, error)
	GetAll() ([]entity.User, error)
	GetByFullName(id string) (*entity.User, error)
}

type userService struct {
	repository repository.UserRepository
}

var (
	userRepo repository.UserRepository
)

func NewUserService(_repository repository.UserRepository) UserService {
	return &userService{
		repository: _repository,
	}
}

func (userService *userService) Save(user *entity.User) (*entity.User, error) {
	return userService.repository.Save(user)
}
func (userService *userService) GetAll() ([]entity.User, error) {
	return userService.repository.GetAll()
}
func (userService *userService) GetByFullName(username string) (*entity.User, error) {
	return userService.repository.Find(username)
}
