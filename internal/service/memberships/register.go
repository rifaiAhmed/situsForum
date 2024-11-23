package memberships

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/rifaiAhmed/fastcampus/internal/model/memberships"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New("user or email already exist")
	}
	fmt.Println("===============")
	fmt.Println(req.Password)
	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	now := time.Now()
	model := memberships.UserModel{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(pass),
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: "",
		UpdatedBy: "",
	}
	return s.membershipRepo.CreateUser(ctx, model)
}
