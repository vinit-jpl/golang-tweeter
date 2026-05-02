package post

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func (s *postService) DeletePost(ctx context.Context, postID, userID int64) (int, error) {
	// check if  the tweet exists
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

	// if exist => soft delete

	err = s.postRepo.SoftDeletePost(ctx, postID, time.Now())

	if err != nil {
		return http.StatusInternalServerError, err
	}


	// return
	return http.StatusOK, nil
}
