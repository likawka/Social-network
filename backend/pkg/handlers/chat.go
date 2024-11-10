package handlers

import (
	"kood/social-network/pkg/api"
	"kood/social-network/pkg/db/queries"
	"kood/social-network/pkg/services"
	"net/http"
	"encoding/json"
)

// HandleGetChats returns a list of chats.
// @Summary Get chats
// @Tags chats
// @Accept json
// @Produce json
// @Response 200 {object} api.Response{payload=api.ChatList}
// @Response 400 {object} api.Response{eror=api.ErrorDetails}
// @Response 500 {object} api.Response{eror=api.ErrorDetails}
// @Router /chats [get]
func HandleGetChats(w http.ResponseWriter, r *http.Request) {
	var user *api.UserResponseInfo
	user, authenticated := services.AuthenticateUser(r)
	if !authenticated {
		services.HTTPError(w, http.StatusUnauthorized, "Unauthorized", "User is not authenticated", false, user, nil)
		return
	}

	var chats api.ChatList
	chatRepo := queries.NewChatsRepository()
	chats.IndividualChats, _ = chatRepo.IndividualChats(user.ID)
	chats.GroupChats, _ = chatRepo.GroupChats(user.ID)
	payload := chats
	services.RespondWithSuccess(w, http.StatusOK, "Chats fetched successfully", authenticated, payload, nil, user)

}

// HandleCreateChat creates a new chat.
// @Summary Create chat
// @Description Creates a new chat with the provided data.
// @Tags chats
// @Accept json
// @Produce json
// @Param chat body api.CrateChat true "Chat Data"
// @Security BearerAuth
// @Response 200 {object} api.Response{payload=api.Chat}
// @Response 400 {object} api.Response{eror=api.ErrorDetails}
// @Response 500 {object} api.Response{eror=api.ErrorDetails}
// @Router /chats [post]
func HandleCreateChat(w http.ResponseWriter, r *http.Request) {
	authenticated := false
	var user *api.UserResponseInfo

	user, authenticated = services.AuthenticateUser(r)
	if !authenticated {
		services.HTTPError(w, http.StatusUnauthorized, "Unauthorized", "User is not authenticated", false, user, nil)
		return
	}

	var chatForm api.CrateChat
	if err := json.NewDecoder(r.Body).Decode(&chatForm); err != nil {
		services.HTTPError(w, http.StatusBadRequest, "Invalid request payload", err.Error(), authenticated, nil, nil)
		return
	}

	chatRepo := queries.NewChatsRepository()
	chatHash, err := chatRepo.Create(user.ID, chatForm.Member.ID)
	if err != nil {
		services.HTTPError(w, http.StatusInternalServerError, "Internal Server Error", err.Error(), authenticated, user, nil)
		return
	}

	services.RespondWithSuccess(w, http.StatusCreated, "Chat created successfully", authenticated, api.CreateChatResponse{chatHash}, nil, user)
}
