package relation

import (
	"database/sql"
	"encoding/json"
	"github.com/quangpham789/golang-assessment/handler/errors"
	"github.com/quangpham789/golang-assessment/service/relation"
	"github.com/quangpham789/golang-assessment/utils"
	"log"
	"net/http"
	"net/mail"
	"strings"
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

// RelationFriendRequest request to create friend or get common friend list
type RelationFriendRequest struct {
	Friends []string `json:"friends"`
}

// RelationSubAndBlockRequest request to create friend or get common friend list
type RelationSubAndBlockRequest struct {
	Requestor string `json:"requestor"`
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
	// Convert body request to struct of Handler
	relationReq := RelationFriendRequest{}
	if err := json.NewDecoder(r.Body).Decode(&relationReq); err != nil {
		errors.JsonResponseError(w, err)
		return
	}

	// Validate body relation request
	input, err := validateRelationInput(relationReq)
	if err != nil {
		errors.JsonResponseError(w, err)
		return
	}

	// Call service create relation
	result, err := relations.relationsService.CreateFriendRelation(r.Context(), input)
	if err != nil {
		log.Println("CreateRelation error", err)
		utils.JsonResponse(w, http.StatusForbidden, result)
	}
	utils.JsonResponse(w, http.StatusCreated, result)
}

// GetAllFriendOfUser end point to get list friend of user
func (relations RelationsHandler) GetAllFriendOfUser(w http.ResponseWriter, r *http.Request) {
	// Convert body request to struct of Handler
	getRelationReq := GetRelationRequest{}
	if err := json.NewDecoder(r.Body).Decode(&getRelationReq); err != nil {
		errors.JsonResponseError(w, err)
		return
	}

	// Validate body relation request
	input, err := validateGetRelationInput(getRelationReq)
	if err != nil {
		errors.JsonResponseError(w, err)
		return
	}

	// Call service get list of friend
	result, err := relations.relationsService.GetAllFriendsOfUser(r.Context(), input)
	if err != nil {
		log.Println("GetAllFriendOfUser error", err)
		utils.JsonResponse(w, http.StatusForbidden, result)
	}
	utils.JsonResponse(w, http.StatusOK, result)
}

func (relations RelationsHandler) GetCommonFriend(w http.ResponseWriter, r *http.Request) {
	// Convert body request to struct of Handler
	getRelationReq := RelationFriendRequest{}
	if err := json.NewDecoder(r.Body).Decode(&getRelationReq); err != nil {
		errors.JsonResponseError(w, err)
		return
	}

	// Validate relation request
	input, err := validateRelationCommonInput(getRelationReq)
	if err != nil {
		errors.JsonResponseError(w, err)
		return
	}

	// Call service get common friend list
	result, err := relations.relationsService.GetCommonFriendList(r.Context(), input)
	if err != nil {
		log.Println("GetCommonFriend call service error", err)
		utils.JsonResponse(w, http.StatusForbidden, result)
	}
	utils.JsonResponse(w, http.StatusOK, result)
}

// CreateSubscriptionRelation end point to create friend relation
func (relations RelationsHandler) CreateSubscriptionRelation(w http.ResponseWriter, r *http.Request) {
	// Convert body request to struct of Handler
	relationReq := RelationSubAndBlockRequest{}
	if err := json.NewDecoder(r.Body).Decode(&relationReq); err != nil {
		errors.JsonResponseError(w, err)
		return
	}

	// Validate body relation request
	input, err := validateSubAndBlockRelationInput(relationReq)
	if err != nil {
		errors.JsonResponseError(w, err)
		return
	}

	// Call service create relation
	result, err := relations.relationsService.CreateSubscriptionRelation(r.Context(), input)
	if err != nil {
		log.Println("CreateRelation error", err)
		utils.JsonResponse(w, http.StatusForbidden, result)
	}
	utils.JsonResponse(w, http.StatusCreated, result)
}

// CreateBlockRelation end point to create friend relation
func (relations RelationsHandler) CreateBlockRelation(w http.ResponseWriter, r *http.Request) {
	// Convert body request to struct of Handler
	relationReq := RelationSubAndBlockRequest{}
	if err := json.NewDecoder(r.Body).Decode(&relationReq); err != nil {
		errors.JsonResponseError(w, err)
		return
	}

	// Validate body relation request
	input, err := validateSubAndBlockRelationInput(relationReq)
	if err != nil {
		errors.JsonResponseError(w, err)
		return
	}

	// Call service create relation
	result, err := relations.relationsService.CreateBlockRelation(r.Context(), input)
	if err != nil {
		log.Println("CreateRelation error", err)
		utils.JsonResponse(w, http.StatusForbidden, result)
	}
	utils.JsonResponse(w, http.StatusCreated, result)
}

// GetEmailReceive end point to create friend relation
func (relations RelationsHandler) GetEmailReceive(w http.ResponseWriter, r *http.Request) {
	// Convert body request to struct of Handler
	relationReq := EmailReceiveRequest{}
	if err := json.NewDecoder(r.Body).Decode(&relationReq); err != nil {
		errors.JsonResponseError(w, err)
		return
	}

	// Validate body relation request
	input, err := validateEmailReceiveInput(relationReq)
	if err != nil {
		errors.JsonResponseError(w, err)
		return
	}

	// Call service create relation
	result, err := relations.relationsService.GetEmailReceive(r.Context(), input)
	if err != nil {
		log.Println("CreateRelation error", err)
		utils.JsonResponse(w, http.StatusForbidden, result)
	}
	utils.JsonResponse(w, http.StatusCreated, result)
}

func validateRelationInput(relationReq RelationFriendRequest) (relation.CreateRelationsInput, error) {
	//check of slice of friend is empty
	if len(relationReq.Friends) < 2 {
		return relation.CreateRelationsInput{}, errors.ErrDataIsEmpty
	}
	requesterEmail := strings.TrimSpace(relationReq.Friends[0])
	if requesterEmail == "" {
		return relation.CreateRelationsInput{}, errors.ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return relation.CreateRelationsInput{}, errors.ErrInvalidEmail
	}

	addresseeEmail := strings.TrimSpace(relationReq.Friends[1])
	if addresseeEmail == "" {
		return relation.CreateRelationsInput{}, errors.ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(addresseeEmail); err != nil {
		return relation.CreateRelationsInput{}, errors.ErrInvalidEmail
	}

	return relation.CreateRelationsInput{
		RequesterEmail: requesterEmail,
		AddresseeEmail: addresseeEmail,
	}, nil
}

func validateSubAndBlockRelationInput(relationReq RelationSubAndBlockRequest) (relation.CreateRelationsInput, error) {

	requesterEmail := strings.TrimSpace(relationReq.Requestor)
	if requesterEmail == "" {
		return relation.CreateRelationsInput{}, errors.ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return relation.CreateRelationsInput{}, errors.ErrInvalidEmail
	}

	addresseeEmail := strings.TrimSpace(relationReq.Target)
	if addresseeEmail == "" {
		return relation.CreateRelationsInput{}, errors.ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(addresseeEmail); err != nil {
		return relation.CreateRelationsInput{}, errors.ErrInvalidEmail
	}

	return relation.CreateRelationsInput{
		RequesterEmail: requesterEmail,
		AddresseeEmail: addresseeEmail,
	}, nil
}

func validateGetRelationInput(relationReq GetRelationRequest) (relation.GetAllFriendsInput, error) {
	requesterEmail := strings.TrimSpace(relationReq.Email)
	if requesterEmail == "" {
		return relation.GetAllFriendsInput{}, errors.ErrEmailCannotBeBlank
	}

	return relation.GetAllFriendsInput{
		Email: requesterEmail,
	}, nil
}

func validateRelationCommonInput(relationReq RelationFriendRequest) (relation.CommonFriendsInput, error) {
	//check if slice of friend is empty
	if len(relationReq.Friends) < 2 {
		return relation.CommonFriendsInput{}, errors.ErrDataIsEmpty
	}
	requesterEmail := strings.TrimSpace(relationReq.Friends[0])
	if requesterEmail == "" {
		return relation.CommonFriendsInput{}, errors.ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return relation.CommonFriendsInput{}, errors.ErrInvalidEmail
	}

	addresseeEmail := strings.TrimSpace(relationReq.Friends[1])
	if addresseeEmail == "" {
		return relation.CommonFriendsInput{}, errors.ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(addresseeEmail); err != nil {
		return relation.CommonFriendsInput{}, errors.ErrInvalidEmail
	}

	return relation.CommonFriendsInput{
		FirstEmail:  requesterEmail,
		SecondEmail: addresseeEmail,
	}, nil
}

func validateEmailReceiveInput(relationReq EmailReceiveRequest) (relation.EmailReceiveInput, error) {

	requesterEmail := strings.TrimSpace(relationReq.Sender)
	if requesterEmail == "" {
		return relation.EmailReceiveInput{}, errors.ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return relation.EmailReceiveInput{}, errors.ErrInvalidEmail
	}

	return relation.EmailReceiveInput{
		Sender: requesterEmail,
		Text:   relationReq.Text,
	}, nil
}
