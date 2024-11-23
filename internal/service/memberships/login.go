package memberships

import (
	"context"
	"errors"
	"fmt"

	"github.com/rifaiAhmed/fastcampus/internal/model/memberships"
	"github.com/rifaiAhmed/fastcampus/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "")
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("email not found")
	}
	fmt.Println(user.Password)
	fmt.Println(req.Password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("email or password not valid")
	}
	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", err
	}

	return token, nil
}
