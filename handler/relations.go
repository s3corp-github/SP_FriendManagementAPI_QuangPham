package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/quangpham789/golang-assessment/service"
	"github.com/quangpham789/golang-assessment/utils"
	"log"
	"net/http"
	"net/mail"
	"strings"
)

// RelationsHandler create relation handler contain relation service
type RelationsHandler struct {
	relationsService service.RelationServ
}

// NewRelationsHandler create relation handler contain RelationsHandler
func NewRelationsHandler(db *sql.DB) RelationsHandler {
	return RelationsHandler{
		relationsService: service.NewRelationService(db),
	}
}

// RelationRequest request to create friend or get common friend list
type RelationRequest struct {
	Friends []string `json:"friends"`
}

// GetRelationRequest request to get list friend of user
type GetRelationRequest struct {
	Email string `json:"email"`
}

// CreateRelation end point to create friend relation
func (relations RelationsHandler) CreateRelation(w http.ResponseWriter, r *http.Request) {
	// Convert body request to struct of Handler
	relationReq := RelationRequest{}
	if err := json.NewDecoder(r.Body).Decode(&relationReq); err != nil {
		JsonResponseError(w, err)
		return
	}

	// Validate body relation request
	input, err := validateRelationInput(relationReq)
	if err != nil {
		JsonResponseError(w, err)
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
		JsonResponseError(w, err)
		return
	}

	// Validate body relation request
	input, err := validateGetRelationInput(getRelationReq)
	if err != nil {
		JsonResponseError(w, err)
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
	getRelationReq := RelationRequest{}
	if err := json.NewDecoder(r.Body).Decode(&getRelationReq); err != nil {
		JsonResponseError(w, err)
		return
	}

	// Validate relation request
	input, err := validateRelationCommonInput(getRelationReq)
	if err != nil {
		JsonResponseError(w, err)
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

func validateRelationInput(relationReq RelationRequest) (service.CreateRelationsInput, error) {
	//check of slice of friend is empty
	if len(relationReq.Friends) < 2 {
		return service.CreateRelationsInput{}, errDataIsEmpty
	}
	requesterEmail := strings.TrimSpace(relationReq.Friends[0])
	if requesterEmail == "" {
		return service.CreateRelationsInput{}, errEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return service.CreateRelationsInput{}, errInvalidEmail
	}

	addresseeEmail := strings.TrimSpace(relationReq.Friends[1])
	if addresseeEmail == "" {
		return service.CreateRelationsInput{}, errEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(addresseeEmail); err != nil {
		return service.CreateRelationsInput{}, errInvalidEmail
	}

	return service.CreateRelationsInput{
		RequesterEmail: requesterEmail,
		AddresseeEmail: addresseeEmail,
	}, nil
}

func validateGetRelationInput(relationReq GetRelationRequest) (service.GetAllFriendsInput, error) {
	requesterEmail := strings.TrimSpace(relationReq.Email)
	if requesterEmail == "" {
		return service.GetAllFriendsInput{}, errEmailCannotBeBlank
	}

	return service.GetAllFriendsInput{
		Email: requesterEmail,
	}, nil
}

func validateRelationCommonInput(relationReq RelationRequest) (service.CommonFriendsInput, error) {
	//check if slice of friend is empty
	if len(relationReq.Friends) < 2 {
		return service.CommonFriendsInput{}, errDataIsEmpty
	}
	requesterEmail := strings.TrimSpace(relationReq.Friends[0])
	if requesterEmail == "" {
		return service.CommonFriendsInput{}, errEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return service.CommonFriendsInput{}, errInvalidEmail
	}

	addresseeEmail := strings.TrimSpace(relationReq.Friends[1])
	if addresseeEmail == "" {
		return service.CommonFriendsInput{}, errEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(addresseeEmail); err != nil {
		return service.CommonFriendsInput{}, errInvalidEmail
	}

	return service.CommonFriendsInput{
		FirstEmail:  requesterEmail,
		SecondEmail: addresseeEmail,
	}, nil
}
