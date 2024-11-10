package queries

import (
	"database/sql"
	"kood/social-network/pkg/models"
)

type SesionsRepository struct {
	DB *sql.DB
}

func NewSesionsRepository() *SesionsRepository {
	return &SesionsRepository{DB: DBWrapper}
}

func (repo *SesionsRepository) CreateSession(s *models.Session) error {
	stmt, err := repo.DB.Prepare("INSERT INTO active_sessions (user_id, session_id, created_at, expires_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(s.UserID, s.SessionID, s.CreatedAt, s.ExpiresAt)
	if err != nil {
		return err
	}

	return nil
}

func (repo *SesionsRepository) DeleteSession(sessionID string) error {
	stmt, err := repo.DB.Prepare("DELETE FROM active_sessions WHERE session_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(sessionID)
	if err != nil {
		return err
	}
	return nil
}

// // UpdateLastActivity updates the last activity timestamp for a session
// func UpdateLastActivity(sessionID string) error {
// 	stmt, err := dbHandler.MainDB.Prepare("UPDATE active_sessions SET last_activity = ? WHERE session_id = ?")
// 	if err != nil {
// 		return err
// 	}
// 	_, err = stmt.Exec(time.Now().UTC(), sessionID)
// 	return err
// }
