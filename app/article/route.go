package article

import (
	"github.com/depsea/api/core"
	"github.com/depsea/api/middlewares"
	"github.com/gin-gonic/gin"
)

// New -
func New(e *gin.Engine) {
	e.GET("/articles", list)
	e.GET("/articles/:article_id", detail)
	e.POST("/articles", middlewares.Auth(), create)
	e.PUT("/articles/:article_id", middlewares.Auth(), update)
	e.DELETE("/articles/:article_id", middlewares.Auth(), delete)
}

func list(ctx *gin.Context) {

}

func detail(ctx *gin.Context) {
	_id := core.ParseIDParam(ctx, "article_id")

	core.Result(ctx, _id)
}

func create(ctx *gin.Context) {

}

func update(ctx *gin.Context) {

}

func delete(ctx *gin.Context) {

}
