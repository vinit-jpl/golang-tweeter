package post

import (
	"context"
	"errors"
	"go-tweets/internal/model"
	"net/http"
	"time"
)

func (s *postService) LikeOrUnlikePost(ctx context.Context, postID, userID int64) (int, error) {

	// check if post or tweet exists or not
	postExists, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if postExists == nil {
		return http.StatusNotFound, errors.New("cannot like as tweet does not exists")
	}

	// check user already like or not
	isUserAlreadyLikePost, err := s.postRepo.IsUserAlreadyLikePost(ctx, postExists.ID, postExists.UserID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if isUserAlreadyLikePost {
		err := s.postRepo.DeleteLikePost(ctx, postID, userID)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	} else {
		// else store data
		now := time.Now()
		err := s.postRepo.StoreLikePost(ctx, &model.PostLikeModel{
			UserID:    userID,
			PostID:    postID,
			CreatedAt: now,
			UpdatedAt: now,
		})

		if err != nil {
			return http.StatusInternalServerError, err
		}

	}

	// return
	return http.StatusOK, nil
}
