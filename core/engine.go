package core

import (
	"github.com/gin-gonic/gin"
)

// Engine -
type Engine struct {
}

// NewEngine -
func NewEngine() *gin.Engine {
	r := gin.Default()

	NewRoutes(r)

	return r
}
