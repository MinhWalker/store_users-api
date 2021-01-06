package app

import (
	"github.com/MinhWalker/store_users-api/controllers/ping"
	"github.com/MinhWalker/store_users-api/controllers/users"
)

func mapUrls()  {
	router.GET("/ping", ping.Ping)

	router.GET("/user/:user_id", users.GetUser)
	router.POST("/user/create", users.CreateUser)
}