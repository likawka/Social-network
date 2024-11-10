package queries

import (
	"database/sql"
	"fmt"
	"kood/social-network/pkg/api"
	"time"
)

type FollowsRepository struct {
	DB *sql.DB
}

func NewFollowsRepository() *FollowsRepository {
	return &FollowsRepository{DB: DBWrapper}
}

func (repo *FollowsRepository) CreateFollowRequest(followerID, followeeID int) error {
	_, err := repo.DB.Exec(
		"INSERT INTO followers (follower_id, followee_id, created_at, status) VALUES (?, ?, ?, 'pending')",
		followerID, followeeID, time.Now().UTC())
	if err != nil {
		return err
	}
	return nil
}

func (repo *FollowsRepository) ProcessNotificationResponse(userID int, form *api.NotificationResponse, table string) error {
	var updateQuery string
	var deleteQuery string

	switch table {
	case "followers":
		updateQuery = "UPDATE followers SET status = ? WHERE follower_id = ? AND id = ?"
	case "group_invitations":
		updateQuery = "UPDATE group_invitations SET status = ? WHERE user_id = ? AND id = ?"
	case "events":
		updateQuery = "INSERT INTO event_response (event_id, user_id, status) VALUES (?, ?, ?)"
	default:
		return fmt.Errorf("invalid table: %s", table)
	}

	result, err := repo.DB.Exec(updateQuery, form.Status, userID, form.ReqestID)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 0 {
		return fmt.Errorf("no rows affected in table %s with ReqestID %d", table, form.ReqestID)
	}

	if form.NotificationInfo.IDRef != 0 {
		deleteQuery = "DELETE FROM notifications WHERE user_id = ? AND type = ? AND id_ref = ?"
		_, err = repo.DB.Exec(deleteQuery, userID, form.NotificationInfo.Type, form.NotificationInfo.IDRef)
	} else {
		deleteQuery = "DELETE FROM notifications WHERE id = ?"
		_, err = repo.DB.Exec(deleteQuery, form.NotificationInfo.ID)
	}

	if err != nil {
		return err
	}

	return nil
}

func (repo *FollowsRepository) RemoveFollow(followerID, followeeID int) error {
	_, err := repo.DB.Exec(
		"DELETE FROM followers WHERE follower_id = ? AND followee_id = ?",
		followerID, followeeID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *FollowsRepository) CheckFollow(followerID, followeeID int) bool {
	var count int
	err := repo.DB.QueryRow(
		"SELECT COUNT(*) FROM followers WHERE follower_id = ? AND followee_id = ?",
		followerID, followeeID).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (repo *FollowsRepository) GetFollows(userID int, followType string) ([]api.UserResponseInfo, error) {
	var query string

	switch followType {
	case "followers":
		query = `
		SELECT u.id, u.nickname
		FROM users u
		JOIN followers f ON u.id = f.follower_id
		WHERE f.followee_id = ? AND f.status = 'accepted'
		`
	case "following":
		query = `
		SELECT u.id, u.nickname
		FROM users u
		JOIN followers f ON u.id = f.followee_id
		WHERE f.follower_id = ? AND f.status = 'accepted'
		`
	default:
		return nil, fmt.Errorf("invalid follow type: %s", followType)
	}

	rows, err := repo.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []api.UserResponseInfo

	for rows.Next() {
		var user api.UserResponseInfo
		if err := rows.Scan(&user.ID, &user.Nickname); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}