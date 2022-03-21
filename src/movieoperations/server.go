package main

import (
	"fmt"
	"moviesdemo/controller"
	router "moviesdemo/http"
	movieMongoRepo "moviesdemo/repository/mongodb"
	service "moviesdemo/service"
	"net/http"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
)

func main() {

	const port string = ":5050"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Running")
	})
	movieRepository, err := movieMongoRepo.NewMongoRepository()
	if err != nil {
		fmt.Println("Movie repo initialize error")
	}
	movieService := service.NewMovieService(movieRepository)
	movieController := controller.NewMovieController(movieService)
	httpRouter.GET("/movies", movieController.GetAll)
	httpRouter.GET("/movie/{id}", movieController.GetById)
	httpRouter.POST("/movie", movieController.Save)

	httpRouter.SERVE(port)
}
