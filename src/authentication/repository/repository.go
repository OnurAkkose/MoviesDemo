package repository

import (
	"authentication/entity"
)

type UserRepository interface {
	Find(username string) (*entity.User, error)
	GetAll() ([]entity.User, error)
	Save(user *entity.User) (*entity.User, error)
}
