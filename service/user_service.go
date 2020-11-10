package service

import (
	"FriendManagementAPI/common"
	"FriendManagementAPI/database"
	"FriendManagementAPI/models"
	"log"
	"strings"
)

type IUserService interface {
	CreateUser(req *models.FriendListRequest) (*models.ResultResponse, error)
	CreateFriendConnection(req *models.FriendConnectionRequest) (*models.ResultResponse, error)
	GetFriendList(req *models.FriendListRequest) (*models.FriendListResponse, error)
	GetCommonFriendsList(req *models.CommonFriendRequest) (*models.FriendListResponse, error)
	CreateSubscribeFriend(req *models.SubscriptionRequest) (*models.ResultResponse, error)
	CreateBlockFriend(req *models.BlockRequest) (*models.ResultResponse, error)
	CreateReceiveUpdate(req *models.SendUpdateEmailRequest) (*models.SendUpdateEmailResponse, error)
}
type Store struct {
	Db database.Database
}

func (st Store) CreateUser(req *models.FriendListRequest) (*models.ResultResponse, error) {
	response := &models.ResultResponse{}
	if err := st.Db.CreateUserByEmail(req.Email); err != nil {
		return response, err
	}
	response.Success = true
	return response, nil
}

func (st Store) CreateFriendConnection(req *models.FriendConnectionRequest) (*models.ResultResponse, error) {
	response := &models.ResultResponse{}
	countUser, err := st.Db.GetUserByRequest(req.Friends[0], req.Friends[1])
	if err != nil {
		return response, err
	}
	if countUser <= 1 {
		log.Printf("Email address does not exist")
		return response, err
	}
	countFriendLst, err := st.Db.GetFriendListByRequest(req)
	if err != nil {
		return response, err
	}
	if countFriendLst == 1 {
		log.Printf("Can't be add, Because they are connected as friends")
		return response, nil
	}
	countBlockLst, err := st.Db.GetBlockListByRequest(req)
	if err != nil {
		return response, err
	}
	if countBlockLst == 1 || countBlockLst == 2 {
		log.Printf("Can't be add, Because they are Blocked")
		return response, nil
	}
	// Add Friend
	if err := st.Db.CreateFriend(req); err != nil {
		if err != nil {
			response.Success = false
			return response, err
		}
	}
	response.Success = true
	return response, nil
}

func (st Store) GetFriendList(req *models.FriendListRequest) (*models.FriendListResponse, error) {
	response := &models.FriendListResponse{}
	rep, err := st.Db.GetFriendListByEmail(req.Email)
	if err != nil {
		return response, err
	}
	response.Success = true
	response.Friends = rep.Friends
	response.Count = len(rep.Friends)
	return response, nil
}

func (st Store) GetCommonFriendsList(req *models.CommonFriendRequest) (*models.FriendListResponse, error) {
	response := &models.FriendListResponse{}
	friendListEmailOne, err := st.Db.GetFriendListByEmail(req.Friends[0])
	if err != nil {
		return response, err
	}
	friendListEmailTwo, err := st.Db.GetFriendListByEmail(req.Friends[1])
	if err != nil {
		return response, err
	}
	commonFriend := []string{}
	for _, friendA := range friendListEmailOne.Friends {
		for _, friendB := range friendListEmailTwo.Friends {
			if friendA == friendB {
				commonFriend = append(commonFriend, friendA)
			}
		}
	}
	response.Success = true
	response.Friends = commonFriend
	response.Count = len(commonFriend)
	return response, nil
}

func (st Store) CreateSubscribeFriend(req *models.SubscriptionRequest) (*models.ResultResponse, error) {
	response := &models.ResultResponse{}
	countUser, err := st.Db.GetUserByRequest(req.Requestor, req.Target)
	if err != nil {
		return response, err
	}
	if countUser <= 1 {
		log.Printf("Email address does not exist")
		return response, err
	}
	countSubscribe, err := st.Db.GetTargetSubscribeByRequest(req.Requestor, req.Target)
	if err != nil {
		return response, err
	}
	if countSubscribe == 1 {
		log.Printf("%s Added %s to Subscription", req.Requestor, req.Target)
		return response, err
	}
	if err := st.Db.CreateSubscribeFriendByRequestorAndTarget(req.Requestor, req.Target); err != nil {
		return response, err
	}
	response.Success = true
	return response, nil
}

func (st Store) CreateBlockFriend(req *models.BlockRequest) (*models.ResultResponse, error) {
	response := &models.ResultResponse{}
	countUser, err := st.Db.GetUserByRequest(req.Requestor, req.Target)
	if err != nil {
		return response, err
	}
	if countUser <= 1 {
		log.Printf("Email address does not exist")
		return response, err
	}
	countBlocked, err := st.Db.GetTargetBlockedByRequest(req.Requestor, req.Target)
	if err != nil {
		return response, err
	}
	if countBlocked == 1 {
		log.Printf("Can't block, because %s blocked by %s", req.Target, req.Requestor)
		return response, err
	}
	if err := st.Db.CreateBlockFriendByRequestorAndTarget(req.Requestor, req.Target); err != nil {
		return response, err
	}
	response.Success = true
	return response, nil
}

func (st Store) CreateReceiveUpdate(req *models.SendUpdateEmailRequest) (*models.SendUpdateEmailResponse, error) {
	response := &models.SendUpdateEmailResponse{}
	countUser, err := st.Db.GetUserByEmail(req.Sender)
	if err != nil {
		return response, err
	}
	if countUser == 0 {
		log.Printf("Email address does not exist")
		return response, err
	}
	blockedLst, err := st.Db.GetAllBlockerByEmail(req.Sender)
	if err != nil {
		return response, err
	}
	allUser, err := st.Db.GetAllUser()
	if err != nil {
		return response, err
	}
	friendLst, err := st.Db.GetFriendListByEmail(req.Sender)
	if err != nil {
		return response, err
	}
	subscriber, err := st.Db.GetAllSubscriberByEmail(req.Sender)
	if err != nil {
		return response, err
	}
	recipient := []string{}
	for _, user := range allUser {
		boolIsBlocked := common.CheckIsExist(blockedLst.Blocked, user.Email)
		if !boolIsBlocked {
			boolIsFriend := common.CheckIsExist(friendLst.Friends, user.Email)
			boolIsSubscribe := common.CheckIsExist(subscriber.Subscription, user.Email)
			boolIsMention := strings.Contains(req.Text, user.Email)
			if boolIsFriend || boolIsSubscribe || boolIsMention {
				recipient = append(recipient, user.Email)
			}
		}
	}
	response.Success = true
	response.Recipients = recipient
	return response, nil
}
