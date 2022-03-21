package repository

import (
	"moviesdemo/entity"
)

type MovieRepository interface {
	Find(id string) (*entity.Movie, error)
	GetAll() ([]entity.Movie, error)
	Save(movie *entity.Movie) (*entity.Movie, error)
}
