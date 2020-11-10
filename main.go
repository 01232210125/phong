package main

import (
	"FriendManagementAPI/database"
	"FriendManagementAPI/handlers"
	"log"
	"net/http"
)

func main() {
	database, err := database.Initialize()
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer database.Conn.Close()

	httpHandler := handlers.NewHandler(database)
	log.Println("Server started on: http://localhost:8080")
	http.ListenAndServe(":8080", httpHandler)
}
