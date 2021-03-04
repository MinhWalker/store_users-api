package app

import (
	"github.com/MinhWalker/store_users-api/logger"
	"github.com/gin-gonic/gin"
)

var(
	router = gin.Default()
)

func StartApplication()  {
	mapUrls()

	logger.Log.Info("about to start application...")
	router.Run(":8080")
}