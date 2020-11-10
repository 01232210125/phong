package service

import (
	"FriendManagementAPI/database"
	"FriendManagementAPI/models"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateReceiveUpdate(t *testing.T) {
	db := createConnectionForTest()
	defer db.Conn.Close()
	testCases := []struct {
		name           string
		sender         string
		text           string
		expectedResult *models.SendUpdateEmailResponse
	}{
		{
			name:   "retieve failed by email address does not exist",
			sender: "john1@example",
			text:   "Hello World! kate@example",
			expectedResult: &models.SendUpdateEmailResponse{
				Success: false,
			},
		},
		{
			name:   "retieve empty by email address blocked",
			sender: "tomy@example",
			text:   "Hello World!",
			expectedResult: &models.SendUpdateEmailResponse{
				Success:    true,
				Recipients: []string{},
			},
		},
		{
			name:   "Success",
			sender: "john@example",s
			text:   "Hello World! kate@example",
			expectedResult: &models.SendUpdateEmailResponse{
				Success:    true,
				Recipients: []string{"lisa@example", "kate@example", "tom@example"},
			},
		},
	}
	store := Store{db}
	err := insertReceiveUpdate(db.Conn)
	require.NoError(t, err)
	for _, tt := range testCases {
		req := &models.SendUpdateEmailRequest{
			Sender: tt.sender,
			Text:   tt.text,
		}
		response, err := store.CreateReceiveUpdate(req)
		require.NoError(t, err)
		require.Equal(t, tt.expectedResult, response)
	}
}
func insertReceiveUpdate(db *sql.DB) error {
	query :=
		`
		truncate block;
		truncate friend;
		truncate subscription;
		truncate userprofile cascade;
		insert into userprofile (email) values ('john@example');
		insert into userprofile (email) values ('lisa@example');
		insert into userprofile (email) values ('kate@example');
		insert into userprofile (email) values ('tom@example');
		insert into userprofile (email) values ('jerry@example');
		insert into userprofile (email) values ('tomy@example');
		insert into userprofile (email) values ('adam@example');
		insert into friend (emailuserone, emailusertwo) values ('john@example','lisa@example');
		insert into friend (emailuserone, emailusertwo) values ('tomy@example','adam@example');
		insert into subscription (requestor, target) values ('tom@example','john@example');
		insert into block (requestor, target) values ('jerry@example','john@example');
		insert into block (requestor, target) values ('adam@example','tomy@example');
		`
	_, err := db.Exec(query)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}

func TestCreateBlockFriend(t *testing.T) {
	db := createConnectionForTest()
	defer db.Conn.Close()
	testCases := []struct {
		name           string
		requestor      string
		target         string
		expectedResult *models.ResultResponse
	}{
		{
			name:      "Success",
			requestor: "john@example",
			target:    "lisa@example",
			expectedResult: &models.ResultResponse{
				Success: true,
			},
		},
		{
			name:      "retieve failed by email address does not exist",
			requestor: "tom@example",
			target:    "andy@example",
			expectedResult: &models.ResultResponse{
				Success: false,
			},
		},
		{
			name:      "retieve failed by email blocked",
			requestor: "tom@example",
			target:    "jerry@example",
			expectedResult: &models.ResultResponse{
				Success: false,
			},
		},
	}
	store := Store{db}
	err := insertBlock(db.Conn)
	require.NoError(t, err)
	for _, tt := range testCases {
		req := &models.BlockRequest{
			Requestor: tt.requestor,
			Target:    tt.target,
		}
		response, err := store.CreateBlockFriend(req)
		require.NoError(t, err)
		require.Equal(t, tt.expectedResult, response)
	}
}

func insertBlock(db *sql.DB) error {
	query :=
		`
		truncate block;
		truncate userprofile cascade;
		insert into userprofile (email) values ('john@example');
		insert into userprofile (email) values ('lisa@example');
		insert into userprofile (email) values ('tom@example');
		insert into userprofile (email) values ('jerry@example');
		insert into block (requestor, target) values ('tom@example','jerry@example');
		`
	_, err := db.Exec(query)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}

func TestCreateSubscribeFriend(t *testing.T) {
	db := createConnectionForTest()
	defer db.Conn.Close()
	testCases := []struct {
		name           string
		requestor      string
		target         string
		expectedResult *models.ResultResponse
	}{
		{
			name:      "Success",
			requestor: "andy@example",
			target:    "john@example",
			expectedResult: &models.ResultResponse{
				Success: true,
			},
		},
		{
			name:      "retieve failed by email address does not exist",
			requestor: "lisa@example",
			target:    "john@example",
			expectedResult: &models.ResultResponse{
				Success: false,
			},
		},
		{
			name:      "retieve failed by target email address added to subscription",
			requestor: "tom@example",
			target:    "jerry@example",
			expectedResult: &models.ResultResponse{
				Success: false,
			},
		},
	}
	store := Store{db}
	err := insertSubscribeFriend(db.Conn)
	require.NoError(t, err)
	for _, tt := range testCases {
		req := &models.SubscriptionRequest{
			Requestor: tt.requestor,
			Target:    tt.target,
		}
		response, err := store.CreateSubscribeFriend(req)
		require.NoError(t, err)
		require.Equal(t, tt.expectedResult, response)
	}
}
func insertSubscribeFriend(db *sql.DB) error {
	query :=
		`
		truncate userprofile cascade;
		truncate subscription;
		insert into userprofile (email) values ('andy@example');
		insert into userprofile (email) values ('john@example');
		insert into userprofile (email) values ('tom@example');
		insert into userprofile (email) values ('jerry@example');
		insert into subscription (requestor, target) values ('tom@example','jerry@example');
		`
	_, err := db.Exec(query)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}

func TestGetCommonFriendsList(t *testing.T) {
	db := createConnectionForTest()
	defer db.Conn.Close()
	testCases := []struct {
		name           string
		friend         []string
		expectedResult *models.FriendListResponse
	}{
		{
			name:   "Success",
			friend: []string{"andy@example", "john@example"},
			expectedResult: &models.FriendListResponse{
				Success: true,
				Friends: []string{"ana@example", "tom@example", "jerry@example"},
				Count:   3,
			},
		},
		{
			name:   "Empty",
			friend: []string{"lisa@example", "john@example"},
			expectedResult: &models.FriendListResponse{
				Success: true,
				Friends: []string{},
				Count:   0,
			},
		},
	}
	store := Store{db}
	err := insertCommonFriend(db.Conn)
	require.NoError(t, err)
	for _, tt := range testCases {
		req := &models.CommonFriendRequest{
			Friends: tt.friend,
		}
		response, err := store.GetCommonFriendsList(req)
		require.NoError(t, err)
		require.Equal(t, tt.expectedResult, response)
	}
}

func insertCommonFriend(db *sql.DB) error {
	query :=
		`
		truncate friend;
		truncate userprofile cascade;
		insert into userprofile (email) values ('andy@example');
		insert into userprofile (email) values ('john@example');
		insert into userprofile (email) values ('ana@example');
		insert into userprofile (email) values ('tom@example');
		insert into userprofile (email) values ('jerry@example');
		insert into userprofile (email) values ('lisa@example');
		insert into friend (emailuserone, emailusertwo) values ('andy@example','ana@example');
		insert into friend (emailuserone, emailusertwo) values ('john@example','ana@example');
		insert into friend (emailuserone, emailusertwo) values ('tom@example','andy@example');
		insert into friend (emailuserone, emailusertwo) values ('tom@example','john@example');
		insert into friend (emailuserone, emailusertwo) values ('jerry@example','andy@example');
		insert into friend (emailuserone, emailusertwo) values ('john@example','jerry@example');
		insert into friend (emailuserone, emailusertwo) values ('lisa@example','john@example');
		`
	_, err := db.Exec(query)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}
func TestCreateFriendConnection(t *testing.T) {
	db := createConnectionForTest()
	defer db.Conn.Close()
	testCases := []struct {
		name           string
		friend         []string
		expectedResult *models.ResultResponse
	}{
		{
			name:   "Success",
			friend: []string{"andy@example", "john@example"},
			expectedResult: &models.ResultResponse{
				Success: true,
			},
		},
		{
			name:   "Success",
			friend: []string{"john@example", "ana@example"},
			expectedResult: &models.ResultResponse{
				Success: true,
			},
		},
		{
			name:   "retieve failed by email address does not exist",
			friend: []string{"andy@example", "lisa@example"},
			expectedResult: &models.ResultResponse{
				Success: false,
			},
		},
		{
			name:   "retieve failed by they are connected as friends",
			friend: []string{"andy@example", "john@example"},
			expectedResult: &models.ResultResponse{
				Success: false,
			},
		},
		{
			name:   "retieve failed by they are Blocked",
			friend: []string{"tom@example", "jerry@example"},
			expectedResult: &models.ResultResponse{
				Success: false,
			},
		},
		{
			name:   "retieve failed by they are Blocked",
			friend: []string{"jerry@example", "tom@example"},
			expectedResult: &models.ResultResponse{
				Success: false,
			},
		},
	}
	store := Store{db}
	err := insertConnectFriend(db.Conn)
	require.NoError(t, err)
	for _, tt := range testCases {
		req := &models.FriendConnectionRequest{
			Friends: tt.friend,
		}
		response, err := store.CreateFriendConnection(req)
		require.NoError(t, err)
		require.Equal(t, tt.expectedResult, response)
	}
}

func insertConnectFriend(db *sql.DB) error {
	query :=
		`
		truncate friend;
		truncate userprofile cascade;
		insert into userprofile (email) values ('andy@example');
		insert into userprofile (email) values ('john@example');
		insert into userprofile (email) values ('ana@example');
		insert into userprofile (email) values ('tom@example');
		insert into userprofile (email) values ('jerry@example');
		insert into block (requestor, target) values ('tom@example','jerry@example');
		`
	_, err := db.Exec(query)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}
func TestGetFriendList(t *testing.T) {
	db := createConnectionForTest()
	defer db.Conn.Close()
	testCases := []struct {
		name           string
		email          string
		expectedResult *models.FriendListResponse
	}{
		{
			name:  "success",
			email: "phong",
			expectedResult: &models.FriendListResponse{
				Success: true,
				Friends: []string{"hien", "thinh", "dat"},
				Count:   3,
			},
		},
		{
			name:  "empty",
			email: "phong123",
			expectedResult: &models.FriendListResponse{
				Success: true,
				Count:   0,
			},
		},
	}
	store := Store{db}
	err := insertFriend(db.Conn)
	require.NoError(t, err)
	for _, tt := range testCases {
		req := &models.FriendListRequest{
			Email: tt.email,
		}
		response, err := store.GetFriendList(req)
		require.NoError(t, err)
		require.Equal(t, tt.expectedResult, response)
	}
}

func insertFriend(db *sql.DB) error {
	query :=
		`
		truncate friend;
		truncate userprofile cascade;
		insert into userprofile (email) values ('join');
		insert into userprofile (email) values ('henry');
		insert into userprofile (email) values ('phong');
		insert into userprofile (email) values ('hien');
		insert into userprofile (email) values ('thinh');
		insert into userprofile (email) values ('dat');
		insert into friend (emailuserone, emailusertwo) values ('phong','hien');
		insert into friend (emailuserone, emailusertwo) values ('phong','thinh');
		insert into friend (emailuserone, emailusertwo) values ('phong','dat');
		insert into friend (emailuserone, emailusertwo) values ('join','henry');
		`
	_, err := db.Exec(query)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}

// create connection with postgres db
func createConnectionForTest() database.Database {
	db := database.Database{}
	// Open the connection
	conn, err := sql.Open("postgres", "postgres://postgres:Hien123456@localhost:5432/FriendMangementDB?sslmode=disable")

	if err != nil {
		panic(err)
	}
	db.Conn = conn
	// check the connection
	err = db.Conn.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}
