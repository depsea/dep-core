package app

import (
	"github.com/depsea/dep-core/app/article"
	"github.com/depsea/dep-core/app/role"
	"github.com/depsea/dep-core/app/tag"
	"github.com/depsea/dep-core/app/user"
	"github.com/gin-gonic/gin"
)

// NewRouter init routes
func NewRouter(e *gin.Engine) {
	article.New(e)
	tag.New(e)
	role.New(e)
	user.New(e)
}
