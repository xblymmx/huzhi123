package article

import (
	"github.com/gin-gonic/gin"
	"github.com/xblymmx/huzhi123/model"
	"strconv"
)

func queryList(c *gin.Context) {
	var articles []model.Article
	//var categoryId int
	var pageNo int
	var pageSize int
	//var startTime string
	//var endTime string
	var err error

	if pageNo, err = strconv.Atoi(c.Query("pageNo")); err != nil {
		pageNo = 1
		err = nil
	}

	if pageNo < 1 {
		pageNo = 1
	}

	if pageSize, err = strconv.Atoi(c.Query("pageSize")); err != nil {
		pageSize = 20
	}

	if pageSize <= 0 {
		pageSize = 20
	}
	offset := (pageNo-1) * pageSize

	model.DB.Limit(pageSize).Offset(offset).Find(articles)
}