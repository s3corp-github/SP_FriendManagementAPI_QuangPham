package friends

import (
	"context"
	"regexp"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/models"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/utils"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/repository"
)

var (
	emailValidationRegex = "[_A-Za-z0-9-\\+]+(\\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\\.[A-Za-z0-9]+)*(\\.[A-Za-z]{2,})"
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
		return err
	}

	targetUser, err := serv.userRepo.GetUserByEmail(ctx, input.TargetEmail)
	if err != nil {
		return err
	}

	checkExists, err := checkUserFriendRelation(ctx, serv.friendRepo, requesterUser.ID, targetUser.ID)
	if err != nil {
		return err
	}
	if !checkExists {
		return ErrRelationIsExists
	}

	relationFriendInput := models.UserFriend{
		RequesterID:  requesterUser.ID,
		TargetID:     targetUser.ID,
		RelationType: utils.FriendRelation,
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

// checkUserFriendRelation function check valid to create friends
func checkUserFriendRelation(ctx context.Context, repo repository.FriendsRepo, requesterID int, targetID int) (bool, error) {
	checkExistFriend, err := repo.CheckFriendRelationExist(ctx, requesterID, targetID)
	if err != nil {
		return false, err
	}

	if checkExistFriend {
		return false, ErrRelationIsExists
	}

	checkRequesterIDBlock, err := repo.CheckBlockRelationExist(ctx, requesterID, targetID)
	if err != nil {
		return false, err
	}

	checkTargetIDBlock, err := repo.CheckBlockRelationExist(ctx, targetID, requesterID)
	if err != nil {
		return false, err
	}

	if checkRequesterIDBlock || checkTargetIDBlock {
		return false, nil
	}

	return true, nil
}

// CreateSubscription function create subscription friends
func (serv FriendService) CreateSubscription(ctx context.Context, input CreateRelationsInput) error {
	requesterUser, err := serv.userRepo.GetUserByEmail(ctx, input.RequesterEmail)
	if err != nil {
		return err
	}

	targetUser, err := serv.userRepo.GetUserByEmail(ctx, input.TargetEmail)
	if err != nil {
		return err
	}

	checkExists, err := serv.friendRepo.CheckSubscriptionRelationExist(ctx, requesterUser.ID, targetUser.ID)
	if err != nil {
		return err
	}
	if checkExists {
		return ErrTwoEmailInvalidCreateSub
	}

	relationFriendInput := models.UserFriend{
		RequesterID:  requesterUser.ID,
		TargetID:     targetUser.ID,
		RelationType: utils.Subscribe,
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

	checkExists, err := serv.friendRepo.CheckBlockRelationExist(ctx, requesterUser.ID, targetUser.ID)
	if err != nil {
		return err
	}
	if checkExists {
		return ErrTwoEmailInvalidCreateBlock
	}

	relationFriendInput := models.UserFriend{
		RequesterID:  requesterUser.ID,
		TargetID:     targetUser.ID,
		RelationType: utils.Blocked,
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
	regex := regexp.MustCompile(emailValidationRegex)
	emailChain := regex.FindAllString(text, -1)
	emails := make([]string, len(emailChain))
	for index, emailCharacter := range emailChain {
		emails[index] = emailCharacter
	}
	return emails
}
