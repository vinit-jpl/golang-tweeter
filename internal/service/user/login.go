package user

import (
	"context"
	"errors"
	"go-tweets/internal/dto"
	"go-tweets/pkg/jwt"
	"net/http"

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
		return  "", "", http.StatusInternalServerError, err 
	}

	// get refresh token if exists

	// if not generate refresh token and save in db

	// return
}
