package memberships

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rifaiAhmed/fastcampus/internal/model/memberships"
)

type membershipService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string, error)
}

type Handler struct {
	*gin.Engine

	membershipService membershipService
}

func NewHandler(api *gin.Engine, membershipService membershipService) *Handler {
	return &Handler{
		Engine:            api,
		membershipService: membershipService,
	}
}

func (h *Handler) RegisterRouter() {
	route := h.Group("memberships")
	route.GET("/ping", h.Ping)
	route.POST("/signup", h.SignUp)
	route.POST("/login", h.Login)
}
