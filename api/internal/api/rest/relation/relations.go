package relation

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/mail"
	"strings"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/api/rest"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/controller/relation"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/pkg/utils"
)

// RelationsHandler create relation handler contain relation service
type RelationsHandler struct {
	relationsService relation.RelationServ
}

// NewRelationsHandler create relation handler contain RelationsHandler
func NewRelationsHandler(db *sql.DB) RelationsHandler {
	return RelationsHandler{
		relationsService: relation.NewRelationService(db),
	}
}

// RelationsFriendRequest request to create friend or get common friend list
type RelationsFriendRequest struct {
	Friends []string `json:"friends"`
}

// RelationsSubAndBlockRequest request to create friend or get common friend list
type RelationsSubAndBlockRequest struct {
	Requester string `json:"requester"`
	Target    string `json:"target"`
}

// GetRelationRequest request to get list friend of user
type GetRelationRequest struct {
	Email string `json:"email"`
}

// EmailReceiveRequest request to get email receive
type EmailReceiveRequest struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

// CreateFriendsRelation end point to create friend relation
func (relations RelationsHandler) CreateFriendsRelation(w http.ResponseWriter, r *http.Request) {
	log.Println("----Start API CreateFriendsRelation----")

	relationReq := RelationsFriendRequest{}
	if err := json.NewDecoder(r.Body).Decode(&relationReq); err != nil {
		log.Println("CreateFriendsRelation: error when decode request ", err)
		rest.JsonResponseError(w, err)
		return
	}

	input, err := validateRelationInput(relationReq)
	if err != nil {
		log.Println("CreateFriendsRelation: error validate input ", err)
		rest.JsonResponseError(w, err)
		return
	}

	result, err := relations.relationsService.CreateFriendRelation(r.Context(), input)
	if err != nil {
		log.Println("CreateRelation error", err)
		rest.JsonResponseError(w, err)
		return
	}
	utils.ResponseJson(w, http.StatusCreated, result)
	log.Println("----End API CreateFriendsRelation----")
}

// GetAllFriendOfUser end point to get list friend of user
func (relations RelationsHandler) GetAllFriendOfUser(w http.ResponseWriter, r *http.Request) {
	log.Println("----Start API GetAllFriendOfUser----")
	getRelationReq := GetRelationRequest{}
	if err := json.NewDecoder(r.Body).Decode(&getRelationReq); err != nil {
		log.Println("GetAllFriendOfUser: error when decode request ", err)
		rest.JsonResponseError(w, err)
		return
	}

	input, err := validateGetRelationInput(getRelationReq)
	if err != nil {
		log.Println("GetAllFriendOfUser: error validate input ", err)
		rest.JsonResponseError(w, err)
		return
	}

	result, err := relations.relationsService.GetAllFriendsOfUser(r.Context(), input)
	if err != nil {
		log.Println("GetAllFriendOfUser error", err)
		rest.JsonResponseError(w, err)
		return
	}
	utils.ResponseJson(w, http.StatusOK, result)
	log.Println("----End API GetAllFriendOfUser----")
}

// GetCommonFriend end point to get list common friend
func (relations RelationsHandler) GetCommonFriend(w http.ResponseWriter, r *http.Request) {
	log.Println("----Start API GetCommonFriend----")
	getRelationReq := RelationsFriendRequest{}
	if err := json.NewDecoder(r.Body).Decode(&getRelationReq); err != nil {
		log.Println("GetCommonFriend: error when decode request ", err)
		rest.JsonResponseError(w, err)
		return
	}

	input, err := validateRelationCommonInput(getRelationReq)
	if err != nil {
		log.Println("GetCommonFriend: error validate input ", err)
		rest.JsonResponseError(w, err)
		return
	}

	result, err := relations.relationsService.GetCommonFriendList(r.Context(), input)
	if err != nil {
		rest.JsonResponseError(w, err)
		return
	}
	utils.ResponseJson(w, http.StatusOK, result)
	log.Println("----End API GetCommonFriend----")
}

// CreateSubscriptionRelation end point to create friend relation
func (relations RelationsHandler) CreateSubscriptionRelation(w http.ResponseWriter, r *http.Request) {
	log.Println("----Start API CreateSubscriptionRelation----")
	relationReq := RelationsSubAndBlockRequest{}
	if err := json.NewDecoder(r.Body).Decode(&relationReq); err != nil {
		log.Println("CreateSubscriptionRelation: error when decode request ", err)
		rest.JsonResponseError(w, err)
		return
	}

	input, err := validateSubAndBlockRelationInput(relationReq)
	if err != nil {
		log.Println("CreateSubscriptionRelation: error validate input ", err)
		rest.JsonResponseError(w, err)
		return
	}

	result, err := relations.relationsService.CreateSubscriptionRelation(r.Context(), input)
	if err != nil {
		log.Println("CreateRelation error", err)
		rest.JsonResponseError(w, err)
		return
	}
	utils.ResponseJson(w, http.StatusCreated, result)
	log.Println("----End API CreateSubscriptionRelation----")
}

// CreateBlockRelation end point to create friend relation
func (relations RelationsHandler) CreateBlockRelation(w http.ResponseWriter, r *http.Request) {
	log.Println("----Start API CreateBlockRelation----")
	relationReq := RelationsSubAndBlockRequest{}
	if err := json.NewDecoder(r.Body).Decode(&relationReq); err != nil {
		log.Println("CreateBlockRelation: error when decode request ", err)
		rest.JsonResponseError(w, err)
		return
	}

	input, err := validateSubAndBlockRelationInput(relationReq)
	if err != nil {
		log.Println("CreateBlockRelation: error validate input ", err)
		rest.JsonResponseError(w, err)
		return
	}

	result, err := relations.relationsService.CreateBlockRelation(r.Context(), input)
	if err != nil {
		log.Println("CreateRelation error", err)
		rest.JsonResponseError(w, err)
		return
	}
	utils.ResponseJson(w, http.StatusCreated, result)
	log.Println("----End API CreateBlockRelation----")
}

// GetEmailReceive end point to create friend relation
func (relations RelationsHandler) GetEmailReceive(w http.ResponseWriter, r *http.Request) {
	log.Println("----Start API GetEmailReceive----")
	relationReq := EmailReceiveRequest{}
	if err := json.NewDecoder(r.Body).Decode(&relationReq); err != nil {
		log.Println("GetEmailReceive: error when decode request ", err)
		rest.JsonResponseError(w, err)
		return
	}

	input, err := validateEmailReceiveInput(relationReq)
	if err != nil {
		log.Println("GetEmailReceive: error validate input ", err)
		rest.JsonResponseError(w, err)
		return
	}

	result, err := relations.relationsService.GetEmailReceive(r.Context(), input)
	if err != nil {
		log.Println("CreateRelation error", err)
		rest.JsonResponseError(w, err)
		return
	}
	utils.ResponseJson(w, http.StatusOK, result)
	log.Println("----End API GetEmailReceive----")
}

// validateRelationInput function validate create relation request
func validateRelationInput(relationReq RelationsFriendRequest) (relation.CreateRelationsInput, error) {
	//check of slice of friend is empty
	if len(relationReq.Friends) < 2 {
		return relation.CreateRelationsInput{}, rest.ErrDataIsEmpty
	}
	requesterEmail := strings.TrimSpace(relationReq.Friends[0])
	if requesterEmail == "" {
		return relation.CreateRelationsInput{}, rest.ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return relation.CreateRelationsInput{}, rest.ErrInvalidEmail
	}

	addresseeEmail := strings.TrimSpace(relationReq.Friends[1])
	if addresseeEmail == "" {
		return relation.CreateRelationsInput{}, rest.ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(addresseeEmail); err != nil {
		return relation.CreateRelationsInput{}, rest.ErrInvalidEmail
	}

	//check email requester and addressee is the same
	if requesterEmail == addresseeEmail {
		return relation.CreateRelationsInput{}, rest.ErrRequesterEmailAndAddresseeEmail
	}

	return relation.CreateRelationsInput{
		RequesterEmail: requesterEmail,
		AddresseeEmail: addresseeEmail,
	}, nil
}

// validateSubAndBlockRelationInput function validate create sub and block relation request
func validateSubAndBlockRelationInput(relationReq RelationsSubAndBlockRequest) (relation.CreateRelationsInput, error) {

	requesterEmail := strings.TrimSpace(relationReq.Requester)
	if requesterEmail == "" {
		return relation.CreateRelationsInput{}, rest.ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return relation.CreateRelationsInput{}, rest.ErrInvalidEmail
	}

	addresseeEmail := strings.TrimSpace(relationReq.Target)
	if addresseeEmail == "" {
		return relation.CreateRelationsInput{}, rest.ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(addresseeEmail); err != nil {
		return relation.CreateRelationsInput{}, rest.ErrInvalidEmail
	}

	if requesterEmail == addresseeEmail {
		return relation.CreateRelationsInput{}, rest.ErrRequesterEmailAndAddresseeEmail
	}

	return relation.CreateRelationsInput{
		RequesterEmail: requesterEmail,
		AddresseeEmail: addresseeEmail,
	}, nil
}

// validateGetRelationInput function validate get friend request
func validateGetRelationInput(relationReq GetRelationRequest) (relation.GetAllFriendsInput, error) {
	requesterEmail := strings.TrimSpace(relationReq.Email)
	if requesterEmail == "" {
		return relation.GetAllFriendsInput{}, rest.ErrEmailCannotBeBlank
	}
	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return relation.GetAllFriendsInput{}, rest.ErrInvalidEmail
	}

	return relation.GetAllFriendsInput{
		Email: requesterEmail,
	}, nil
}

// validateRelationCommonInput function validate get common friend request
func validateRelationCommonInput(relationReq RelationsFriendRequest) (relation.CommonFriendsInput, error) {
	//check if slice of friend is empty
	if len(relationReq.Friends) < 2 {
		return relation.CommonFriendsInput{}, rest.ErrDataIsEmpty
	}
	requesterEmail := strings.TrimSpace(relationReq.Friends[0])
	if requesterEmail == "" {
		return relation.CommonFriendsInput{}, rest.ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return relation.CommonFriendsInput{}, rest.ErrInvalidEmail
	}

	addresseeEmail := strings.TrimSpace(relationReq.Friends[1])
	if addresseeEmail == "" {
		return relation.CommonFriendsInput{}, rest.ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(addresseeEmail); err != nil {
		return relation.CommonFriendsInput{}, rest.ErrInvalidEmail
	}

	return relation.CommonFriendsInput{
		FirstEmail:  requesterEmail,
		SecondEmail: addresseeEmail,
	}, nil
}

func validateEmailReceiveInput(relationReq EmailReceiveRequest) (relation.EmailReceiveInput, error) {
	requesterEmail := strings.TrimSpace(relationReq.Sender)
	if requesterEmail == "" {
		return relation.EmailReceiveInput{}, rest.ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return relation.EmailReceiveInput{}, rest.ErrInvalidEmail
	}

	return relation.EmailReceiveInput{
		Sender: requesterEmail,
		Text:   relationReq.Text,
	}, nil
}
