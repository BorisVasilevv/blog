package handler

import (
	"crud/internal/core/interface/service"
	"crud/internal/core/model"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

type handlerComment struct {
	Id_post int    `json:"id_post"`
	Body    string `json:"body"`
	Author  string `json:"author"`
}

func CreateComment(service service.ComService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var comment handlerComment

		login := c.GetString("user")
		id_post := c.GetInt("id_post")

		if err := c.BindJSON(&comment); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверное тело запроса"})

			return
		}

		comment.Id_post = id_post
		comment.Author = login

		id, err := service.CreateComment(c.Request.Context(), model.Comment(comment))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"comment": id})
	}
}

func GetComment(service service.ComService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		numberId, err := strconv.Atoi(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверно передан id комментария"})

			return
		}

		сomment, err := service.GetComment(c.Request.Context(), numberId)

		if err != nil {
			slog.Error(err.Error())

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка получения комментария"})

			return

		}

		c.JSON(http.StatusOK, handlerComment(сomment))

	}
}
