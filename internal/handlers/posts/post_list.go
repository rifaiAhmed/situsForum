package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllPost(c *gin.Context) {
	ctx := c.Request.Context()
	pageIndexStr := c.Query("pageIndex")
	sizeIndexStr := c.Query("pageSize")

	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("invalid page index").Error(),
		})
		return
	}

	pageSize, err := strconv.Atoi(sizeIndexStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("invalid page size").Error(),
		})
		return
	}

	response, err := h.postService.GetAllPost(ctx, pageSize, pageIndex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
