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
		router.Group("/post")
		{
			api.GET("/post", handler.GetPosts(postService))
			api.POST("/post", handler.CreatePost(postService), middleware.AuthMiddleware)
			api.GET("/post/:id", handler.GetPost(postService))
			api.PUT("/post/:id", handler.ChangePost(postService), middleware.AuthMiddleware)
			api.GET("/post/:id/like", handler.LikePost(postService))
			api.POST("/post/:id_post/comment", handler.CreateComment(comService))
		}
		api.GET("/comment/:id", handler.GetComment(comService))
	}

	return router
}
