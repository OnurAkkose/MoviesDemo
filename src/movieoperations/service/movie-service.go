package service

import (
	"moviesdemo/entity"
	"moviesdemo/repository"
)

type MovieService interface {
	Save(movie *entity.Movie) (*entity.Movie, error)
	GetAll() ([]entity.Movie, error)
	GetById(id string) (*entity.Movie, error)
}

type movieService struct {
	repository repository.MovieRepository
}

var (
	movieRepo repository.MovieRepository
)

func NewMovieService(_repository repository.MovieRepository) MovieService {
	return &movieService{
		repository: _repository,
	}
}

func (movieService *movieService) Save(movie *entity.Movie) (*entity.Movie, error) {
	return movieService.repository.Save(movie)
}
func (movieService *movieService) GetAll() ([]entity.Movie, error) {
	return movieService.repository.GetAll()
}
func (movieService *movieService) GetById(id string) (*entity.Movie, error) {
	return movieService.repository.Find(id)
}
