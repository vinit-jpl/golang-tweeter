package comment

import (
	"context"
	"errors"
	"go-tweets/internal/model"
	"net/http"
	"time"
)

func (s *commentService) LikeOrUnlikeComment(ctx context.Context, commentID, userID int64) (int, error) {

	// check if comment exists
	commentExists, err := s.CommentRepo.GetComment(ctx, commentID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if commentExists == nil {
		return http.StatusNotFound, errors.New("comment not found")
	}

	// check if user already liked the comment
	isUserAlreadyLikedComment, err := s.CommentRepo.IsUserAlreadyLikeComment(ctx, commentID, userID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// if user liked, delete the data
	if isUserAlreadyLikedComment {
		err := s.CommentRepo.DeleteLikeComment(ctx, commentID, userID)

		if err != nil {
			return http.StatusInternalServerError, err
		}
	} else {
		// else, store data
		now := time.Now()
		err := s.CommentRepo.StoreLikeComment(ctx, &model.CommentLikeModel{
			UserID:    userID,
			CommentID: commentID,
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
