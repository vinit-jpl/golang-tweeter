package post

import (
	"context"
	"database/sql"
	"fmt"
	"go-tweets/internal/model"
	"log"
)

func (r *postRepository) GetPostByID(ctx context.Context, postID int64) (*model.PostWithUserModel, error) {

	query := `
		SELECT p.id, p.title, p.content, p.user_id, p.created_at, p.updated_at, u.username, COUNT(pl.id) AS like_count 
		FROM posts AS p
		JOIN users as u ON p.user_id = u.id
		LEFT JOIN post_likes AS pl ON pl.post_id = p.id
		WHERE p.id = ?
		AND deleted_at IS NULL
		GROUP BY p.id, p.title, p.content, p.user_id, p.created_at, p.updated_at, u.username
	`

	log.Println("in get post by id")
	row := r.db.QueryRowContext(ctx, query, postID)

	var result model.PostWithUserModel

	err := row.Scan(
		&result.ID,
		&result.Title,
		&result.Content,
		&result.UserID,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.Username,
		&result.LikeCount,
	)

	fmt.Println("result ", result)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil

}
