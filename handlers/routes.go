package handlers

import (
	"FriendManagementAPI/database"
	"FriendManagementAPI/service"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var dbInstance database.Database

// NewHandler create router
func NewHandler(db database.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = db
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)
	router.Route("/api", users)
	return router
}

func users(router chi.Router) {
	st := service.Store{dbInstance}
	router.Post("/registration", createUser(st))
	router.Post("/friendConnection", createFriendConnection(st))
	router.Post("/friendList", getFriendList(st))
	router.Post("/commonFriend", getCommonFriendsList(st))
	router.Post("/subscribeFriend", createSubscribeFriend(st))
	router.Post("/blockFriend", createBlockFriend(st))
	router.Post("/receiveUpdates", receiveUpdate(st))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	render.Render(w, r, ErrNotFound)
}
func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, ErrMethodNotAllowed)
}
