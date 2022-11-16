package friends

import (
	"context"

	"github.com/friendsofgo/errors"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/utils"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/repository"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/models"
	"github.com/volatiletech/null/v8"
)

var (
	isRelationExistFunc         = isRelationExist
	isValidToCreateRelationFunc = isValidToCreateRelation
)

// CreateRelationsInput param using for create friend friends
type CreateRelationsInput struct {
	RequesterEmail string
	AddresseeEmail string
}

// GetAllFriendsInput param using for get list friend of users
type GetAllFriendsInput struct {
	Email string
}

// CommonFriendsInput param using for get list common friend of two users
type CommonFriendsInput struct {
	FirstEmail  string
	SecondEmail string
}

// EmailReceiveInput param using for receive email
type EmailReceiveInput struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

// FriendListResponse response for API get list friend of users
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

// GetAllFriends get all friends of users by input
func (serv FriendService) GetAllFriends(ctx context.Context, input GetAllFriendsInput) ([]string, error) {
	user, err := serv.userRepo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	userIDs, err := serv.friendRepo.GetRelationIDsOfUser(ctx, user.ID, utils.FriendRelation)
	if err != nil {
		return nil, err
	}

	emails, err := serv.userRepo.GetListEmailByIDs(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	return emails, nil
}

// CreateFriend function implement create friend relationship
func (serv FriendService) CreateFriend(ctx context.Context, input CreateRelationsInput) error {
	requester, err := serv.userRepo.GetUserByEmail(ctx, input.RequesterEmail)
	if err != nil {
		return errors.New(ErrMessageRequesterEmailFromRequest)
	}

	// requester email not exist
	if requester.Email == "" {
		return errors.New(ErrMessageRequesterEmailNotExist)
	}

	//get addresseeId from email request
	addressee, err := serv.userRepo.GetUserByEmail(ctx, input.AddresseeEmail)
	if err != nil {
		return errors.New(ErrMessageAddresseeEmailFromRequest)
	}
	// requester email not exist
	if addressee.Email == "" {
		return errors.New(ErrMessageAddresseeEmailNotExist)
	}

	isValid, err := isValidToCreateRelationFunc(ctx, serv.friendRepo, requester.ID, addressee.ID, utils.FriendRelation)
	if !isValid || err != nil {
		return err
	}

	// insert friends two email, relationType 1 is friend
	relationFriendInput := models.UserFriend{
		RequesterID:  requester.ID,
		TargetID:     addressee.ID,
		RelationType: null.IntFrom(utils.FriendRelation),
	}
	return serv.friendRepo.CreateUserFriend(ctx, relationFriendInput)

}

// GetCommonFriends function implement get common friend
func (serv FriendService) GetCommonFriends(ctx context.Context, input CommonFriendsInput) ([]string, error) {
	var listCommonFriend []string

	//get requesterId from email request
	firstUser, err := serv.userRepo.GetUserByEmail(ctx, input.FirstEmail)
	if err != nil {
		return nil, err
	}

	//get first friend list
	firstIdsList, err := serv.friendRepo.GetRelationIDsOfUser(ctx, firstUser.ID, utils.FriendRelation)
	if err != nil {
		return nil, err
	}

	//get requesterId from email request
	secondUser, err := serv.userRepo.GetUserByEmail(ctx, input.SecondEmail)
	if err != nil {
		return nil, err
	}

	//get second friend list
	secondIdsList, err := serv.friendRepo.GetRelationIDsOfUser(ctx, secondUser.ID, utils.FriendRelation)
	if err != nil {
		return nil, err
	}

	//Intersection two list friend
	listCommonIds := utils.Intersection(firstIdsList, secondIdsList)

	listCommonFriend, err = serv.userRepo.GetListEmailByIDs(ctx, listCommonIds)
	if err != nil {
		return nil, err
	}

	return listCommonFriend, nil

}

// isRelationExist function check friends is exists
func isRelationExist(ctx context.Context, repo repository.FriendsRepo, requesterID int, addresseeID int, relationType int) (bool, error) {
	isExistRelation, err := repo.IsRelationExist(ctx, requesterID, addresseeID, relationType)
	if err != nil {
		return false, err
	}

	return isExistRelation, nil
}

// isRelationExist function check valid to create friends
func isValidToCreateRelation(ctx context.Context, repo repository.FriendsRepo, requesterID int, addresseeID int, relationType int) (bool, error) {
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
		return false, errors.New(ErrMessageUnableCreateRelation)
	}
	return isValid, nil
}

// CreateSubscription function create subscription friends
func (serv FriendService) CreateSubscription(ctx context.Context, input CreateRelationsInput) (CreateRelationsResponse, error) {
	//get requesterId from email request
	requester, err := serv.userRepo.GetUserByEmail(ctx, input.RequesterEmail)
	if err != nil {
		return CreateRelationsResponse{Success: false}, err
	}

	if requester.Email == "" {
		return CreateRelationsResponse{Success: false}, errors.New(ErrMessageRequesterEmailNotExist)
	}

	//get addresseeId from email request
	target, err := serv.userRepo.GetUserByEmail(ctx, input.AddresseeEmail)
	if err != nil {
		return CreateRelationsResponse{Success: false}, err
	}

	if target.Email == "" {
		return CreateRelationsResponse{Success: false}, errors.New(ErrMessageAddresseeEmailNotExist)
	}

	//check if users blocked or users is friend return false
	isValid, err := isValidToCreateRelationFunc(ctx, serv.friendRepo, requester.ID, target.ID, utils.Subscribe)
	//if users has friends block or friend return false
	if !isValid || err != nil {
		return CreateRelationsResponse{Success: false}, err
	}

	// insert friends friend two email, relationType 1 is friend
	relationFriendInput := models.Relation{
		RequesterID:    requester.ID,
		AddresseeID:    target.ID,
		RequesterEmail: requester.Email,
		AddresseeEmail: target.Email,
		RelationType:   utils.Subscribe,
	}
	result, err := serv.friendRepo.CreateRelation(ctx, relationFriendInput)

	return CreateRelationsResponse{Success: result}, err
}

// CreateBlock function create block friends
func (serv FriendService) CreateBlock(ctx context.Context, input CreateRelationsInput) (CreateRelationsResponse, error) {
	//get requesterId from email request
	requester, err := serv.userRepo.GetUserByEmail(ctx, input.RequesterEmail)
	if err != nil {
		return CreateRelationsResponse{Success: false}, err
	}

	//get addresseeId from email request
	target, err := serv.userRepo.GetUserByEmail(ctx, input.AddresseeEmail)
	if err != nil {
		return CreateRelationsResponse{Success: false}, err
	}

	//check if users blocked or users is friend return false
	isValid, err := isValidToCreateRelationFunc(ctx, serv.friendRepo, requester.ID, target.ID, utils.Block)
	//if users has friends block or friend return false
	if !isValid || err != nil {
		return CreateRelationsResponse{Success: false}, err
	}

	// insert friends friend two email, relationType 1 is friend
	relationFriendInput := models.Relation{
		RequesterID:    requester.ID,
		AddresseeID:    target.ID,
		RequesterEmail: requester.Email,
		AddresseeEmail: target.Email,
		RelationType:   utils.Block,
	}
	result, err := serv.friendRepo.CreateRelation(ctx, relationFriendInput)
	if err != nil {
		return CreateRelationsResponse{Success: false}, err
	}

	// if insert block is success then delete subscribe friends
	err = serv.friendRepo.DeleteRelation(ctx, requester.ID, target.ID, utils.Subscribe)

	return CreateRelationsResponse{Success: result}, err
}

// GetEmailReceive function email receive info
func (serv FriendService) GetEmailReceive(ctx context.Context, input EmailReceiveInput) (EmailReceiveResponse, error) {
	//get requesterId from email request
	requester, err := serv.userRepo.GetUserByEmail(ctx, input.Sender)

	if err != nil {
		return EmailReceiveResponse{Success: false}, errors.New(ErrMessageEmailNotExist)
	}

	//get lst id friend
	friendIDs, err := serv.friendRepo.GetRelationIDsOfUser(ctx, requester.ID, utils.FriendRelation)

	if err != nil {
		return EmailReceiveResponse{Success: false}, err
	}

	// get lst id subscription
	subscribeIDs, err := serv.friendRepo.GetRequesterIDRelation(ctx, requester.ID, utils.Subscribe)
	if err != nil {
		return EmailReceiveResponse{Success: false}, err
	}

	// get users who mentioned in text
	mentionedEmails := utils.FindEmailFromText(input.Text)
	userIDsMentionedEmails, err := serv.userRepo.GetUserIDsByEmail(ctx, mentionedEmails)

	// get requester id of block friends
	blocksID, err := serv.friendRepo.GetRequesterIDRelation(ctx, requester.ID, utils.Block)

	if err != nil {
		return EmailReceiveResponse{Success: false}, err
	}

	// users can receiver update
	receiversIDnoMentioned := append(friendIDs, subscribeIDs...)
	receiversID := utils.GetReceiverID(utils.UniqueSlice(append(receiversIDnoMentioned, userIDsMentionedEmails...)), append(blocksID, requester.ID))

	emails, err := serv.userRepo.GetListEmailByIDs(ctx, receiversID)

	if err != nil {
		return EmailReceiveResponse{Success: false}, err
	}
	return EmailReceiveResponse{
		Success:    true,
		Recipients: emails,
	}, nil
}
