package services

import (
	"kood/social-network/pkg/api"
	"kood/social-network/pkg/db/queries"
	"kood/social-network/pkg/models"
	"log"
	"net/http"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"golang.org/x/crypto/bcrypt"
)

func CreateSession(userID int) (string, time.Time, error) {
	sessionRepo := queries.NewSesionsRepository()
	sessionID := uuid.New().String()
	expiresAt := time.Now().UTC().Add(24 * time.Hour)
	session := models.Session{
		UserID:    userID,
		SessionID: sessionID,
		CreatedAt: time.Now().UTC(),
		ExpiresAt: expiresAt,
	}
	if err := sessionRepo.CreateSession(&session); err != nil {
		return "", time.Time{}, err
	}
	return sessionID, expiresAt, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println("Password comparison failed:", err)
		return err
	}
	return nil
}

func GetRouteParams(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func DecodeForm(r *http.Request, dest interface{}) error {
	decoder := schema.NewDecoder()
	return decoder.Decode(dest, r.Form)
}

func getSessionID(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func AuthenticateUser(r *http.Request) (*api.UserResponseInfo, bool) {
	authenticated := false
	var user *api.UserResponseInfo
	sessionID, err := getSessionID(r)
	if err != nil {
		return &api.UserResponseInfo{ID: 0, Nickname: ""}, false
	}

	userRepo := queries.NewUsersRepository()

	u, err := userRepo.GetUserBySessionID(sessionID)
	if err != nil || u == nil {
		return &api.UserResponseInfo{ID: 0, Nickname: ""}, false
	}

	user = &api.UserResponseInfo{
		ID:       u.ID,
		Nickname: u.Nickname,
	}
	authenticated = true

	return user, authenticated
}

func TrimAndNormalizeSpaces(input string) string {
	trimmed := strings.TrimSpace(input)
	words := strings.FieldsFunc(trimmed, unicode.IsSpace)
	result := strings.Join(words, " ")
	return result
}


func LoverCase(input string) string {
	return strings.ToLower(input)
}