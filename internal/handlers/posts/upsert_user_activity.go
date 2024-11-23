package posts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rifaiAhmed/fastcampus/internal/model/posts"
)

func (h *Handler) UpSertUserActivity(c *gin.Context) {
	ctx := c.Request.Context()
	var request posts.UserActivityRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userId := c.GetInt64("userID")
	postIDstr := c.Param("postID")
	postID, err := strconv.ParseInt(postIDstr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = h.postService.UpSertUserActivity(ctx, postID, userId, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}
