package article

import (
	"github.com/gin-gonic/gin"
)

// New -
func New(e *gin.Engine) {
	e.GET("/articles", list)
	e.GET("/articles/:article_id", detail)
	e.POST("/articles", create)
	e.PUT("/articles/:article_id", update)
	e.DELETE("/articles/:article_id", delete)
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
