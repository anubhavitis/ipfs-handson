package server

import (
	"fileverse-test/src/services/handlers"

	"github.com/gin-gonic/gin"
)

func apiRoutes(group *gin.RouterGroup) {
	group.GET(GetApiConst, handlers.GetFileHandler)
	group.POST(UploadApiConst, handlers.UploadFileHandler)
}
