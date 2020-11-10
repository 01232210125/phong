package handlers

import (
	"FriendManagementAPI/models"
	"FriendManagementAPI/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
)

func createUser(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.FriendListRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrBadRequest)
			return
		}
		response, err := service.CreateUser(req)
		if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
		// send Result response
		json.NewEncoder(w).Encode(response)
	}
}

func createFriendConnection(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.FriendConnectionRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrBadRequest)
			return
		}
		response, err := service.CreateFriendConnection(req)
		if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
		// send Result response
		json.NewEncoder(w).Encode(response)
	}
}

func getFriendList(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.FriendListRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrBadRequest)
			return
		}
		response, err := service.GetFriendList(req)
		if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
		// send Result response
		json.NewEncoder(w).Encode(response)
	}
}

func getCommonFriendsList(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.CommonFriendRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrBadRequest)
			return
		}
		response, err := service.GetCommonFriendsList(req)
		if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
		// send Result response
		json.NewEncoder(w).Encode(response)
	}
}

func createSubscribeFriend(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.SubscriptionRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrBadRequest)
			return
		}
		response, err := service.CreateSubscribeFriend(req)
		if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
		// send Result response
		json.NewEncoder(w).Encode(response)
	}
}

func createBlockFriend(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.BlockRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrBadRequest)
			return
		}
		response, err := service.CreateBlockFriend(req)
		if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
		// send Result response
		json.NewEncoder(w).Encode(response)
	}
}

func receiveUpdate(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.SendUpdateEmailRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrBadRequest)
			return
		}
		response, err := service.CreateReceiveUpdate(req)
		if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
		// send Result response
		json.NewEncoder(w).Encode(response)
	}
}
