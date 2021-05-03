package app

import (
	"github.com/MinhWalker/store_users-api/src/controllers/ping"
	"github.com/MinhWalker/store_users-api/src/controllers/users"
)

func mapUrls()  {
	router.GET("/ping", ping.Ping)

	router.POST("/users/create", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)

	router.GET("/internal/users/search", users.Search)
	router.POST("/users/login", users.Login)
}