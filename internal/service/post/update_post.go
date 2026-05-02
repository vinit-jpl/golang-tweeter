package post

import (
	"context"
	"errors"
	"fmt"
	"go-tweets/internal/dto"
	"go-tweets/internal/model"
	"net/http"
	"time"
)

func (s *postService) UpdatePost(ctx context.Context, req *dto.CreateOrUpdatePostRequest, postID int64, userID int64) (int, error) {

	// check if tweet is present in DB
	postExists, err := s.postRepo.GetPostByID(ctx, postID)
	fmt.Println("post exists: ", postExists)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	if postExists == nil {
		return http.StatusNotFound, errors.New("tweet not found")
	}

	if postExists.UserID != userID {
		return http.StatusNotFound, errors.New("tweet not found of the logged in user")
	}

	// if exists update the data
	err = s.postRepo.UpdatePost(ctx, &model.PostModel{
		Title:     req.Title,
		Content:   req.Content,
		UpdatedAt: time.Now(),
	}, postID)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	// retutrn
	return http.StatusOK, nil
}
