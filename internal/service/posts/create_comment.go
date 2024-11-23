package posts

import (
	"context"
	"strconv"
	"time"

	"github.com/rifaiAhmed/fastcampus/internal/model/posts"
)

func (s *service) CreateComment(ctx context.Context, UserID, PostID int64, request posts.CreateCommentRequest) error {
	now := time.Now()
	model := posts.CommentModel{
		PostID:         PostID,
		UserID:         UserID,
		CommentContent: request.CommentContent,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBy:      strconv.FormatInt(UserID, 10),
		UpdatedBy:      strconv.FormatInt(UserID, 10),
	}

	err := s.postRepo.CreateComment(ctx, model)
	if err != nil {
		return err
	}

	return nil
}
