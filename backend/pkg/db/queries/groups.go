package queries

import (
	"database/sql"
	"fmt"
	"kood/social-network/pkg/api"
	"time"
)

type GroupsRepository struct {
	DB *sql.DB
}

func NewGroupsRepository() *GroupsRepository {
	return &GroupsRepository{DB: DBWrapper}
}

func (repo *GroupsRepository) Create(CreatorID int, group *api.GroupCreateRequest) (int, error) {
	res, err := repo.DB.Exec(
		"INSERT INTO groups (creator_id, title, description, banner_color, created_at) VALUES (?, ?, ?, ?, ?)",
		CreatorID, group.Title, group.Description, group.BannerColor, time.Now().UTC())
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (repo *GroupsRepository) GetGroups(userID int) ([]api.Group, error) {
	rows, err := repo.DB.Query(`
        SELECT 
            g.id, 
            g.title, 
            g.description, 
            g.banner_color, 
            g.created_at,
            g.member_count,
            COALESCE(gm.role, '') AS role
        FROM groups g
        LEFT JOIN group_memberships gm 
            ON g.id = gm.group_id AND gm.user_id = ?
		ORDER BY g.member_count DESC
    `, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []api.Group

	for rows.Next() {
		var group api.Group
		var role string
		err := rows.Scan(&group.ID, &group.Title, &group.Description, &group.BannerColor, &group.CreatedAt, &group.MemberInfo.TotalMembers, &role)
		if err != nil {
			return nil, err
		}

		group.MemberInfo.Roles = role

		groups = append(groups, group)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

func (repo *GroupsRepository) IsUserMemberOfGroup(userID, groupID int) (bool, error) {
	var count int
	query := `
        SELECT COUNT(*) 
        FROM group_memberships 
        WHERE group_id = ? AND user_id = ?`
	err := repo.DB.QueryRow(query, groupID, userID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (repo *GroupsRepository) GroupRequest(UserID int, request *api.GroupRequest) error {
	isMember, err := repo.IsUserMemberOfGroup(UserID, request.GroupID)
	if err != nil {
		return err
	}

	if !isMember && request.RequestType != "j_req" {
		return fmt.Errorf("user is not a member of group")
	}

	_, err = repo.DB.Exec(
		`INSERT INTO group_invitations (group_id, user_id, request_type, status, created_at) 
         VALUES (?, ?, ?, 'pending', CURRENT_TIMESTAMP)`,
		request.GroupID, UserID, request.RequestType,
	)
	if err != nil {
		return err
	}

	return nil
}

func (repo *GroupsRepository) GroupResponse(response *api.NotificationResponse) error {
	_, err := repo.DB.Exec(
		`UPDATE group_invitations SET status = ? WHERE id = ?`,
		response.Status, response.ReqestID,
	)
	if err != nil {
		return err
	}
	_, err = repo.DB.Exec(
		"DELETE FROM notifications WHERE id = ?",
		response.NotificationInfo.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *GroupsRepository) CreateGroupEvent(CreatorID int, event *api.GroupEventCreate) (int, error) {
	isMember, err := repo.IsUserMemberOfGroup(CreatorID, event.GroupID)
	if err != nil {
		return 0, err
	}
	if !isMember {
		return 0, fmt.Errorf("user is not a member of group")
	}

	var eventID int
	err = repo.DB.QueryRow(
		`INSERT INTO group_events (group_id, creator_id, title, description, event_time, created_at) 
		 VALUES (?, ?, ?, ?, ?, ?) RETURNING id`,
		event.GroupID, CreatorID, event.Title, event.Description, event.Date, time.Now().UTC(),
	).Scan(&eventID)
	if err != nil {
		return 0, err
	}

	return eventID, nil
}

func (repo *GroupsRepository) GetGroup(groupID int) (*api.Group, error) {
	var group api.Group

	err := repo.DB.QueryRow(`
		SELECT 
			g.title, 
			g.description, 
			g.banner_color, 
			g.created_at,
			g.member_count,
			u.nickname AS creator_nickname
		FROM groups g
		JOIN users u ON g.creator_id = u.id
		WHERE g.id = ?
	`, groupID).Scan(&group.Title, &group.Description, &group.BannerColor, &group.CreatedAt, &group.MemberInfo.TotalMembers, &group.CreatorInfo.Nickname)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &group, nil
}

func (repo *GroupsRepository) GetGroupMembers(group *api.Group) error {
	rows, err := repo.DB.Query(`
		SELECT 
			u.id, 
			u.nickname, 
			gm.role
		FROM group_memberships gm
		JOIN users u ON gm.user_id = u.id
		WHERE gm.group_id = ?
	`, group.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var member api.GroupMember
		err := rows.Scan(&member.ID, &member.Nickname, &member.Role)
		if err != nil {
			return err
		}

		group.Members = append(group.Members, member)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}

func (repo *GroupsRepository) GetGroupEvents(group *api.Group) error {
	rows, err := repo.DB.Query(`
		SELECT 
			ge.id, 
			ge.title, 
			ge.description, 
			ge.event_time, 
			ge.created_at,
			u.nickname AS creator_nickname
		FROM group_events ge
		JOIN users u ON ge.creator_id = u.id
		WHERE ge.group_id = ?
		ORDER BY ge.event_time DESC
	`, group.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var event api.GroupEvent
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.CreatedAt, &event.CreatorInfo.Nickname)
		if err != nil {
			return err
		}

		group.Events = append(group.Events, event)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}