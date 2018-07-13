package common

import (
	"src/github.com/gin-gonic/gin"
	"github.com/xblymmx/huzhi123/constant"
)

func SendErrJSON(msg string, code int, args ...interface{}) {
	if len(args) != 1 || len(args) != 2 {
		panic("one or two args are required")
	}

	var ctx *gin.Context
	var retCode = constant.Code.ERROR

	if len(args) == 1 {
		if c, ok := args[0].(*gin.Context); !ok {
			panic(constant.Msg.GinContextRequired)
		} else {
			ctx = c
		}
	} else if len(args) == 2 {
		codeNo, ok := args[0].(int)
		if !ok {
			panic("ret code should be int")
		}
		retCode = codeNo

		c, ok := args[1].(*gin.Context)
		if !ok {
			panic(constant.Msg.GinContextRequired)
		}
		ctx = c
	}

	ctx.JSON(code, gin.H{
		"code": retCode,
		"msg":  msg,
		"data": gin.H{},
	})

	ctx.Abort()
}

