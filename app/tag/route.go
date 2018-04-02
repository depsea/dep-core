package tag

import (
	"github.com/gin-gonic/gin"
)

// New -
func New(e *gin.Engine) {
	e.GET("/tags", list)
	e.GET("/tags/:tag_id", detail)
	e.POST("/tags", create)
	e.PUT("/tags/:tag_id", update)
	e.DELETE("/tags/:tag_id", delete)
}

func list(ctx *gin.Context) {

}

func detail(ctx *gin.Context) {

}

func create(ctx *gin.Context) {

}

func update(ctx *gin.Context) {

}

func delete(ctx *gin.Context) {

}
