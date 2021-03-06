package app

import (
	"github.com/MinhWalker/store_users-api/src/Repository/user"
	"github.com/MinhWalker/store_users-api/src/controllers/ping"
	"github.com/MinhWalker/store_users-api/src/controllers/users"
	"github.com/MinhWalker/store_users-api/src/services"
)

func mapUrls()  {
	userHandler := users.NewUserHandler(services.NewUserService(user.NewUserRepository()))

	router.GET("/ping", ping.Ping)

	//Users API
	router.POST("/users/create", userHandler.Create)
	router.GET("/users/:user_id", userHandler.Get)
	router.PUT("/users/:user_id", userHandler.Update)
	router.PATCH("/users/:user_id", userHandler.Update)
	router.DELETE("/users/:user_id", userHandler.Delete)

	router.GET("/internal/users/search", userHandler.Search)
	router.POST("/users/login", userHandler.Login)
}