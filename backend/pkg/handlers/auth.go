package handlers

import (
	"database/sql"
	"encoding/json"
	"kood/social-network/pkg/api"
	"kood/social-network/pkg/db/queries"
	"kood/social-network/pkg/services"
	"net/http"
	"time"
)

// HandleRegister registers a new user.
// @Summary Register a new user
// @Description Registers a new user with the provided registration details and optional photo.
// @Tags auth
// @Accept json
// @Produce json
// @Param body body api.RegistrationRequest true "Registration details"
// @Response 200 {object} api.Response{payload=api.UserResponseInfo}
// @Response 400 {object} api.Response{eror=api.ErrorDetails}
// @Response 500 {object} api.Response{eror=api.ErrorDetails}
// @Router /auth/register [post]
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	authenticated := false
	var registrForm api.RegistrationRequest

	if err := json.NewDecoder(r.Body).Decode(&registrForm); err != nil {
		services.HTTPError(w, http.StatusBadRequest, "Invalid request payload", err.Error(), authenticated, nil, nil)
		return
	}
	registrForm.Nickname = services.LoverCase(registrForm.Nickname)
	registrForm.Email = services.LoverCase(registrForm.Email)

	registrForm.AboutMe = services.TrimAndNormalizeSpaces(registrForm.AboutMe)

	validationErrors := services.ValidateOperation("registration", registrForm)
	if len(validationErrors) > 0 {
		services.HTTPError(w, http.StatusBadRequest, "Validation error", "Validation error", authenticated, nil, validationErrors)
		return
	}

	hashedPassword, err := services.HashPassword(registrForm.Password)
	if err != nil {
		services.HTTPError(w, http.StatusInternalServerError, "Error hashing password", "Error hashing password", authenticated, nil, nil)
		return
	}
	registrForm.Password = string(hashedPassword)
	registrForm.CreatedAt = time.Now().UTC()

	userRepo := queries.NewUsersRepository()
	_, existingUser, err := userRepo.ChekUserByEmail(registrForm.Email)
	if err != nil && err != sql.ErrNoRows {
		services.HTTPError(w, http.StatusInternalServerError, "Error checking user by email", "Error checking user by email", authenticated, nil, nil)
		return
	}
	if existingUser != nil {
		services.HTTPError(w, http.StatusConflict, "User already registered", "User already registered", authenticated, nil, nil)
		return
	}

	userResponse, err := userRepo.CreateUser(&registrForm)
	if err != nil {
		services.HTTPError(w, http.StatusInternalServerError, "Error creating user", err.Error(), authenticated, nil, nil)
		return
	}

	sessionID, expiresAt, err := services.CreateSession(userResponse.ID)
	if err != nil {
		services.HTTPError(w, http.StatusInternalServerError, "Error creating session", err.Error(), authenticated, nil, nil)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionID,
		Expires:  expiresAt,
		HttpOnly: false,
		Path:     "/",
	})

	authenticated = true

	services.RespondWithSuccess(w, http.StatusOK, "User registered successfully", authenticated, nil, nil, userResponse)
}

// HandleLogin logs in a user.
// @Summary Log in a user
// @Description Logs in a user with the provided email and password.
// @Tags auth
// @Accept json
// @Produce json
// @Param body body api.LoginRequest true "Login credentials"
// @Response 200 {object} api.Response{payload=api.UserResponseInfo}
// @Response 400 {object} api.Response{eror=api.ErrorDetails}
// @Response 500 {object} api.Response{eror=api.ErrorDetails}
// @Router /auth/login [post]
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	authenticated := false
	var loginForm api.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&loginForm); err != nil {
		services.HTTPError(w, http.StatusBadRequest, "Invalid request payload", "Invalid request payload", authenticated, nil, nil)
		return
	}
	loginForm.Email = services.LoverCase(loginForm.Email)
	userRepo := queries.NewUsersRepository()
	storedUser, userResponse, err := userRepo.ChekUserByEmail(loginForm.Email)
	if err != nil {
		services.HTTPError(w, http.StatusUnauthorized, "Invalid email or password", "User not found", authenticated, nil, nil)
		return
	}

	if err := services.CheckPassword(storedUser.Password, loginForm.Password); err != nil {
		services.HTTPError(w, http.StatusUnauthorized, "Invalid email or password", "Invalid password", authenticated, nil, nil)
		return
	}

	sessionID, expiresAt, err := services.CreateSession(userResponse.ID)
	if err != nil {
		services.HTTPError(w, http.StatusInternalServerError, "Error creating session", err.Error(), authenticated, nil, nil)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionID,
		Expires:  expiresAt,
		HttpOnly: true,
		Path:     "/",
	})

	authenticated = true
	services.RespondWithSuccess(w, http.StatusOK, "User logged in successfully", authenticated, nil, nil, userResponse)
}

// HandleLogout logs out a user
// @Summary Log out a user
// @Description Logs out the currently authenticated user.
// @Tags auth
// @Security BearerAuth
// @Accept json
// @Produce json
// @Response 200 {object} api.Response
// @Response 400 {object} api.Response{eror=api.ErrorDetails}
// @Response 500 {object} api.Response{eror=api.ErrorDetails}
// @Router /auth/logout [delete]
func HandleLogout(w http.ResponseWriter, r *http.Request) {
	authenticated := true
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			services.HTTPError(w, http.StatusUnauthorized, "Missing session ID", "Missing session ID", authenticated, nil, nil)
			return
		}
		services.HTTPError(w, http.StatusBadRequest, "Error reading cookie", "Error reading cookie", authenticated, nil, nil)
		return
	}
	sessionID := cookie.Value
	sessionRepo := queries.NewSesionsRepository()
	if err := sessionRepo.DeleteSession(sessionID); err != nil {
		services.HTTPError(w, http.StatusInternalServerError, "Error deleting session", err.Error(), authenticated, nil, nil)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().UTC().Add(-1 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})

	authenticated = false
	services.RespondWithSuccess(w, http.StatusOK, "User logged out successfully", authenticated, nil, nil, nil)
}
