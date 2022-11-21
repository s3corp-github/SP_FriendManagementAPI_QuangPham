package friends

import (
	"context"
	"regexp"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/models"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/utils"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/repository"
	"github.com/volatiletech/null/v8"
)

var (
	isRelationExistFunc         = isRelationExist
	isValidToCreateRelationFunc = isValidToCreateRelation
)

// CreateRelationsInput param using for create friends
type CreateRelationsInput struct {
	RequesterEmail string
	TargetEmail    string
}

// GetAllFriendsInput param using for get list friend of users
type GetAllFriendsInput struct {
	Email string
}

// CommonFriendsInput param using for get list common friend of two users
type CommonFriendsInput struct {
	RequesterEmail string
	TargetEmail    string
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

// GetFriends get all friends of users by input
func (serv FriendService) GetFriends(ctx context.Context, input GetAllFriendsInput) ([]string, error) {
	user, err := serv.userRepo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	userIDs, err := serv.friendRepo.GetRelationIDs(ctx, user.ID, utils.FriendRelation)
	if err != nil {
		return nil, err
	}

	return serv.userRepo.GetEmailsByIDs(ctx, userIDs)
}

// CreateFriend function implement create friend relationship
func (serv FriendService) CreateFriend(ctx context.Context, input CreateRelationsInput) error {
	requesterUser, err := serv.userRepo.GetUserByEmail(ctx, input.RequesterEmail)
	if err != nil {
		return ErrRequestEmailInvalid
	}

	targetUser, err := serv.userRepo.GetUserByEmail(ctx, input.TargetEmail)
	if err != nil {
		return ErrTargetEmailInvalid
	}

	isValid, err := isValidToCreateRelationFunc(ctx, serv.friendRepo, requesterUser.ID, targetUser.ID, utils.FriendRelation)
	if !isValid {
		return ErrRelationIsExists
	}
	if err != nil {
		return err
	}

	relationFriendInput := models.UserFriend{
		RequesterID:  requesterUser.ID,
		TargetID:     targetUser.ID,
		RelationType: null.IntFrom(utils.FriendRelation),
	}

	return serv.friendRepo.CreateUserFriend(ctx, relationFriendInput)

}

// GetCommonFriends function implement get common friend
func (serv FriendService) GetCommonFriends(ctx context.Context, input CommonFriendsInput) ([]string, error) {
	firstUser, err := serv.userRepo.GetUserByEmail(ctx, input.RequesterEmail)
	if err != nil {
		return nil, err
	}

	firstIdsList, err := serv.friendRepo.GetRelationIDs(ctx, firstUser.ID, utils.FriendRelation)
	if err != nil {
		return nil, err
	}

	secondUser, err := serv.userRepo.GetUserByEmail(ctx, input.TargetEmail)
	if err != nil {
		return nil, err
	}

	secondIdsList, err := serv.friendRepo.GetRelationIDs(ctx, secondUser.ID, utils.FriendRelation)
	if err != nil {
		return nil, err
	}

	userIds := getCommonUserIDs(firstIdsList, secondIdsList)

	return serv.userRepo.GetEmailsByIDs(ctx, userIds)
}

// isRelationExist function check friends is exists
func isRelationExist(ctx context.Context, repo repository.FriendsRepo, requesterID int, targetID int, relationType int) (bool, error) {
	return repo.IsRelationExist(ctx, requesterID, targetID, relationType)
}

// isRelationExist function check valid to create friends
func isValidToCreateRelation(ctx context.Context, repo repository.FriendsRepo, requesterID int, targetID int, relationType int) (bool, error) {
	isExistRelation, err := isRelationExistFunc(ctx, repo, requesterID, targetID, relationType)
	if err != nil {
		return false, err
	}

	if isExistRelation {
		return false, ErrRelationIsExists
	}

	var isValid bool

	switch relationType {
	case utils.FriendRelation:
		isRequesterIDBlock, err := isRelationExistFunc(ctx, repo, requesterID, targetID, utils.Blocked)
		isTargetIDBlock, err := isRelationExistFunc(ctx, repo, requesterID, targetID, utils.Blocked)
		if err != nil {
			return false, err
		}

		if !isRequesterIDBlock && !isTargetIDBlock {
			isValid = true
		}
	case utils.Subscribe:
		isValid = true
	case utils.Blocked:
		isValid = true
	}

	return isValid, nil
}

// CreateSubscription function create subscription friends
func (serv FriendService) CreateSubscription(ctx context.Context, input CreateRelationsInput) error {
	requester, err := serv.userRepo.GetUserByEmail(ctx, input.RequesterEmail)
	if err != nil {
		return err
	}

	target, err := serv.userRepo.GetUserByEmail(ctx, input.TargetEmail)
	if err != nil {
		return err
	}

	isValid, err := isValidToCreateRelationFunc(ctx, serv.friendRepo, requester.ID, target.ID, utils.Subscribe)
	if err != nil {
		return err
	}

	if !isValid {
		return ErrTwoEmailInvalidCreateSub
	}

	relationFriendInput := models.UserFriend{
		RequesterID:  requester.ID,
		TargetID:     target.ID,
		RelationType: null.IntFrom(utils.Subscribe),
	}

	return serv.friendRepo.CreateUserFriend(ctx, relationFriendInput)
}

// CreateBlock function create block friends
func (serv FriendService) CreateBlock(ctx context.Context, input CreateRelationsInput) error {
	requesterUser, err := serv.userRepo.GetUserByEmail(ctx, input.RequesterEmail)
	if err != nil {
		return err
	}

	targetUser, err := serv.userRepo.GetUserByEmail(ctx, input.TargetEmail)
	if err != nil {
		return err
	}

	isValid, err := isValidToCreateRelationFunc(ctx, serv.friendRepo, requesterUser.ID, targetUser.ID, utils.Blocked)
	if err != nil {
		return err
	}

	if !isValid {
		return ErrTwoEmailInvalidCreateBlock
	}

	relationFriendInput := models.UserFriend{
		RequesterID:  requesterUser.ID,
		TargetID:     targetUser.ID,
		RelationType: null.IntFrom(utils.Blocked),
	}

	if err = serv.friendRepo.CreateUserFriend(ctx, relationFriendInput); err != nil {
		return ErrTwoEmailInvalidCreateBlock
	}

	if err = serv.friendRepo.DeleteRelation(ctx, requesterUser.ID, targetUser.ID, utils.Subscribe); err != nil {
		return ErrTwoEmailInvalidCreateBlock
	}

	return nil
}

// GetEmailReceive function email receive info
func (serv FriendService) GetEmailReceive(ctx context.Context, input EmailReceiveInput) ([]string, error) {
	requester, err := serv.userRepo.GetUserByEmail(ctx, input.Sender)
	if err != nil {
		return nil, err
	}

	friendIDs, err := serv.friendRepo.GetRelationIDs(ctx, requester.ID, utils.FriendRelation)
	if err != nil {
		return nil, err
	}

	subscribeIDs, err := serv.friendRepo.GetRequesterIDFriends(ctx, requester.ID, utils.Subscribe)
	if err != nil {
		return nil, err
	}

	mentionedEmails := findEmailFromText(input.Text)
	userIDsMentionedEmails, err := serv.userRepo.GetUserIDsByEmail(ctx, mentionedEmails)
	if err != nil {
		return nil, err
	}

	blocksID, err := serv.friendRepo.GetRequesterIDFriends(ctx, requester.ID, utils.Blocked)
	if err != nil {
		return nil, err
	}

	receiversIDnoMentioned := append(friendIDs, subscribeIDs...)
	receiversID := getReceiverID(uniqueSlice(append(receiversIDnoMentioned, userIDsMentionedEmails...)), append(blocksID, requester.ID))

	return serv.userRepo.GetEmailsByIDs(ctx, receiversID)
}

// getCommonUserIDs find the same elements of two array
func getCommonUserIDs(userID1, userID2 []int) []int {
	userIDMap := make(map[int]bool)
	for _, item := range userID1 {
		userIDMap[item] = true
	}

	var commonUserIDs []int
	for _, item := range userID2 {
		if _, ok := userIDMap[item]; ok {
			commonUserIDs = append(commonUserIDs, item)
		}
	}
	return commonUserIDs
}

// getReceiverID get slice of receiver id
func getReceiverID(userIDs1, userIDs2 []int) []int {

	sameElements := getCommonUserIDs(userIDs1, userIDs2)

	for _, v := range sameElements {
		userIDs1 = removeIndex(userIDs1, indexOf(v, userIDs1))
	}

	return userIDs1
}

// removeIndex remove index in slice
func removeIndex(userIDs []int, index int) []int {
	return append(userIDs[:index], userIDs[index+1:]...)
}

// indexOf get index of value in slice
func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

// uniqueSlice remove duplicate element in slice
func uniqueSlice(intSlice []int) []int {
	keys := make(map[int]bool)
	var list []int

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// findEmailFromText return email mentioned in text
func findEmailFromText(text string) []string {

	regex := regexp.MustCompile(utils.EmailValidationRegex)

	emailChain := regex.FindAllString(text, -1)

	emails := make([]string, len(emailChain))

	for index, emailCharacter := range emailChain {
		emails[index] = emailCharacter
	}
	return emails
}
