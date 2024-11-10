package queries

import (
	"database/sql"
	"kood/social-network/pkg/api"
	"time"
)

type UsersRepository struct {
	DB *sql.DB
}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{DB: DBWrapper}
}

func (repo *UsersRepository) CreateUser(u *api.RegistrationRequest) (*api.UserResponseInfo, error) {
	stmt, err := repo.DB.Prepare(`INSERT INTO users (nickname, date_of_birth, about_me, first_name, last_name, email, password_hash ,created_at, avatar) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Nickname, u.DateOfBirth, u.AboutMe, u.FirstName, u.LastName, u.Email, u.Password, u.CreatedAt, u.Avatar)
	if err != nil {
		return nil, err
	}

	var newUser api.UserResponseInfo
	err = repo.DB.QueryRow(`SELECT id, nickname FROM users WHERE email = ?`, u.Email).Scan(&newUser.ID, &newUser.Nickname)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (repo *UsersRepository) ChekUserByEmail(email string) (*api.LoginRequest, *api.UserResponseInfo, error) {
	var user api.LoginRequest
	var userD api.UserResponseInfo
	err := repo.DB.QueryRow(`SELECT email, password_hash, id, nickname FROM users WHERE email = ?`, email).Scan(
		&user.Email, &user.Password, &userD.ID, &userD.Nickname)
	if err == sql.ErrNoRows {
		return nil, nil, err
	} else if err != nil {
		return nil, nil, err
	}
	return &user, &userD, nil
}

func (repo *UsersRepository) GetUserBySessionID(sessionID string) (*api.UserResponseInfo, error) {
	var user api.UserResponseInfo
	err := repo.DB.QueryRow(`
		SELECT u.id, u.nickname 
		FROM users u
		INNER JOIN active_sessions s ON u.id = s.user_id
		WHERE s.session_id = ?
	`, sessionID).Scan(&user.ID, &user.Nickname)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (repo *UsersRepository) GetUsers(userID int, users *[]api.User) error {
	rows, err := repo.DB.Query(`
		SELECT 
			id, 
			nickname, 
			first_name, 
			last_name, 
			date_of_birth, 
			avatar, 
			about_me, 
			banner_color, 
			profile_visibility, 
			created_at, 
			post_count, 
			comment_count, 
			follower_count, 
			following_count, 
			last_active
		FROM users
		WHERE id != ?
	`, userID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var user api.User
		if err := rows.Scan(
			&user.ID,
			&user.Nickname,
			&user.FirstName,
			&user.LastName,
			&user.DateOfBirth,
			&user.AvatarPath,
			&user.AboutMe,
			&user.BannerColor,
			&user.ProfileVisibility,
			&user.CreatedAt,
			&user.PostCount,
			&user.CommentCount,
			&user.FollowerCount,
			&user.FollowingCount,
			&user.LastActive,
		); err != nil {
			return err
		}
		*users = append(*users, user)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}

func (repo *UsersRepository) GetUserByID(userID, id int) (*api.UserPage, error) {
	var userPage api.UserPage
	err := repo.DB.QueryRow(`
		SELECT 
			u.id, 
			u.nickname, 
			u.first_name, 
			u.last_name, 
			u.date_of_birth, 
			u.avatar, 
			u.about_me, 
			u.banner_color, 
			u.profile_visibility, 
			u.created_at, 
			u.post_count, 
			u.comment_count, 
			u.follower_count, 
			u.following_count, 
			u.last_active
		FROM users u
		WHERE u.id = ?
	`, id).Scan(
		&userPage.User.ID,
		&userPage.User.Nickname,
		&userPage.User.FirstName,
		&userPage.User.LastName,
		&userPage.User.DateOfBirth,
		&userPage.User.AvatarPath,
		&userPage.User.AboutMe,
		&userPage.User.BannerColor,
		&userPage.User.ProfileVisibility,
		&userPage.User.CreatedAt,
		&userPage.User.PostCount,
		&userPage.User.CommentCount,
		&userPage.User.FollowerCount,
		&userPage.User.FollowingCount,
		&userPage.User.LastActive,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &userPage, nil
}


func (repo *UsersRepository) UpdateLastActivity(userID int) error {
	_, err := repo.DB.Exec(`
		UPDATE users
		SET last_activity = ?
		WHERE id = ?
	`, time.Now().UTC(), userID)
	if err != nil {
		return err
	}
	return nil
}