package user

import (
	"src/github.com/gin-gonic/gin"
	"github.com/xblymmx/huzhi123/common"
	"github.com/xblymmx/huzhi123/constant"
	"net/http"
	"github.com/xblymmx/huzhi123/utils"
	"strings"
	"github.com/xblymmx/huzhi123/model"
)

const (
	activeDuration = 24 * 60 * 60
)

// user sign up
func SignUp(c *gin.Context) {
	type UserBind struct {
		UserName string `json:"user_name" binding:"required,min=4,max=20"`
		Password string `json:"password" binding:"required,min=4,max=20"`
		Email string `json:"email" binding:"required"`
	}

	var userBind UserBind
	var err error

	err = c.ShouldBindJSON(&userBind)
	if err != nil {
		common.SendErrJSON(constant.Msg.InvalidUserSignUpParameter, http.StatusBadRequest, c)
		return
	}

	userBind.UserName = strings.TrimSpace(utils.AvoidXSS(userBind.UserName))
	userBind.Email = strings.TrimSpace(userBind.Email)

	// check if user already exists
	var user model.User
	err = model.DB.Where("email = ? OR user_name = ?", userBind.Email, userBind.UserName).First(&user).Error
	if err == nil {
		if user.UserName == userBind.UserName {
			common.SendErrJSON(constant.Msg.UserNameExists, http.StatusBadRequest, c)
			return
		} else {
			common.SendErrJSON(constant.Msg.UserEmailExists, http.StatusBadRequest, c)
			return
		}
	}

	var newUser model.User
	encryptedPwd, err := utils.GenerateEncryptedPassword(userBind.Password)
	if err != nil {
		common.SendErrJSON(constant.Msg.ServerError, http.StatusInternalServerError, c)
		return
	}

	newUser.UserName = userBind.UserName
	newUser.Email = userBind.Email
	newUser.Password = encryptedPwd
	newUser.Status = constant.User.StatusNotActive
	newUser.Gender = constant.User.GenderUnknown
	newUser.Role = constant.User.RoleBasic

	err = model.DB.Create(&newUser).Error
	if err != nil {
		common.SendErrJSON(constant.Msg.CreateUserError, http.StatusInternalServerError, c)
		return
	}

	// todo: redis cache active user
	//redisConn := model.RedisPool.Get()
	//redisConn.Do("SET", )

	// use goroutine to send email
	go func() {
		//utils.SendMail()
	}()

	c.JSON(http.StatusOK, gin.H{
		"code": constant.Code.SUCCESS,
		"msg": constant.Msg.SUCCESS,
		"data": newUser,
	})

}
