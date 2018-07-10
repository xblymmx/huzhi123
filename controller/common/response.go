package common

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/xblymmx/huzhi123/config"
)

func SendErrJSON(msg string, args ...interface{}) {
	if len(args) == 0 {
		panic("missing gin.Context")
	}

	var ctx *gin.Context
	var errNo = config.ErrorCode.ERROR

	if len(args) == 1 {
		if c, ok := args[0].(*gin.Context); !ok {
			panic("args should be gin.Context")
		} else {
			ctx = c
		}
	} else if len(args) == 2 {
		eNo, ok := args[0].(int)
		if !ok {
			panic("err code should be int")
		}
		errNo = eNo

		c, ok := args[1].(*gin.Context)
		if !ok {
			panic("missing gin.Context")
		}
		ctx = c
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": errNo,
		"msg": msg,
		"data": gin.H{},
	})

	ctx.Abort()
}

