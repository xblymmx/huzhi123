package category

import (
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/xblymmx/huzhi123/model"
	"github.com/xblymmx/huzhi123/common"
	"github.com/xblymmx/huzhi123/constant"
	"net/http"
	"github.com/xblymmx/huzhi123/utils"
)

func save(c *gin.Context, isCreated bool) {
	var category model.Category
	var err error

	if err := c.ShouldBindJSON(&category); err != nil {
		common.SendErrJSON(constant.Msg.GinBindingError, http.StatusBadRequest, c)
		return
	}

	category.Name = utils.AvoidXSS(category.Name)
	category.Name = strings.TrimSpace(category.Name)

	if category.Name == "" {
		common.SendErrJSON(constant.Msg.InvalidCategoryName, http.StatusBadRequest, c)
	}

	//if utf8.RuneCountInString(category.Name) > model.

	if category.ParentID != 0 {
		var parentCategory model.Category
		if err := model.DB.First(&parentCategory, category.ParentID).Error; err != nil {
			common.SendErrJSON(constant.Msg.InvalidCategoryParentID, http.StatusInternalServerError, c)
			return
		}
	}

	var toUpdateCate model.Category
	if isCreated { // create new category
		if err := model.DB.Create(&category).Error; err != nil {
			common.SendErrJSON("error while create category", http.StatusInternalServerError, c)
			return
		}
	} else { // update category
		if err := model.DB.First(&toUpdateCate, category.ID).Error; err != nil {
			common.SendErrJSON(constant.Msg.CategoryNotExist, http.StatusInternalServerError, c)
			return
		}
		err = model.DB.Model(&toUpdateCate).Updates(&category).Error
		if err != nil {
			common.SendErrJSON(constant.Msg.UpdateCategoryError, http.StatusInternalServerError, c)
			return
		}
	}

	var categoryData model.Category
	if isCreated {
		categoryData = category
	} else {
		categoryData = toUpdateCate
	}

	c.JSON(http.StatusOK, gin.H{
		"code": constant.Code.SUCCESS,
		"msg": constant.Msg.SUCCESS,
		"data": categoryData,
	})
}

func Create(c *gin.Context) {
	save(c, true)
}

func Update(c *gin.Context) {
	save(c, false)
}

func List(c *gin.Context) {
	var categories []model.Category
	var err error

	err = model.DB.Order("sequence asc").Find(&categories).Error
	if err != nil {
		common.SendErrJSON(constant.Msg.QueryCategoryError, http.StatusInternalServerError, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": constant.Code.SUCCESS,
		"msg": constant.Msg.SUCCESS,
		"data": categories,
	})

}
