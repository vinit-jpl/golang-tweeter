package comment

import (
	"context"
	"errors"
	"go-tweets/internal/dto"
	"go-tweets/internal/model"
	"net/http"
	"time"
)

func (s *commentService) CreateComment(ctx context.Context, req *dto.StoreCommentRequest, userID int64) (int, error) {

	// check if tweet  data is available
	if postExists, err := s.PostRepo.GetPostByID(ctx, req.PostId); err != nil {
		return http.StatusInternalServerError, err

	} else if postExists == nil {
		return http.StatusNotFound, errors.New("post not found")
	}

	// store comment
	now := time.Now()
	err := s.CommentRepo.StoreComment(ctx, &model.CommentModel{
		PostID:    req.PostId,
		UserID:    userID,
		Content:   req.Content,
		CreatedAt: now,
		UpdatedAt: now,
	})

	if err != nil {
		return http.StatusInternalServerError, err
	}

	// return
	return http.StatusCreated, nil
}
