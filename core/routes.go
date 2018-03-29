package core

import (
	"github.com/depsea/api/app/article"
	"github.com/depsea/api/app/category"
	"github.com/depsea/api/app/role"
	"github.com/depsea/api/app/user"
	"github.com/gin-gonic/gin"
)

// NewRoutes init routes
func NewRoutes(e *gin.Engine) {
	article.New(e)
	category.New(e)
	role.New(e)
	user.New(e)
}
