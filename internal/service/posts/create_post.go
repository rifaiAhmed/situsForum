package posts

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/rifaiAhmed/fastcampus/internal/model/posts"
)

func (s *service) CreatePost(ctx context.Context, userId int64, req posts.CreatePostRequest) error {
	postHastags := strings.Join(req.PostHastags, ",")

	now := time.Now()
	model := posts.PostModel{
		UserID:      userId,
		PostTitle:   req.PostTitle,
		PostContent: req.PostContent,
		PostHastags: postHastags,
		CreatedAt:   now,
		UpdatedAt:   now,
		CreatedBy:   strconv.FormatInt(userId, 10),
		UpdatedBy:   strconv.FormatInt(userId, 10),
	}

	err := s.postRepo.CreatePost(ctx, model)
	if err != nil {
		return err
	}

	return nil
}
