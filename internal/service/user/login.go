package user

import (
	"context"
	"errors"
	"go-tweets/internal/dto"
	"go-tweets/internal/model"
	"go-tweets/pkg/jwt"
	refreshtoken "go-tweets/pkg/refreshToken"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Login(ctx context.Context, req *dto.LoginRequest) (string, string, int, error) {

	// user is registerd or not
	userExists, err := s.userRepo.GetUserByEmailOrUsername(ctx, req.Email, "")

	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if userExists == nil {
		return "", "", http.StatusNotFound, errors.New("wrong email or password")

	}

	err = bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(req.Password))

	if err != nil {
		return "", "", http.StatusNotFound, errors.New("wrong password")
	}

	// generate token
	token, err := jwt.CreateToken(userExists.ID, userExists.Username, s.cfg.SecretJwt)

	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	// get refresh token if exists
	now := time.Now()
	refreshTokenExists, err := s.userRepo.GetRefreshToken(ctx, userExists.ID, now)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if refreshTokenExists != nil {
		return token, refreshTokenExists.RefreshToken, http.StatusOK, nil
	}

	// if not generate refresh token and save in db
	refreshToken, err := refreshtoken.GenerateRefreshToken()
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	err = s.userRepo.StoreRefreshToken(ctx,
		&model.RefreshTokenModel{
			UserID:       userExists.ID,
			RefreshToken: refreshToken,
			CreatedAt:    now,
			UpdatedAt:    now,
			ExpiredAt:    time.Now().Add(7 * 24 * time.Hour),
		})

	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	// return
	return token, refreshToken, http.StatusOK, nil

}
