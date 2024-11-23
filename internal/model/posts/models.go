package posts

import "time"

type (
	CreatePostRequest struct {
		PostTitle   string   `json:"postTitle"`
		PostContent string   `json:"postContent"`
		PostHastags []string `json:"postHastags"`
	}
)

type (
	PostModel struct {
		ID          int64     `db:"id"`
		UserID      int64     `db:"user_id"`
		PostTitle   string    `db:"post_title"`
		PostContent string    `db:"post_content"`
		PostHastags string    `db:"post_hastags"`
		CreatedAt   time.Time `db:"created_at"`
		UpdatedAt   time.Time `db:"updated_at"`
		CreatedBy   string    `db:"created_by"`
		UpdatedBy   string    `db:"updated_by"`
	}
)

type (
	GetAllPostResponse struct {
		Data       []Post     `json:"data"`
		Pagination Pagination `json:"pagination"`
	}

	Post struct {
		ID          int64    `json:"id"`
		UserID      int64    `json:"userID"`
		Username    string   `json:"username"`
		PostTitle   string   `json:"postTitle"`
		PostContent string   `json:"postContent"`
		PostHastags []string `json:"postHastags"`
	}

	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}
)
