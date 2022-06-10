package api

import (
	"github.com/zhSou/zhSou-go/util/highlight"
	"math"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/zhSou/zhSou-go/global"
	"github.com/zhSou/zhSou-go/service"
	"github.com/zhSou/zhSou-go/util/cutpage"
)

/// TODO 高亮显示请求字段
type highLight struct {
	PreTag  string `json:"preTag"`
	PostTag string `json:"postTag"`
}

/// 查询请求
type queryRequest struct {
	Query       string    `json:"query"`
	Page        int       `json:"page"`
	Limit       int       `json:"limit"`
	FilterWords []string  `json:"filterWord"`
	HighLight   highLight `json:"highLight"`
}

type document struct {
	Url  string `json:"url"`
	Text string `json:"text"`
}

/// 查询结果响应
type queryResponseRecord struct {
	Id       int      `json:"id"`
	Text     string   `json:"text"`
	Document document `json:"document"`
}

type queryResponse struct {
	UseTime   float64               `json:"useTime"`
	Total     int                   `json:"total"`
	Page      int                   `json:"page"`
	TotalPage int                   `json:"totalPage"`
	Records   []queryResponseRecord `json:"records"`
}

func QueryHandler(c *gin.Context) {
	startTime := time.Now()
	var qr queryRequest
	err := c.BindJSON(&qr)
	if err != nil {
		c.JSON(200, BuildErrorResponse(RequestFormatError, "请求格式有误"))
	}

	ret := service.SearchWithFilter(qr.Query, qr.FilterWords)
	ids := ret.DocIds

	records := make([]queryResponseRecord, 0)
	for _, pageId := range cutpage.CutPage[int](ids, qr.Page, qr.Limit) {
		dataRecord, err := global.GetDataReader().Read(uint32(pageId))
		if err != nil {
			c.JSON(200, BuildErrorResponse(UnknownError, "数据读取异常："+strconv.Itoa(pageId)))
			return
		}
		records = append(records, queryResponseRecord{
			Id:   pageId,
			Text: highlight.HighLight(dataRecord.Text, ret.Words, qr.HighLight.PreTag, qr.HighLight.PostTag),
			Document: document{
				Url:  dataRecord.Url,
				Text: dataRecord.Text,
			},
		})
	}

	c.JSON(200, BuildSuccessResponse(queryResponse{
		UseTime:   time.Since(startTime).Seconds(),
		Total:     len(ids),
		TotalPage: int(math.Ceil(float64(len(ids)) / float64(qr.Limit))),
		Page:      qr.Page,
		Records:   records,
	}))
}
