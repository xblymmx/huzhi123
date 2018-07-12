package category

import (
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/xblymmx/huzhi123/model"
	"github.com/microcosm-cc/bluemonday"
	"github.com/xblymmx/huzhi123/controller/common"
	"github.com/xblymmx/huzhi123/constant"
)

func Save(c *gin.Context, isCreated bool) {
	var category model.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		common.SendErrJSON(constant.ErrorMsg.GinBindingError, c)
		return
	}

	category.Name = bluemonday.UGCPolicy().Sanitize(category.Name)
	category.Name = strings.TrimSpace(category.Name)

	if category.Name == "" {
		common.SendErrJSON(constant.ErrorMsg.InvalidCategoryName, c)
	}

	//if utf8.RuneCountInString(category.Name) > model.

	if category.ParentID != 0 {
		var parentCategory model.Category
		if err := model.DB.First(&parentCategory, category.ParentID).Error; err != nil {
			common.SendErrJSON(constant.ErrorMsg.InvalidCategoryParentID, c)
			return
		}
	}

	//var newCategory model.Category
	//if isCreated {
	//	if err := model.DB.Create(&category).Error; err != nil {
	//		common.SendErrJSON("error while create category", c)
	//	}
	//}
}
