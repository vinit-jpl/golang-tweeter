package middleware

import (
	"errors"
	"go-tweets/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secretkey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		header := ctx.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)

		if header == "" {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		userID, username, err := jwt.ValidateToken(header, secretkey, true)

		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		ctx.Set("userID", userID)
		ctx.Set("username", username)
		ctx.Next()

	}
}

func AuthRefreshTokenMiddleware(secretkey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		header := ctx.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)

		if header == "" {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		userID, username, err := jwt.ValidateToken(header, secretkey, false)

		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		ctx.Set("userID", userID)
		ctx.Set("username", username)
		ctx.Next()

	}
}
