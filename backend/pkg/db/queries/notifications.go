package queries

import (
	"database/sql"
)

type NotificationsRepository struct {
	DB *sql.DB
}

func NewNotificationsRepository() *NotificationsRepository {
	return &NotificationsRepository{DB: DBWrapper}
}

func (repo *NotificationsRepository) GetUnreadNotificationsCount(userID int) (int, error) {
	var count int

	query := `
		SELECT COUNT(*)
		FROM notifications
		WHERE user_id = ? AND is_read = 0
	`

	err := repo.DB.QueryRow(query, userID).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	return count, nil
}
