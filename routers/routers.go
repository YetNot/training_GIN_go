package routers

import (
	"Gin/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	apiV1Group := router.Group("/api/v1")
	apiV1Group.GET("article", handlers.GetAllArticles)
	apiV1Group.POST("article", handlers.PostAllArticles)
	apiV1Group.GET("article/:id", handlers.GetArticleById)
	apiV1Group.PUT("article/:id", handlers.UpdateArticleById)
	apiV1Group.DELETE("article/:id", handlers.DeleteArticleById)
	return router
}