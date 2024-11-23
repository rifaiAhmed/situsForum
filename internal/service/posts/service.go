package posts

import (
	"context"

	"github.com/rifaiAhmed/fastcampus/internal/configs"
	"github.com/rifaiAhmed/fastcampus/internal/model/posts"
)

type postRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, model posts.CommentModel) error
	UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	GetUserActivy(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error)
	GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllPostResponse, error)
}

type service struct {
	cfg      *configs.Config
	postRepo postRepository
}

func NewService(postRepo postRepository, config *configs.Config) *service {
	return &service{
		postRepo: postRepo,
		cfg:      config,
	}
}
