package http

import (
	"crud/internal/core/interface/service"
	"crud/internal/transport/handler"
	"crud/internal/transport/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(service service.AuthService, postService service.PostService, comService service.ComService) *gin.Engine {
	router := gin.New()

	router.POST("/register", handler.RegisterUser(service))
	router.POST("/authorization", handler.LogIn(service))

	api := router.Group("/api")
	{
		api.GET("/posts", handler.GetPosts(postService))
		api.POST("/post", middleware.AuthMiddleware, handler.CreatePost(postService))
		api.GET("/post/:id", handler.GetPost(postService))
		api.PUT("/post/:id", middleware.AuthMiddleware, handler.ChangePost(postService))
		api.GET("/post/:id/like", handler.LikePost(postService))
		api.POST("/post/comment/:id_post", middleware.AuthMiddleware, handler.CreateComment(comService))
		api.GET("/comment/:id", handler.GetComment(comService))
	}

	return router
}
