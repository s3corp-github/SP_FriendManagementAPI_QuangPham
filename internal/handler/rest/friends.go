package rest

import (
	"encoding/json"
	"log"
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

// RelationFriendsRequest request to create friend or get common friend list
type RelationFriendsRequest struct {
	Requester string `json:"requester"`
	Target    string `json:"target"`
}

// GetRelationRequest request to get list friend of users
type GetRelationRequest struct {
	Email string `json:"email"`
}

// EmailReceiveRequest request to get email receive
type EmailReceiveRequest struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

// CreateRelationsResponse response for API create a friend
type CreateRelationsResponse struct {
	Success bool `json:"success"`
}

// CreateFriends end point to create friend
func (h Handler) CreateFriends(w http.ResponseWriter, r *http.Request) {
	log.Println("----Start API CreateFriends----")

	friendsResp := CreateRelationsResponse{}
	friendsReq := FriendsRequest{}
	if err := json.NewDecoder(r.Body).Decode(&friendsReq); err != nil {
		log.Println("CreateFriends: error when decode request ", err)
		utils.JsonResponseError(w, err)
		return
	}

	input, err := validateRelationInput(friendsReq)
	if err != nil {
		log.Println("CreateFriends: error validate input ", err)
		utils.JsonResponseError(w, err)
		return
	}

	result, err := h.friendService.CreateFriend(r.Context(), input)
	if err != nil {
		log.Println("CreateRelation error", err)
		utils.JsonResponseError(w, err)
		return
	}

	if result {
		friendsResp = CreateRelationsResponse{
			Success: true,
		}
	} else {
		friendsResp = CreateRelationsResponse{
			Success: false,
		}
	}

	utils.ResponseJson(w, http.StatusCreated, friendsResp)
	log.Println("----End API CreateFriends----")
}

// GetAllFriendOfUser end point to get list friend of users
func (h Handler) GetAllFriendOfUser(w http.ResponseWriter, r *http.Request) {
	log.Println("----Start API GetAllFriendOfUser----")

	getRelationReq := GetRelationRequest{}
	if err := json.NewDecoder(r.Body).Decode(&getRelationReq); err != nil {
		log.Println("GetAllFriendOfUser: error when decode request ", err)
		utils.JsonResponseError(w, err)
		return
	}

	input, err := validateGetRelationInput(getRelationReq)
	if err != nil {
		log.Println("GetAllFriendOfUser: error validate input ", err)
		utils.JsonResponseError(w, err)
		return
	}

	result, err := h.friendService.GetAllFriends(r.Context(), input)
	if err != nil {
		log.Println("GetAllFriendOfUser error", err)
		utils.JsonResponseError(w, err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, result)
	log.Println("----End API GetAllFriendOfUser----")
}

// GetCommonFriend end point to get list common friend
func (h Handler) GetCommonFriend(w http.ResponseWriter, r *http.Request) {
	log.Println("----Start API GetCommonFriend----")

	getRelationReq := FriendsRequest{}
	if err := json.NewDecoder(r.Body).Decode(&getRelationReq); err != nil {
		log.Println("GetCommonFriend: error when decode request ", err)
		utils.JsonResponseError(w, err)
		return
	}

	input, err := validateRelationCommonInput(getRelationReq)
	if err != nil {
		log.Println("GetCommonFriend: error validate input ", err)
		utils.JsonResponseError(w, err)
		return
	}

	result, err := h.friendService.GetCommonFriends(r.Context(), input)
	if err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, result)
	log.Println("----End API GetCommonFriend----")
}

// CreateSubscription end point to create friend
func (h Handler) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	log.Println("----Start API CreateSubscription----")

	relationReq := RelationFriendsRequest{}
	if err := json.NewDecoder(r.Body).Decode(&relationReq); err != nil {
		log.Println("CreateSubscription: error when decode request ", err)
		utils.JsonResponseError(w, err)
		return
	}

	input, err := validateSubAndBlockRelationInput(relationReq)
	if err != nil {
		log.Println("CreateSubscription: error validate input ", err)
		utils.JsonResponseError(w, err)
		return
	}

	result, err := h.friendService.CreateSubscription(r.Context(), input)
	if err != nil {
		log.Println("CreateRelation error", err)
		utils.JsonResponseError(w, err)
		return
	}

	utils.ResponseJson(w, http.StatusCreated, result)
	log.Println("----End API CreateSubscription----")
}

// CreateBlock end point to create friend
func (h Handler) CreateBlock(w http.ResponseWriter, r *http.Request) {
	log.Println("----Start API CreateBlock----")

	relationReq := RelationFriendsRequest{}
	if err := json.NewDecoder(r.Body).Decode(&relationReq); err != nil {
		log.Println("CreateBlock: error when decode request ", err)
		utils.JsonResponseError(w, err)
		return
	}

	input, err := validateSubAndBlockRelationInput(relationReq)
	if err != nil {
		log.Println("CreateBlock: error validate input ", err)
		utils.JsonResponseError(w, err)
		return
	}

	result, err := h.friendService.CreateBlock(r.Context(), input)
	if err != nil {
		log.Println("CreateRelation error", err)
		utils.JsonResponseError(w, err)
		return
	}

	utils.ResponseJson(w, http.StatusCreated, result)
	log.Println("----End API CreateBlock----")
}

// GetEmailReceive end point to create friend
func (h Handler) GetEmailReceive(w http.ResponseWriter, r *http.Request) {
	log.Println("----Start API GetEmailReceive----")

	relationReq := EmailReceiveRequest{}
	if err := json.NewDecoder(r.Body).Decode(&relationReq); err != nil {
		log.Println("GetEmailReceive: error when decode request ", err)
		utils.JsonResponseError(w, err)
		return
	}

	input, err := validateEmailReceiveInput(relationReq)
	if err != nil {
		log.Println("GetEmailReceive: error validate input ", err)
		utils.JsonResponseError(w, err)
		return
	}

	result, err := h.friendService.GetEmailReceive(r.Context(), input)
	if err != nil {
		log.Println("CreateRelation error", err)
		utils.JsonResponseError(w, err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, result)
	log.Println("----End API GetEmailReceive----")
}

// validateRelationInput function validate create friends request
func validateRelationInput(relationReq FriendsRequest) (friends.CreateRelationsInput, error) {
	//check of slice of friend is empty
	if len(relationReq.Friends) < 2 {
		return friends.CreateRelationsInput{}, ErrDataIsEmpty
	}

	requesterEmail := strings.TrimSpace(relationReq.Friends[0])
	if requesterEmail == "" {
		return friends.CreateRelationsInput{}, ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return friends.CreateRelationsInput{}, ErrInvalidEmail
	}

	addresseeEmail := strings.TrimSpace(relationReq.Friends[1])
	if addresseeEmail == "" {
		return friends.CreateRelationsInput{}, ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(addresseeEmail); err != nil {
		return friends.CreateRelationsInput{}, ErrInvalidEmail
	}

	//check email requester and addressee is the same
	if requesterEmail == addresseeEmail {
		return friends.CreateRelationsInput{}, ErrRequesterEmailAndAddresseeEmail
	}

	return friends.CreateRelationsInput{
		RequesterEmail: requesterEmail,
		AddresseeEmail: addresseeEmail,
	}, nil
}

// validateSubAndBlockRelationInput function validate create sub and block friends request
func validateSubAndBlockRelationInput(relationReq RelationFriendsRequest) (friends.CreateRelationsInput, error) {

	requesterEmail := strings.TrimSpace(relationReq.Requester)
	if requesterEmail == "" {
		return friends.CreateRelationsInput{}, ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return friends.CreateRelationsInput{}, ErrInvalidEmail
	}

	addresseeEmail := strings.TrimSpace(relationReq.Target)
	if addresseeEmail == "" {
		return friends.CreateRelationsInput{}, ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(addresseeEmail); err != nil {
		return friends.CreateRelationsInput{}, ErrInvalidEmail
	}

	if requesterEmail == addresseeEmail {
		return friends.CreateRelationsInput{}, ErrRequesterEmailAndAddresseeEmail
	}

	return friends.CreateRelationsInput{
		RequesterEmail: requesterEmail,
		AddresseeEmail: addresseeEmail,
	}, nil
}

// validateGetRelationInput function validate get friend request
func validateGetRelationInput(relationReq GetRelationRequest) (friends.GetAllFriendsInput, error) {
	requesterEmail := strings.TrimSpace(relationReq.Email)
	if requesterEmail == "" {
		return friends.GetAllFriendsInput{}, ErrEmailCannotBeBlank
	}
	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return friends.GetAllFriendsInput{}, ErrInvalidEmail
	}

	return friends.GetAllFriendsInput{
		Email: requesterEmail,
	}, nil
}

// validateRelationCommonInput function validate get common friend request
func validateRelationCommonInput(relationReq FriendsRequest) (friends.CommonFriendsInput, error) {
	//check if slice of friend is empty
	if len(relationReq.Friends) < 2 {
		return friends.CommonFriendsInput{}, ErrDataIsEmpty
	}
	requesterEmail := strings.TrimSpace(relationReq.Friends[0])
	if requesterEmail == "" {
		return friends.CommonFriendsInput{}, ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return friends.CommonFriendsInput{}, ErrInvalidEmail
	}

	addresseeEmail := strings.TrimSpace(relationReq.Friends[1])
	if addresseeEmail == "" {
		return friends.CommonFriendsInput{}, ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(addresseeEmail); err != nil {
		return friends.CommonFriendsInput{}, ErrInvalidEmail
	}

	return friends.CommonFriendsInput{
		FirstEmail:  requesterEmail,
		SecondEmail: addresseeEmail,
	}, nil
}

func validateEmailReceiveInput(relationReq EmailReceiveRequest) (friends.EmailReceiveInput, error) {
	requesterEmail := strings.TrimSpace(relationReq.Sender)
	if requesterEmail == "" {
		return friends.EmailReceiveInput{}, ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return friends.EmailReceiveInput{}, ErrInvalidEmail
	}

	return friends.EmailReceiveInput{
		Sender: requesterEmail,
		Text:   relationReq.Text,
	}, nil
}
