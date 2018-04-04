package core

import (
	"github.com/gin-gonic/gin"
)

// Result -
func Result(ctx *gin.Context, data interface{}, err ...interface{}) {
	code := 0
	message := "success"

	if len(err) > 0 {
		code = err[0].(int)
		message = "error"

		if len(err) == 2 {
			message = err[1].(string)
		}
	}

	if data == nil {
		data = gin.H{}
	}

	ctx.JSON(200, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

// BodyError -
func BodyError(ctx *gin.Context, args ...interface{}) {
	code := 41000
	message := "body parse error."
	data := gin.H{}

	Result(ctx, data, code, message)
}

// DBError -
func DBError(ctx *gin.Context, err error) {
	print(err)
	Result(ctx, nil, 50000, err.Error())
}
