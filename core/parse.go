package core

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ParseIDParam -
func ParseIDParam(ctx *gin.Context, field string) int {
	_id, err := strconv.Atoi(ctx.Param(field))

	if err != nil {
		Result(ctx, gin.H{}, 40000, fmt.Sprintf("%s param parse error.", field))
		return 0
	}

	return _id
}
