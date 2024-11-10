package queries

import (
	"database/sql"
	"kood/social-network/pkg/api"
	"time"
)

type PostsRepository struct {
	DB *sql.DB
}

func NewPostsRepository() *PostsRepository {
	return &PostsRepository{DB: DBWrapper}
}

func (repo *PostsRepository) Create(UserID int, post *api.PostCreateRequest) (int, error) {
	var stmt *sql.Stmt
	var err error

	if post.Groupinfo.ID > 0 {
		stmt, err = repo.DB.Prepare("INSERT INTO posts (user_id, title, content, created_at, image, privacy, group_id) VALUES (?, ?, ?, ?, ?, ?, ?)")
	} else {
		stmt, err = repo.DB.Prepare("INSERT INTO posts (user_id, title, content, created_at, image, privacy) VALUES (?, ?, ?, ?, ?, ?)")
	}

	if err != nil {
		return 0, err
	}

	post.CreatedAt = time.Now().UTC()

	var result sql.Result
	if post.Groupinfo.ID > 0 {
		result, err = stmt.Exec(UserID, post.Title, post.Content, post.CreatedAt, post.Image, post.Privacy, post.Groupinfo.ID)
	} else {
		result, err = stmt.Exec(UserID, post.Title, post.Content, post.CreatedAt, post.Image, post.Privacy)
	}

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (repo *PostsRepository) GetPostsWithoutGroups(userID int) ([]api.Post, error) {
	query := `
	SELECT p.id, p.user_id, u.nickname, p.title, p.content, p.image, p.privacy, p.created_at, 
           p.comment_count, p.like_count, p.dislike_count,
           COALESCE(r.reaction, 'none') AS user_reaction
	FROM posts p
	LEFT JOIN users u ON p.user_id = u.id
	LEFT JOIN followers f ON p.user_id = f.followee_id AND f.follower_id = ?
	LEFT JOIN (
		SELECT "reaction" AS reaction, "object_id"
		FROM "likes_dislikes"
		WHERE "user_id" = ? AND "object_type" = 'post'
	) r ON p.id = r.object_id
	WHERE (p.privacy = 'public'
	   OR (p.privacy = 'private' AND (p.user_id = ? OR f.status = 'accepted')))
	   AND p.group_id IS NULL
	ORDER BY p.created_at DESC
	`

	rows, err := repo.DB.Query(query, userID, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []api.Post

	for rows.Next() {
		var post api.Post
		if err := rows.Scan(
			&post.ID, &post.Userinfo.ID, &post.Userinfo.Nickname, &post.Title, &post.Content, &post.Image,
			&post.Privacy, &post.CreatedAt, &post.CommentCount,
			&post.ReactionInfo.LikeCount, &post.ReactionInfo.DislikeCount,
			&post.ReactionInfo.Status,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (repo *PostsRepository) GetPostByID(userID, postID int) (*api.Post, error) {
	query := `
   SELECT p.id, p.user_id, u.nickname, p.title, p.content, p.image, p.privacy, p.created_at, 
          p.comment_count, p.like_count, p.dislike_count,
          COALESCE(r.reaction, 'none') AS user_reaction
   FROM posts p
   LEFT JOIN users u ON p.user_id = u.id 
   LEFT JOIN followers f ON p.user_id = f.followee_id AND f.follower_id = ?
   LEFT JOIN (
       SELECT "reaction" AS reaction, "object_id"
       FROM "likes_dislikes"
       WHERE "user_id" = ? AND "object_type" = 'post'
   ) r ON p.id = r.object_id
   WHERE p.id = ?
   AND (
       p.privacy = 'public'
       OR (p.privacy = 'private' AND p.user_id = ?)
       OR (p.privacy = 'private' AND f.status = 'accepted')
   )
   `
	var post api.Post
	err := repo.DB.QueryRow(query, userID, userID, postID, userID).Scan(
		&post.ID, &post.Userinfo.ID, &post.Userinfo.Nickname, &post.Title, &post.Content, &post.Image,
		&post.Privacy, &post.CreatedAt, &post.CommentCount,
		&post.ReactionInfo.LikeCount, &post.ReactionInfo.DislikeCount,
		&post.ReactionInfo.Status,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &post, nil
}

func (repo *PostsRepository) GetPostsForGroup(groupID int) ([]api.Post, error) {
	query := `
	SELECT p.id, p.user_id, u.nickname, p.title, p.content, p.image, p.privacy, p.created_at, 
		   p.comment_count, p.like_count, p.dislike_count,
		   COALESCE(r.reaction, 'none') AS user_reaction
	FROM posts p
	LEFT JOIN users u ON p.user_id = u.id
	LEFT JOIN followers f ON p.user_id = f.followee_id AND f.follower_id = ?
	LEFT JOIN (
		SELECT "reaction" AS reaction, "object_id"
		FROM "likes_dislikes"
		WHERE "user_id" = ? AND "object_type" = 'post'
	) r ON p.id = r.object_id
	WHERE p.group_id = ?
	ORDER BY p.created_at DESC
	`

	rows, err := repo.DB.Query(query, groupID, groupID, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []api.Post

	for rows.Next() {
		var post api.Post
		if err := rows.Scan(
			&post.ID, &post.Userinfo.ID, &post.Userinfo.Nickname, &post.Title, &post.Content, &post.Image,
			&post.Privacy, &post.CreatedAt, &post.CommentCount,
			&post.ReactionInfo.LikeCount, &post.ReactionInfo.DislikeCount,
			&post.ReactionInfo.Status,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (repo *PostsRepository) GetPostsForUser(requesterID, targetUserID int) ([]api.Post, error) {
	query := `
	SELECT p.id, p.user_id, u.nickname, p.title, p.content, p.image, p.privacy, p.created_at, 
           p.comment_count, p.like_count, p.dislike_count,
           COALESCE(r.reaction, 'none') AS user_reaction
	FROM posts p
	LEFT JOIN users u ON p.user_id = u.id
	LEFT JOIN followers f ON p.user_id = f.followee_id AND f.follower_id = ?
	LEFT JOIN (
		SELECT "reaction" AS reaction, "object_id"
		FROM "likes_dislikes"
		WHERE "user_id" = ? AND "object_type" = 'post'
	) r ON p.id = r.object_id
	WHERE p.user_id = ?
	AND (p.privacy = 'public'
	   OR (p.privacy = 'private' AND (p.user_id = ? OR f.status = 'accepted')))
	   AND p.group_id IS NULL
	ORDER BY p.created_at DESC
	`

	rows, err := repo.DB.Query(query, requesterID, requesterID, targetUserID, requesterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []api.Post

	for rows.Next() {
		var post api.Post
		if err := rows.Scan(
			&post.ID, &post.Userinfo.ID, &post.Userinfo.Nickname, &post.Title, &post.Content, &post.Image,
			&post.Privacy, &post.CreatedAt, &post.CommentCount,
			&post.ReactionInfo.LikeCount, &post.ReactionInfo.DislikeCount,
			&post.ReactionInfo.Status,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (repo *PostsRepository) GetGroupPostsForUser(requesterID, targetUserID int) ([]api.Post, error) {
	query := `
	SELECT p.id, p.user_id, u.nickname, p.title, p.content, p.image, p.privacy, p.created_at, 
           p.comment_count, p.like_count, p.dislike_count,
           COALESCE(r.reaction, 'none') AS user_reaction,
           g.id, g.title
	FROM posts p
	LEFT JOIN users u ON p.user_id = u.id
	LEFT JOIN followers f ON p.user_id = f.followee_id AND f.follower_id = ?
	LEFT JOIN (
		SELECT "reaction" AS reaction, "object_id"
		FROM "likes_dislikes"
		WHERE "user_id" = ? AND "object_type" = 'post'
	) r ON p.id = r.object_id
	LEFT JOIN group_memberships gm ON gm.group_id = p.group_id AND gm.user_id = ?
	LEFT JOIN groups g ON p.group_id = g.id
	WHERE p.user_id = ?
	AND p.group_id IS NOT NULL
	AND gm.user_id IS NOT NULL
	ORDER BY p.created_at DESC
	`

	rows, err := repo.DB.Query(query, requesterID, requesterID, requesterID, targetUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []api.Post

	for rows.Next() {
		var post api.Post
		var groupID sql.NullInt64
		var groupName sql.NullString

		if err := rows.Scan(
			&post.ID, &post.Userinfo.ID, &post.Userinfo.Nickname, &post.Title, &post.Content, &post.Image,
			&post.Privacy, &post.CreatedAt, &post.CommentCount,
			&post.ReactionInfo.LikeCount, &post.ReactionInfo.DislikeCount,
			&post.ReactionInfo.Status,
			&groupID, &groupName,
		); err != nil {
			return nil, err
		}

		if groupID.Valid && groupName.Valid {
			post.Groupinfo = api.GroupResponseInfo{
				ID:   int(groupID.Int64),
				Name: groupName.String,
			}
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
