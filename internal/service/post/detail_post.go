package post

import (
	"context"
	"errors"
	"go-tweets/internal/dto"
	"net/http"
)

func (s *postService) DetailPost(ctx context.Context, postID int64) (*dto.DetailPostResponse, int, error) {

	// get post by id
	post, err := s.postRepo.GetPostByID(ctx, postID)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if post == nil {
		return nil, http.StatusNotFound, errors.New("tweet not found")
	}

	// get all the comments related to the post
	postIDs := []int64{post.ID}
	comments, err := s.commentRepo.GetCommentsByPostID(ctx, postIDs)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// mapping comment with post

	commentsMap := make([]dto.Comment, 0)
	for _, comment := range comments {
		commentsMap = append(commentsMap, dto.Comment{
			ID:        comment.ID,
			Username:  comment.Username,
			Content:   comment.Content,
			LikeCount: comment.LikeCount,
			CreatedAt: comment.CreatedAt.String(),
			UpdatedAt: comment.UpdatedAt.String(),
		})
	}

	// set response
	return &dto.DetailPostResponse{
		ID:        post.ID,
		Username:  post.Username,
		Title:     post.Title,
		Content:   post.Content,
		LikeCount: post.LikeCount,
		Comments:  commentsMap,
		CreatedAt: post.CreatedAt.String(),
		UpdatedAt: post.UpdatedAt.String(),
	}, http.StatusOK, nil
}
