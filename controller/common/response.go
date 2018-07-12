package common

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/xblymmx/huzhi123/constant"
)

func SendErrJSON(msg string, args ...interface{}) {
	if len(args) != 1 || len(args) != 2 {
		panic("args should be one or two")
	}

	var ctx *gin.Context
	var retCode = constant.ErrorCode.ERROR

	if len(args) == 1 {
		if c, ok := args[0].(*gin.Context); !ok {
			panic("gin.Context is required")
		} else {
			ctx = c
		}
	} else if len(args) == 2 {
		eNo, ok := args[0].(int)
		if !ok {
			panic("err code should be int")
		}
		retCode = eNo

		c, ok := args[1].(*gin.Context)
		if !ok {
			panic("gin.Context is required")
		}
		ctx = c
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": retCode,
		"msg":  msg,
		"data": gin.H{},
	})

	ctx.Abort()
}

