package api

import "github.com/gin-gonic/gin"

func StartServer() {
	r := gin.Default()
	r.POST("/query", QueryHandler)
	_ = r.Run(":8001")
}
