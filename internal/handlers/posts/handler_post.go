package posts

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rifaiAhmed/fastcampus/internal/middleware"
	"github.com/rifaiAhmed/fastcampus/internal/model/posts"
)

type postService interface {
	CreatePost(ctx context.Context, userId int64, req posts.CreatePostRequest) error
	CreateComment(ctx context.Context, UserID, PostID int64, request posts.CreateCommentRequest) error
	UpSertUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error
	GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPostResponse, error)
}

type Handler struct {
	*gin.Engine

	postService postService
}

func NewHandler(api *gin.Engine, postService postService) *Handler {
	return &Handler{
		Engine:      api,
		postService: postService,
	}
}

func (h *Handler) RegisterRouter() {
	route := h.Group("posts")
	route.Use(middleware.AuthMiddleware())

	route.POST("/create", h.CreatePost)
	route.POST("/comment/:postID", h.CreateComment)
	route.PUT("/user_activity/:postID", h.UpSertUserActivity)
	route.GET("/", h.GetAllPost)
}
