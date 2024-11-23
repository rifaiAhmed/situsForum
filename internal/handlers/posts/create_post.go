package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rifaiAhmed/fastcampus/internal/model/posts"
)

func (h *Handler) CreatePost(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.CreatePostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId := c.GetInt64("userID")
	err := h.postService.CreatePost(ctx, userId, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
}
