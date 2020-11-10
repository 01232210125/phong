package mocks

import (
	"FriendManagementAPI/models"

	"github.com/stretchr/testify/mock"
)

type StoreMock struct {
	mock.Mock
}

// func (serMock *ServiceMock) createFriendConnection
func (st StoreMock) CreateFriendConnection(req *models.FriendConnectionRequest) (*models.ResultResponse, error) {
	returnVals := st.Called(req)
	r0 := returnVals.Get(0).(models.ResultResponse)
	var r1 error
	if returnVals.Get(1) != nil {
		r1 = returnVals.Get(1).(error)
	}
	return &r0, r1
}

func (st StoreMock) GetFriendList(req *models.FriendListRequest) (*models.FriendListResponse, error) {
	returnVals := st.Called(req)
	r0 := returnVals.Get(0).(models.FriendListResponse)
	var r1 error
	if returnVals.Get(1) != nil {
		r1 = returnVals.Get(1).(error)
	}
	return &r0, r1
}

func (st StoreMock) GetCommonFriendsList(req *models.CommonFriendRequest) (*models.FriendListResponse, error) {
	returnVals := st.Called(req)
	r0 := returnVals.Get(0).(models.FriendListResponse)
	var r1 error
	if returnVals.Get(1) != nil {
		r1 = returnVals.Get(1).(error)
	}
	return &r0, r1
}

func (st StoreMock) CreateSubscribeFriend(req *models.SubscriptionRequest) (*models.ResultResponse, error) {
	returnVals := st.Called(req)
	r0 := returnVals.Get(0).(models.ResultResponse)
	var r1 error
	if returnVals.Get(1) != nil {
		r1 = returnVals.Get(1).(error)
	}
	return &r0, r1
}

func (st StoreMock) CreateBlockFriend(req *models.BlockRequest) (*models.ResultResponse, error) {
	returnVals := st.Called(req)
	r0 := returnVals.Get(0).(models.ResultResponse)
	var r1 error
	if returnVals.Get(1) != nil {
		r1 = returnVals.Get(1).(error)
	}
	return &r0, r1
}

func (st StoreMock) CreateReceiveUpdate(req *models.SendUpdateEmailRequest) (*models.SendUpdateEmailResponse, error) {
	returnVals := st.Called(req)
	r0 := returnVals.Get(0).(models.SendUpdateEmailResponse)
	var r1 error
	if returnVals.Get(1) != nil {
		r1 = returnVals.Get(1).(error)
	}
	return &r0, r1
}
