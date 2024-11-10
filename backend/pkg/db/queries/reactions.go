package queries

import (
	"database/sql"
	"time"
)

type ReactionRepository struct {
	DB *sql.DB
}

func NewReactionsRepository() *ReactionRepository {
	return &ReactionRepository{DB: DBWrapper}
}

func (repo *ReactionRepository) CreateOrUpdateReaction(userID int, objectType string, objectID int, reactionType string) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var existingReaction string
	checkQuery := `
    SELECT "reaction" FROM "likes_dislikes"
    WHERE "user_id" = ? AND "object_type" = ? AND "object_id" = ?;
    `
	err = tx.QueryRow(checkQuery, userID, objectType, objectID).Scan(&existingReaction)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if existingReaction == reactionType {
		deleteQuery := `
        DELETE FROM "likes_dislikes"
        WHERE "user_id" = ? AND "object_type" = ? AND "object_id" = ?;
        `
		_, err = tx.Exec(deleteQuery, userID, objectType, objectID)
		if err != nil {
			return err
		}
		return nil 
	} else {
		upsertQuery := `
        INSERT INTO "likes_dislikes" ("user_id", "object_type", "object_id", "reaction", "created_at")
        VALUES (?, ?, ?, ?, ?)
        ON CONFLICT ("user_id", "object_type", "object_id") 
        DO UPDATE SET "reaction" = excluded."reaction", "created_at" = excluded."created_at";
        `
		_, err = tx.Exec(upsertQuery, userID, objectType, objectID, reactionType, time.Now().UTC())
		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *ReactionRepository) GetReactionCounts(objectType string, objectID int) (int, int, error) {
	var likes, dislikes int

	queryLikes := `
    SELECT COUNT(*) FROM "likes_dislikes"
    WHERE "object_type" = ? AND "object_id" = ? AND "reaction" = 'like';
    `
	err := repo.DB.QueryRow(queryLikes, objectType, objectID).Scan(&likes)
	if err != nil {
		return 0, 0, err
	}

	queryDislikes := `
    SELECT COUNT(*) FROM "likes_dislikes"
    WHERE "object_type" = ? AND "object_id" = ? AND "reaction" = 'dislike';
    `
	err = repo.DB.QueryRow(queryDislikes, objectType, objectID).Scan(&dislikes)
	if err != nil {
		return 0, 0, err
	}

	var updateQuery string
	var params []interface{}

	switch objectType {
	case "post":
		updateQuery = `
        UPDATE "posts"
        SET "like_count" = ?, "dislike_count" = ?
        WHERE "id" = ?;
        `
		params = append(params, likes, dislikes, objectID)
	case "comment":
		updateQuery = `
        UPDATE "post_comments"
        SET "like_count" = ?, "dislike_count" = ?
        WHERE "id" = ?;
        `
		params = append(params, likes, dislikes, objectID)
	default:
		return 0, 0, nil
	}

	_, err = repo.DB.Exec(updateQuery, params...)
	if err != nil {
		return 0, 0, err
	}

	return likes, dislikes, nil
}
