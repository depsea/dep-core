package category

import (
	"github.com/gin-gonic/gin"
)

// New -
func New(e *gin.Engine) {
	e.GET("/categories", list)
	e.GET("/categories/:category_id", detail)
	e.POST("/categories", create)
	e.PUT("/categories/:category_id", update)
	e.DELETE("/categories/:category_id", delete)
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
