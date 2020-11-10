package models

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Email        string   `json:"email"`
	Friends      []string `json:"friends"`
	Subscription []string `json:"subscription"`
	Blocked      []string `json:"blocked"`
}

type ListUser struct {
	Users []User `json:"users"`
}

type ResultResponse struct {
	Success bool `json:"success"`
}

type FriendConnectionRequest struct {
	Friends []string `json:"friends"`
}

type FriendListRequest struct {
	Email string `json:"email"`
}

type CommonFriendRequest struct {
	Friends []string `json:"friends"`
}

type FriendListResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}

type SubscriptionRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

type BlockRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

type SendUpdateEmailRequest struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

type SendUpdateEmailResponse struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients"`
}

func (friend *FriendConnectionRequest) Bind(r *http.Request) error {
	userEmailOne := friend.Friends[0]
	userEmailTwo := friend.Friends[1]
	if userEmailOne == "" || userEmailTwo == "" {
		return fmt.Errorf("email is a required field")
	}
	if userEmailOne == userEmailTwo {
		log.Print("can't friend connect myself")
		return fmt.Errorf("can't connect myself")
	}
	return nil
}

func (email *FriendListRequest) Bind(r *http.Request) error {
	if email.Email == "" {
		log.Print("email is a required field")
		return fmt.Errorf("email is a required field")
	}
	return nil
}

func (req *SubscriptionRequest) Bind(r *http.Request) error {
	requestor := req.Requestor
	target := req.Target
	if requestor == "" || target == "" {
		return fmt.Errorf("email is a required field")
	}
	if requestor == target {
		log.Print("can't subscribe with myself")
		return fmt.Errorf("can't subscribe myself")
	}
	return nil
}

func (req *BlockRequest) Bind(r *http.Request) error {
	requestor := req.Requestor
	target := req.Target
	if requestor == "" || target == "" {
		return fmt.Errorf("email is a required field")
	}
	if requestor == target {
		log.Print("can't block with myself")
		return fmt.Errorf("can't block with myself")
	}
	return nil
}

func (req *SendUpdateEmailRequest) Bind(r *http.Request) error {
	if req.Sender == "" || req.Text == "" {
		return fmt.Errorf("email and content is a required field")
	}
	return nil
}

func (friend *CommonFriendRequest) Bind(r *http.Request) error {
	userEmailOne := friend.Friends[0]
	userEmailTwo := friend.Friends[1]
	if userEmailOne == "" || userEmailTwo == "" {
		return fmt.Errorf("email is a required field")
	}
	if userEmailOne == userEmailTwo {
		log.Print("can't friend connect myself")
		return fmt.Errorf("can't connect myself")
	}
	return nil
}
