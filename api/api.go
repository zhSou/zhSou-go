package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhSou/zhSou-go/global"
)

func StartServer() {
	conf := global.Config
	r := gin.Default()
	r.POST("/query", QueryHandler)
	r.POST("/getDocuments", GetDocumentsHandler)
	_ = r.Run(fmt.Sprintf("%s:%d", conf.ListenIp, conf.ListenPort))
}
