package role

import (
	"github.com/gin-gonic/gin"
)

// New -
func New(e *gin.Engine) {
	e.GET("/roles", list)
	e.GET("/roles/:role_id", detail)
	e.POST("/roles", create)
	e.PUT("/roles/:role_id", update)
	e.DELETE("/roles/:role_id", delete)
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
