package main

import (
	"authentication/controller"
	router "authentication/http"
	userMongoRepo "authentication/repository/mongodb"
	service "authentication/service"
	"fmt"
	"net/http"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
)

func main() {

	const port string = ":5051"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Running")
	})
	userRepository, err := userMongoRepo.NewMongoRepository()
	if err != nil {
		fmt.Println("User repo initialize error")
	}
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	httpRouter.GET("/user", userController.GetAll)
	httpRouter.GET("/user/{username}", userController.GetByFullName)
	httpRouter.POST("/user", userController.Save)

	httpRouter.SERVE(port)
}
