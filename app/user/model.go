package user

import (
	"github.com/gin-gonic/gin"
)

// New -
func New(e *gin.Engine) {
	e.GET("/users", list)
	e.GET("/users/:user_id", detail)
	e.POST("/users", create)
	e.PUT("/users/:user_id", update)
	e.DELETE("/users/:user_id", delete)
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
