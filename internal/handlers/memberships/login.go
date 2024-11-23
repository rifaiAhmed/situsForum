package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rifaiAhmed/fastcampus/internal/model/memberships"
)

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, err := h.membershipService.Login(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	response := memberships.LoginResponse{
		AccessToken: token,
	}
	c.JSON(http.StatusOK, response)
}
