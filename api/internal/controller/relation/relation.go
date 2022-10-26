package relation

import (
	"context"
	"database/sql"
	"log"

	"github.com/friendsofgo/errors"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/pkg/utils"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/repository"
	models "github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/repository/orm/models"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/repository/relation"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/repository/user"
)

var (
	isRelationExistFunc         = isRelationExist
	isValidToCreateRelationFunc = isValidToCreateRelation
)

// RelationsService type contain repository needed
type RelationsService struct {
	relationsRepository repository.RelationsRepo
	userRepository      repository.UserRepo
	utils               utils.UtilsInf
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

// EmailReceiveInput param using for receive email
type EmailReceiveInput struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
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

// EmailReceiveResponse response for API email receive
type EmailReceiveResponse struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients"`
}

// RelationServ define function of relation
type RelationServ interface {
	CreateFriendRelation(ctx context.Context, input CreateRelationsInput) (CreateRelationsResponse, error)
	CreateSubscriptionRelation(ctx context.Context, input CreateRelationsInput) (CreateRelationsResponse, error)
	CreateBlockRelation(ctx context.Context, input CreateRelationsInput) (CreateRelationsResponse, error)
	GetAllFriendsOfUser(ctx context.Context, input GetAllFriendsInput) (FriendListResponse, error)
	GetCommonFriendList(ctx context.Context, input CommonFriendsInput) (FriendListResponse, error)
	GetEmailReceive(ctx context.Context, input EmailReceiveInput) (EmailReceiveResponse, error)
}

// NewRelationService create new relation service
func NewRelationService(db *sql.DB) RelationServ {
	return RelationsService{
		relationsRepository: relation.NewRelationsRepository(db),
		userRepository:      user.NewUserRepository(db),
	}
}

// GetAllFriendsOfUser implement function get all friend of user by input
func (serv RelationsService) GetAllFriendsOfUser(ctx context.Context, input GetAllFriendsInput) (FriendListResponse, error) {
	//get requesterId from request
	userGetFromReq, err := serv.userRepository.GetUserByEmail(ctx, input.Email)
	if err != nil {
		log.Println("GetAllFriendsOfUser: error get email from request ", err)
		return FriendListResponse{Success: false, Count: 0}, err
	}

	lstIdUser, err := serv.relationsRepository.GetRelationIDsOfUser(ctx, userGetFromReq.ID, utils.FriendRelation)
	if err != nil {
		log.Println("GetAllFriendsOfUser: error get list ids user ", err)
		return FriendListResponse{Success: false, Count: 0}, err
	}

	lstFriends, err := serv.userRepository.GetListEmailByIDs(ctx, lstIdUser)
	if err != nil {
		log.Println("GetAllFriendsOfUser: error get list email by user id ", err)
		return FriendListResponse{Success: false, Count: 0}, err
	}

	return FriendListResponse{
		Success: true,
		Friends: lstFriends,
		Count:   len(lstFriends),
	}, nil
}

// CreateFriendRelation function implement create friend relationship
func (serv RelationsService) CreateFriendRelation(ctx context.Context, input CreateRelationsInput) (CreateRelationsResponse, error) {
	//get requesterId from email request
	requester, err := serv.userRepository.GetUserByEmail(ctx, input.RequesterEmail)
	if err != nil {
		log.Println("CreateFriendRelation: error get email from request ", err)
		return CreateRelationsResponse{Success: false}, errors.New(utils.ErrMessageRequesterEmailFromRequest)
	}

	// requester email not exist
	if requester.Email == "" {
		log.Println("CreateFriendRelation: error requester email not exist ", err)
		return CreateRelationsResponse{Success: false}, errors.New(utils.ErrMessageRequesterEmailNotExist)
	}

	//get addresseeId from email request
	addressee, err := serv.userRepository.GetUserByEmail(ctx, input.AddresseeEmail)
	if err != nil {
		log.Println("CreateFriendRelation: error get email from request ", err)
		return CreateRelationsResponse{Success: false}, errors.New(utils.ErrMessageAddresseeEmailFromRequest)
	}
	// requester email not exist
	if addressee.Email == "" {
		log.Println("CreateFriendRelation: error addressee email not exist ", err)
		return CreateRelationsResponse{Success: false}, errors.New(utils.ErrMessageAddresseeEmailNotExist)
	}

	isValid, err := isValidToCreateRelationFunc(ctx, serv.relationsRepository, requester.ID, addressee.ID, utils.FriendRelation)
	if !isValid || err != nil {
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
	//define slice of list friend
	var listCommonFriend []string

	//get requesterId from email request
	firstUser, err := serv.userRepository.GetUserByEmail(ctx, input.FirstEmail)
	if err != nil {
		log.Println("GetCommonFriendList: error get email from request ", err)
		return FriendListResponse{Success: false, Count: 0}, err
	}

	//get first friend list
	firstIdsList, err := serv.relationsRepository.GetRelationIDsOfUser(ctx, firstUser.ID, utils.FriendRelation)
	if err != nil {
		log.Println("GetCommonFriendList: error get ids list ", err)
		return FriendListResponse{Success: false, Count: 0}, err
	}
	//get requesterId from email request
	secondUser, err := serv.userRepository.GetUserByEmail(ctx, input.SecondEmail)
	if err != nil {
		log.Println("GetCommonFriendList: error get email from request ", err)
		return FriendListResponse{Success: false, Count: 0}, err
	}
	//get second friend list
	secondIdsList, err := serv.relationsRepository.GetRelationIDsOfUser(ctx, secondUser.ID, utils.FriendRelation)
	if err != nil {
		log.Println("GetCommonFriendList: error get ids from email ", err)
		return FriendListResponse{Success: false, Count: 0}, err
	}
	//Intersection two list friend
	//listCommonIds := utils.Intersection(firstIdsList, secondIdsList)
	listCommonIds := utils.Utils{}.Intersection(firstIdsList, secondIdsList)

	listCommonFriend, err = serv.userRepository.GetListEmailByIDs(ctx, listCommonIds)

	if err != nil {
		log.Println("GetCommonFriendList: error get list email from list ids ", err)
		return FriendListResponse{Success: false, Count: 0}, err
	}
	return FriendListResponse{
		Success: true,
		Friends: listCommonFriend,
		Count:   len(listCommonFriend),
	}, nil

}

// isRelationExist function check relation is exists
func isRelationExist(ctx context.Context, repo repository.RelationsRepo, requesterID int, addresseeID int, relationType int) (bool, error) {
	isExistRelation, err := repo.IsRelationExist(ctx, requesterID, addresseeID, relationType)
	if err != nil {
		log.Println("CreateFriendRelation: error when check block relation ", err)
		return false, err
	}

	return isExistRelation, nil
}

// isRelationExist function check valid to create relation
func isValidToCreateRelation(ctx context.Context, repo repository.RelationsRepo, requesterID int, addresseeID int, relationType int) (bool, error) {
	isExistRelation, err := isRelationExistFunc(ctx, repo, requesterID, addresseeID, relationType)

	if err != nil {
		return false, err
	}

	if isExistRelation {
		return false, errors.New("Relation is exists")
	}

	isValid := false

	switch relationType {
	case utils.FriendRelation:
		isRequesterIDBlock, err := isRelationExistFunc(ctx, repo, requesterID, addresseeID, utils.Block)
		isAddresseeIDBlock, err := isRelationExistFunc(ctx, repo, requesterID, addresseeID, utils.Block)

		if err != nil {
			return false, err
		}

		if !isRequesterIDBlock && !isAddresseeIDBlock {
			isValid = true
		}
	case utils.Subscribe:
		isValid = true
	case utils.Block:
		isValid = true

	}
	if err != nil {
		return false, err
	}

	if isValid == false {
		return false, errors.New("Unable to create relation")
	}
	return isValid, nil
}

// CreateSubscriptionRelation function create subscription relation
func (serv RelationsService) CreateSubscriptionRelation(ctx context.Context, input CreateRelationsInput) (CreateRelationsResponse, error) {
	//get requesterId from email request
	requester, err := serv.userRepository.GetUserByEmail(ctx, input.RequesterEmail)
	if err != nil {
		log.Println("CreateSubscriptionRelation: error get email from request ", err)
		return CreateRelationsResponse{Success: false}, err
	}

	if requester.Email == "" {
		log.Println("CreateSubscriptionRelation: error email requester not exist ", err)
		return CreateRelationsResponse{Success: false}, errors.New(utils.ErrMessageRequesterEmailNotExist)
	}

	//get addresseeId from email request
	target, err := serv.userRepository.GetUserByEmail(ctx, input.AddresseeEmail)
	if err != nil {
		log.Println("CreateFriendRelation: error get email from request ", err)
		return CreateRelationsResponse{Success: false}, err
	}

	if target.Email == "" {
		log.Println("CreateSubscriptionRelation: error email target not exist ", err)
		return CreateRelationsResponse{Success: false}, errors.New(utils.ErrMessageAddresseeEmailNotExist)
	}

	//check if user blocked or user is friend return false
	isValid, err := isValidToCreateRelationFunc(ctx, serv.relationsRepository, requester.ID, target.ID, utils.Subscribe)
	//if user has relation block or friend return false
	if !isValid || err != nil {
		return CreateRelationsResponse{Success: false}, err
	}

	// insert relation friend two email, relationType 1 is friend
	relationFriendInput := models.Relation{
		RequesterID:    requester.ID,
		AddresseeID:    target.ID,
		RequesterEmail: requester.Email,
		AddresseeEmail: target.Email,
		RelationType:   utils.Subscribe,
	}
	result, err := serv.relationsRepository.CreateRelation(ctx, relationFriendInput)
	if err != nil {
		return CreateRelationsResponse{Success: false}, err
	}

	return CreateRelationsResponse{Success: result}, nil
}

// CreateBlockRelation function create block relation
func (serv RelationsService) CreateBlockRelation(ctx context.Context, input CreateRelationsInput) (CreateRelationsResponse, error) {
	//get requesterId from email request
	requester, err := serv.userRepository.GetUserByEmail(ctx, input.RequesterEmail)
	if err != nil {
		log.Println("CreateSubscriptionRelation: error get email from request ", err)
		return CreateRelationsResponse{Success: false}, err
	}

	//get addresseeId from email request
	target, err := serv.userRepository.GetUserByEmail(ctx, input.AddresseeEmail)
	if err != nil {
		log.Println("CreateFriendRelation: error get email from request ", err)
		return CreateRelationsResponse{Success: false}, err
	}

	//check if user blocked or user is friend return false
	isValid, err := isValidToCreateRelationFunc(ctx, serv.relationsRepository, requester.ID, target.ID, utils.Block)
	//if user has relation block or friend return false
	if !isValid || err != nil {
		return CreateRelationsResponse{Success: false}, err
	}

	// insert relation friend two email, relationType 1 is friend
	relationFriendInput := models.Relation{
		RequesterID:    requester.ID,
		AddresseeID:    target.ID,
		RequesterEmail: requester.Email,
		AddresseeEmail: target.Email,
		RelationType:   utils.Block,
	}
	result, err := serv.relationsRepository.CreateRelation(ctx, relationFriendInput)
	if err != nil {
		return CreateRelationsResponse{Success: false}, err
	}

	// if insert block is success then delete subcribe relation
	err = serv.relationsRepository.DeleteRelation(ctx, requester.ID, target.ID, utils.Subscribe)

	if err != nil {
		return CreateRelationsResponse{Success: false}, err
	}

	return CreateRelationsResponse{Success: result}, nil
}

// GetEmailReceive function email receive info
func (serv RelationsService) GetEmailReceive(ctx context.Context, input EmailReceiveInput) (EmailReceiveResponse, error) {
	//get requesterId from email request
	requester, err := serv.userRepository.GetUserByEmail(ctx, input.Sender)

	if err != nil {
		log.Println("GetEmailReceive: error when get user by email", err)
		return EmailReceiveResponse{Success: false}, errors.New(utils.ErrMessageEmailNotExist)
	}

	//get lst id friend
	friendIDs, err := serv.relationsRepository.GetRelationIDsOfUser(ctx, requester.ID, utils.FriendRelation)

	if err != nil {
		return EmailReceiveResponse{Success: false}, err
	}

	// get lst id subscription
	subscribeIDs, err := serv.relationsRepository.GetRequesterIDRelation(ctx, requester.ID, utils.Subscribe)
	if err != nil {
		return EmailReceiveResponse{Success: false}, err
	}

	// get user who mentioned in text
	mentionedEmails := utils.Utils{}.FindEmailFromText(input.Text)
	userIDsMentionedEmails, err := serv.userRepository.GetUserIDsByEmail(ctx, mentionedEmails)

	// get requester id of block relation
	blocksID, err := serv.relationsRepository.GetRequesterIDRelation(ctx, requester.ID, utils.Block)

	if err != nil {
		return EmailReceiveResponse{Success: false}, err
	}

	// user can receiver update
	receiversIDnoMentioned := append(friendIDs, subscribeIDs...)
	receiversID := utils.Utils{}.GetReceiverID(utils.Utils{}.UniqueSlice(append(receiversIDnoMentioned, userIDsMentionedEmails...)), append(blocksID, requester.ID))

	emails, err := serv.userRepository.GetListEmailByIDs(ctx, receiversID)

	if err != nil {
		return EmailReceiveResponse{Success: false}, err
	}
	return EmailReceiveResponse{
		Success:    true,
		Recipients: emails,
	}, nil
}
