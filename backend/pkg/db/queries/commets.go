package queries

import (
	"database/sql"
	"kood/social-network/pkg/api"
	"time"
)

type CommentsRepository struct {
	DB *sql.DB
}

func NewCommentsRepository() *CommentsRepository {
	return &CommentsRepository{DB: DBWrapper}
}

func (r *CommentsRepository) Create(userID int, content *api.CommentCreateRequest) (int, error) {
	query := `INSERT INTO post_comments (post_id, user_id, content, created_at) VALUES (?,?,?,?) RETURNING id`
	var commentID int
	content.CreatedAt = time.Now().UTC()
	err := r.DB.QueryRow(query, content.PostID, userID, content.Content, content.CreatedAt).Scan(&commentID)
	if err != nil {
		return 0, err
	}
	return commentID, nil
}

func (r *CommentsRepository) GetCommentsForPost(userID, postID int) (*[]api.Comment, error) {
	query := `
	SELECT c.id, c.post_id, c.user_id, u.nickname, c.content, c.created_at, 
           c.like_count, c.dislike_count,
           COALESCE(r.reaction, 'none') AS user_reaction
	FROM post_comments c
	LEFT JOIN users u ON c.user_id = u.id
	LEFT JOIN (
		SELECT "reaction", "object_id"
		FROM "likes_dislikes"
		WHERE "user_id" = ? AND "object_type" = 'comment'
	) r ON c.id = r.object_id
	WHERE c.post_id = ?
	ORDER BY c.created_at DESC
	`

	rows, err := r.DB.Query(query, userID, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []api.Comment

	for rows.Next() {
		var comment api.Comment
		if err := rows.Scan(
			&comment.ID, &comment.PostID, &comment.UserResponse.ID, &comment.UserResponse.Nickname,
			&comment.Content, &comment.CreatedAt, &comment.ReactionInfo.LikeCount, &comment.ReactionInfo.DislikeCount, &comment.ReactionInfo.Status,
		); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &comments, nil
}
