package handler

import (
	"crud/internal/core/interface/service"
	"crud/internal/core/model"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

type handlerPost struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	ImageURL string `json:"image"`
	Author   string `json:"author"`
	Likes    int    `json:"likes"`
}

func CreatePost(service service.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var post handlerPost

		login := c.GetString("user")

		if err := c.BindJSON(&post); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверное тело запроса"})

			return
		}

		post.Author = login

		id, err := service.CreatePost(c.Request.Context(), model.Post(post))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"post": id})
	}
}

func GetPost(service service.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		numberId, err := strconv.Atoi(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверно передан id поста"})

			return
		}

		post, err := service.GetPost(c.Request.Context(), numberId)

		if err != nil {
			slog.Error(err.Error())

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка получения поста"})

			return

		}

		c.JSON(http.StatusOK, handlerPost(post))

	}
}

func LikePost(service service.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		numberId, err := strconv.Atoi(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверно передан id поста"})

			return
		}

		post, err := service.LikePost(c.Request.Context(), numberId)

		if err != nil {
			slog.Error(err.Error())

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка получения поста"})

			return

		}

		c.JSON(http.StatusOK, handlerPost(post))

	}
}

func GetPosts(service service.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []handlerPost
		posts, err := service.GetPosts(c.Request.Context())

		if err != nil {
			slog.Error(err.Error())

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка получения постов"})

			return

		}

		for _, post := range posts {
			result = append(result, handlerPost(post))
		}
		c.JSON(http.StatusOK, posts)
	}
}

func ChangePost(service service.PostService) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")

		numberId, err := strconv.Atoi(id)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверно передан id поста"})

			return
		}

		var newPost handlerPost

		login := context.GetString("user")

		if err := context.BindJSON(&newPost); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверное тело запроса"})

			return
		}

		newPost.Author = login

		post, err := service.ChangePost(context.Request.Context(), numberId, model.Post(newPost))

		if err != nil {
			slog.Error(err.Error())

			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка изменения поста"})

			return
		}

		context.JSON(http.StatusOK, handlerPost(post))

	}
}
