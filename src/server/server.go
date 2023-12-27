package server

import "github.com/gin-gonic/gin"

func GetServer() *gin.Engine {
	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 500 << 20 // 8 MiB

	apiRoutes(router.Group("/"))

	return router
}
