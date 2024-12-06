package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rifaiAhmed/fastcampus/internal/configs"
	"github.com/rifaiAhmed/fastcampus/pkg/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT

	return func(ctx *gin.Context) {
		header := ctx.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == "" {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}
		userID, username, err := jwt.ValidateToken(header, secretKey)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}
		ctx.Set("userID", userID)
		ctx.Set("username", username)
		ctx.Next()
	}
}
