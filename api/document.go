package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhSou/zhSou-go/global"
	"strconv"
)

func GetDocumentsHandler(c *gin.Context) {
	var idList []int
	err := c.BindJSON(&idList)
	if err != nil {
		c.JSON(200, BuildErrorResponse(RequestFormatError, "请求格式有误"))
		return
	}

	docs := make(map[string]document)

	for _, id := range idList {
		if id < 0 || id > int(global.GetDataReader().Len()) {
			continue
		}
		doc, err := global.GetDataReader().Read(uint32(id))
		if err != nil {
			continue
		}
		docs[strconv.Itoa(id)] = document{
			Url:  doc.Url,
			Text: doc.Text,
		}
	}
	c.JSON(200, BuildSuccessResponse(gin.H{
		"documents": docs,
	}))
}
