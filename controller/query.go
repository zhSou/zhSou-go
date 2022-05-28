package controller

import "github.com/gin-gonic/gin"

/// 高亮显示请求字段
type highLight struct {
	PreTag  string `json:"preTag"`
	PostTag string `json:"postTag"`
}

/// 查询请求
type queryRequest struct {
	Query     string    `json:"query"`
	Page      int       `json:"page"`
	Limit     int       `json:"limit"`
	Order     string    `json:"order"`
	HighLight highLight `json:"highLight"`
}

/// 查询结果响应
type queryResponse struct {
	Id       int                    `json:"id"`
	Text     string                 `json:"text"`
	Score    float32                `json:"score"`
	Document map[string]interface{} `json:"document"`
}

func QueryHandler(c *gin.Context) {
	c.JSON(200, *BuildSuccessResponse(queryResponse{}))
}
