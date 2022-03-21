package controller

import (
	"encoding/json"
	"moviesdemo/entity"
	"moviesdemo/service"
	"net/http"

	"github.com/gorilla/mux"
)

type MovieController interface {
	GetById(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Save(w http.ResponseWriter, r *http.Request)
}

type movieController struct {
	movieService service.MovieService
}

func NewMovieController(_movieService service.MovieService) MovieController {
	return &movieController{
		movieService: _movieService,
	}
}

func (movieController *movieController) Save(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var movie entity.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		return
	}
	result, err := movieController.movieService.Save(&movie)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
func (movieController *movieController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	movies, err := movieController.movieService.GetAll()
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)

}
func (movieController *movieController) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["id"]
	w.Header().Set("content-type", "application/json")
	movie, err := movieController.movieService.GetById(movieId)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)

}
