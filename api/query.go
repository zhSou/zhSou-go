package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhSou/zhSou-go/global"
	"github.com/zhSou/zhSou-go/service"
	"github.com/zhSou/zhSou-go/util/cutpage"
	"math"
	"time"
)

/// 高亮显示请求字段
type highLight struct {
	PreTag  string `json:"preTag"`
	PostTag string `json:"postTag"`
}

/// 查询请求
type queryRequest struct {
	Query string `json:"query"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
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
	_ = c.BindJSON(&qr)

	ids := service.Search(qr.Query)
	pageIds, err := cutpage.CutPage[int](ids, qr.Page, qr.Limit)

	if err != nil {
		return
	}
	var records []queryResponseRecord
	for _, pageId := range pageIds {
		dataRecord, err := global.GetDataReader().Read(uint32(pageId))
		if err != nil {
			return
		}
		records = append(records, queryResponseRecord{
			Id:   pageId,
			Text: dataRecord.Text,
			Document: document{
				Url:  dataRecord.Url,
				Text: dataRecord.Text,
			},
		})
	}

	c.JSON(200, *BuildSuccessResponse(queryResponse{
		UseTime:   time.Since(startTime).Seconds(),
		Total:     len(ids),
		TotalPage: int(math.Ceil(float64(len(ids)) / float64(qr.Limit))),
		Page:      qr.Page,
		Records:   records,
	}))
}
