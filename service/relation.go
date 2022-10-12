package service

import (
	"context"
	"database/sql"
	models "github.com/quangpham789/golang-assessment/models"
	"github.com/quangpham789/golang-assessment/repository"
	"github.com/quangpham789/golang-assessment/repository/relation"
	"github.com/quangpham789/golang-assessment/repository/user"
	"github.com/quangpham789/golang-assessment/utils"
	"log"
)

// RelationsService type contain repository needed
type RelationsService struct {
	relationsRepository repository.RelationsRepo
	userRepository      repository.UserRepo
}

// CreateRelationsInput param using for create friend relation
type CreateRelationsInput struct {
	RequesterEmail string
	AddresseeEmail string
}

// GetAllFriendsInput param using for get list friend of user
type GetAllFriendsInput struct {
	Email string
}

// CommonFriendsInput param using for get list common friend of two user
type CommonFriendsInput struct {
	FirstEmail  string
	SecondEmail string
}

// CreateRelationsResponse response for API create a friend
type CreateRelationsResponse struct {
	Success bool `json:"success"`
}

// FriendListResponse response for API get list friend of user
// and get common friend of two user
type FriendListResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}

// RelationServ define function of relation
type RelationServ interface {
	CreateFriendRelation(ctx context.Context, input CreateRelationsInput) (CreateRelationsResponse, error)
	GetAllFriendsOfUser(ctx context.Context, input GetAllFriendsInput) (FriendListResponse, error)
	GetCommonFriendList(ctx context.Context, input CommonFriendsInput) (FriendListResponse, error)
	isRelation(ctx context.Context, requesterID int, addresseeID int) bool
}

// GetAllFriendsOfUser implement function get all friend of user by input
func (serv RelationsService) GetAllFriendsOfUser(ctx context.Context, input GetAllFriendsInput) (FriendListResponse, error) {
	//get requesterId from request
	user, err := serv.userRepository.GetUserByEmail(ctx, input.Email)
	if err != nil {
		log.Println("GetAllFriendsOfUser: error get email from request ", err)
		return FriendListResponse{Success: false, Count: 0}, err
	}

	//call repository to get all friend
	relationFriend, err := serv.relationsRepository.GetAllRelationFriendOfUser(ctx, user.ID)
	if err != nil {
		log.Println("GetAllFriendsOfUser: error get list friend ", err)
		return FriendListResponse{Success: false, Count: 0}, err
	}
	//check if relationFriend length is 0
	if len(relationFriend) == 0 {
		log.Println("GetAllFriendsOfUser: error get list friend ", err)
		return FriendListResponse{Success: false, Count: 0}, err
	}
	var listFriends []string
	for _, friend := range relationFriend {
		listFriends = append(listFriends, friend.AddresseeEmail)
	}
	return FriendListResponse{
		Success: true,
		Friends: listFriends,
		Count:   len(listFriends),
	}, nil
}

// CreateFriendRelation function implement create friend relationship
func (serv RelationsService) CreateFriendRelation(ctx context.Context, input CreateRelationsInput) (CreateRelationsResponse, error) {
	//get requesterId from email request
	requester, err := serv.userRepository.GetUserByEmail(ctx, input.RequesterEmail)
	if err != nil {
		log.Println("CreateFriendRelation: error get email from request ", err)
		return CreateRelationsResponse{Success: false}, err
	}

	//get addresseeId from email request
	addressee, err := serv.userRepository.GetUserByEmail(ctx, input.AddresseeEmail)
	if err != nil {
		log.Println("CreateFriendRelation: error get email from request ", err)
		return CreateRelationsResponse{Success: false}, err
	}

	//check if user blocked or user is friend return false
	isRelation := serv.isRelation(ctx, requester.ID, addressee.ID)
	//if user has relation block or friend return false
	if !isRelation {
		return CreateRelationsResponse{Success: false}, err
	}

	// insert relation friend two email, relationType 1 is friend
	relationFriendInput := models.Relation{
		RequesterID:    requester.ID,
		AddresseeID:    addressee.ID,
		RequesterEmail: requester.Email,
		AddresseeEmail: addressee.Email,
		RelationType:   utils.FriendRelation,
	}
	result, err := serv.relationsRepository.CreateRelation(ctx, relationFriendInput)
	if err != nil {
		return CreateRelationsResponse{Success: false}, err
	}

	return CreateRelationsResponse{Success: result}, nil
}

// GetCommonFriendList function implement get common friend
func (serv RelationsService) GetCommonFriendList(ctx context.Context, input CommonFriendsInput) (FriendListResponse, error) {
	//get requesterId from email request
	firstUser, err := serv.userRepository.GetUserByEmail(ctx, input.FirstEmail)
	if err != nil {
		log.Println("GetCommonFriendList: error get email from request ", err)
		return FriendListResponse{Success: false, Count: 0}, err
	}

	secondUser, err := serv.userRepository.GetUserByEmail(ctx, input.SecondEmail)
	if err != nil {
		log.Println("GetCommonFriendList: error get email from request ", err)
		return FriendListResponse{Success: false, Count: 0}, err
	}

	//call service get all friend
	relationFriend, err := serv.relationsRepository.GetCommonFriend(ctx, firstUser.ID, secondUser.ID)
	if err != nil {
		log.Println("GetCommonFriendList: error get list friend ", err)
		return FriendListResponse{Success: false, Count: 0}, err
	}
	//check if relationFriend length is 0
	if len(relationFriend) == 0 {
		log.Println("GetCommonFriendList: error list friend is empty ", err)
		return FriendListResponse{Success: false, Count: 0}, err
	}
	var listFriends []string
	// append list email from database to list friend response
	for _, friend := range relationFriend {
		listFriends = append(listFriends, friend.AddresseeEmail)
	}
	return FriendListResponse{
		Success: true,
		Friends: listFriends,
		Count:   len(listFriends),
	}, nil

}

func (serv RelationsService) isRelation(ctx context.Context, requesterID int, addresseeID int) bool {
	relationBlockResult, err := serv.relationsRepository.GetRelationByIdsAndType(ctx, requesterID, addresseeID, utils.Block)
	if relationBlockResult.RelationType == utils.Block || err != nil {
		log.Println("CreateFriendRelation: error when check block relation ", err)
		return false
	}
	relationFriendResult, _ := serv.relationsRepository.GetRelationByIdsAndType(ctx, requesterID, addresseeID, utils.FriendRelation)
	if relationFriendResult.RelationType == utils.FriendRelation {
		log.Println("CreateFriendRelation: error already is friend ", err)
		return false
	}

	relationFriendResultReverse, _ := serv.relationsRepository.GetRelationByIdsAndType(ctx, addresseeID, requesterID, utils.FriendRelation)
	if relationFriendResultReverse.RelationType == utils.FriendRelation {
		log.Println("CreateFriendRelation: error already is friend ", err)
		return false
	}

	return true
}

// NewRelationService create new relation service
func NewRelationService(db *sql.DB) RelationServ {
	return RelationsService{
		relationsRepository: relation.NewRelationsRepository(db),
		userRepository:      user.NewUserRepository(db),
	}
}
