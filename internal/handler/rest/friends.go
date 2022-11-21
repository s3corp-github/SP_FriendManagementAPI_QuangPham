package rest

import (
	"encoding/json"
	"net/http"
	"net/mail"
	"strings"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/utils"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/service/friends"
)

// FriendsRequest request to create friend or get common friend list
type FriendsRequest struct {
	Friends []string `json:"friends"`
}

// CreateFriendsRequest request to create friend or get common friend list
type CreateFriendsRequest struct {
	Requester string `json:"requester"`
	Target    string `json:"target"`
}

// GetFriendsRequest request to get list friend of users
type GetFriendsRequest struct {
	Email string `json:"email"`
}

// EmailReceiveRequest request to get email receive
type EmailReceiveRequest struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

// CreateUserFriendsResponse response for API create a friend
type CreateUserFriendsResponse struct {
	Success bool `json:"success"`
}

// CreateFriends end point to create friend
func (h Handler) CreateFriends(w http.ResponseWriter, r *http.Request) {
	friendsReq := FriendsRequest{}
	if err := json.NewDecoder(r.Body).Decode(&friendsReq); err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	input, err := validateRelationInput(friendsReq)
	if err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	if err = h.friendService.CreateFriend(r.Context(), input); err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	utils.ResponseJson(w, http.StatusCreated, CreateUserFriendsResponse{
		Success: true,
	})
}

// GetFriends end point to get list friend of users
func (h Handler) GetFriends(w http.ResponseWriter, r *http.Request) {
	getRelationReq := GetFriendsRequest{}
	if err := json.NewDecoder(r.Body).Decode(&getRelationReq); err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	input, err := validateGetRelationInput(getRelationReq)
	if err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	result, err := h.friendService.GetFriends(r.Context(), input)
	if err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, result)
}

// GetCommonFriends end point to get list common friend
func (h Handler) GetCommonFriends(w http.ResponseWriter, r *http.Request) {
	getRelationReq := FriendsRequest{}
	if err := json.NewDecoder(r.Body).Decode(&getRelationReq); err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	input, err := validateRelationCommonInput(getRelationReq)
	if err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	result, err := h.friendService.GetCommonFriends(r.Context(), input)
	if err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, result)
}

// CreateSubscription end point to create friend
func (h Handler) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	friendsReq := CreateFriendsRequest{}
	if err := json.NewDecoder(r.Body).Decode(&friendsReq); err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	input, err := validateSubAndBlockRelationInput(friendsReq)
	if err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	if err = h.friendService.CreateSubscription(r.Context(), input); err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	utils.ResponseJson(w, http.StatusCreated, CreateUserFriendsResponse{
		Success: true,
	})
}

// CreateBlock end point to create friend
func (h Handler) CreateBlock(w http.ResponseWriter, r *http.Request) {
	blockReq := CreateFriendsRequest{}
	if err := json.NewDecoder(r.Body).Decode(&blockReq); err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	input, err := validateSubAndBlockRelationInput(blockReq)
	if err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	if err = h.friendService.CreateBlock(r.Context(), input); err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	utils.ResponseJson(w, http.StatusCreated, CreateUserFriendsResponse{
		Success: true,
	})
}

// GetEmailReceive end point to create friend
func (h Handler) GetEmailReceive(w http.ResponseWriter, r *http.Request) {
	relationReq := EmailReceiveRequest{}
	if err := json.NewDecoder(r.Body).Decode(&relationReq); err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	input, err := validateEmailReceiveInput(relationReq)
	if err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	result, err := h.friendService.GetEmailReceive(r.Context(), input)
	if err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, result)
}

// validateRelationInput function validate create friends request
func validateRelationInput(relationReq FriendsRequest) (friends.CreateRelationsInput, error) {
	if len(relationReq.Friends) != 2 {
		return friends.CreateRelationsInput{}, ErrInvalidBodyRequest
	}

	requesterEmail := strings.TrimSpace(relationReq.Friends[0])
	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return friends.CreateRelationsInput{}, ErrInvalidEmail
	}

	targetEmail := strings.TrimSpace(relationReq.Friends[1])
	if _, err := mail.ParseAddress(targetEmail); err != nil {
		return friends.CreateRelationsInput{}, ErrInvalidEmail
	}

	if strings.EqualFold(requesterEmail, targetEmail) {
		return friends.CreateRelationsInput{}, ErrRequesterAndTargetEmail
	}

	return friends.CreateRelationsInput{
		RequesterEmail: requesterEmail,
		TargetEmail:    targetEmail,
	}, nil
}

// validateSubAndBlockRelationInput function validate create sub and block friends request
func validateSubAndBlockRelationInput(relationReq CreateFriendsRequest) (friends.CreateRelationsInput, error) {
	requesterEmail := strings.TrimSpace(relationReq.Requester)
	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return friends.CreateRelationsInput{}, ErrInvalidEmail
	}

	targetEmail := strings.TrimSpace(relationReq.Target)
	if _, err := mail.ParseAddress(targetEmail); err != nil {
		return friends.CreateRelationsInput{}, ErrInvalidEmail
	}

	if requesterEmail == targetEmail {
		return friends.CreateRelationsInput{}, ErrRequesterAndTargetEmail
	}

	return friends.CreateRelationsInput{
		RequesterEmail: requesterEmail,
		TargetEmail:    targetEmail,
	}, nil
}

// validateGetRelationInput function validate get friend request
func validateGetRelationInput(relationReq GetFriendsRequest) (friends.GetAllFriendsInput, error) {
	requesterEmail := strings.TrimSpace(relationReq.Email)
	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return friends.GetAllFriendsInput{}, ErrInvalidEmail
	}

	return friends.GetAllFriendsInput{
		Email: requesterEmail,
	}, nil
}

// validateRelationCommonInput function validate get common friend request
func validateRelationCommonInput(relationReq FriendsRequest) (friends.CommonFriendsInput, error) {
	if len(relationReq.Friends) != 2 {
		return friends.CommonFriendsInput{}, ErrInvalidBodyRequest
	}

	requesterEmail := strings.TrimSpace(relationReq.Friends[0])
	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return friends.CommonFriendsInput{}, ErrInvalidEmail
	}

	targetEmail := strings.TrimSpace(relationReq.Friends[1])
	if _, err := mail.ParseAddress(targetEmail); err != nil {
		return friends.CommonFriendsInput{}, ErrInvalidEmail
	}

	return friends.CommonFriendsInput{
		RequesterEmail: requesterEmail,
		TargetEmail:    targetEmail,
	}, nil
}

// validateEmailReceiveInput function validate get receive email request
func validateEmailReceiveInput(relationReq EmailReceiveRequest) (friends.EmailReceiveInput, error) {
	requesterEmail := strings.TrimSpace(relationReq.Sender)
	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return friends.EmailReceiveInput{}, ErrInvalidEmail
	}

	return friends.EmailReceiveInput{
		Sender: requesterEmail,
		Text:   relationReq.Text,
	}, nil
}
