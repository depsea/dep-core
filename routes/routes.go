package routes

import (
	"github.com/gin-gonic/gin"
)

// Register -
func Register(e *gin.Engine) {
	users := e.Group("/users")
	{
		users.GET("", GetUserList)
		users.GET("/:user_id", GetUserDetail)
	}

	user := e.Group("/user")
	{
		user.GET("/login", Login)
		user.POST("/sign", Sign)
	}

	categories := e.Group("/categories")
	{
		categories.GET("", GetCategoryList)
		categories.GET("/:category_id", GetCategoryDetail)
		categories.POST("", CreateCategory)
	}

	tags := e.Group("/tags")
	{
		tags.GET("", GetTagList)
		tags.GET("/:tag_id", GetTagDetail)
		tags.POST("", CreateTag)
	}

	articles := e.Group("/articles")
	{
		articles.GET("", GetArticleList)
		articles.GET("/:article_id", GetArticleDetail)
		articles.POST("", CreateArticle)
	}

}
