package memberships

import (
	"context"

	"github.com/rifaiAhmed/fastcampus/internal/configs"
	"github.com/rifaiAhmed/fastcampus/internal/model/memberships"
)

type membershipRepo interface {
	GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
}

type service struct {
	cfg            *configs.Config
	membershipRepo membershipRepo
}

func NewService(membershipRepo membershipRepo, config *configs.Config) *service {
	return &service{
		membershipRepo: membershipRepo,
		cfg:            config,
	}
}
