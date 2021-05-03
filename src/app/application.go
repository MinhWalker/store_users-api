package app

import (
	"github.com/MinhWalker/store_users-api/src/logger"
	"github.com/gin-gonic/gin"
)

var(
	router = gin.Default()
)

func StartApplication()  {
	mapUrls()

	logger.Info("about to start application...")
	router.Run(":8081")
}