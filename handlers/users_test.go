package handlers

import (
	"FriendManagementAPI/mocks"
	"FriendManagementAPI/models"
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateFriendConnection(t *testing.T) {
	var jsonStr = []byte(`{"friends":["andy@example","john@example.com"]}`)
	testCases := []struct {
		name         string
		bodyRequest  *bytes.Buffer
		expectedCode int
		expectedBody string
		mockresponse models.ResultResponse
		mockError    error
	}{
		{
			name:        "retieve success",
			bodyRequest: bytes.NewBuffer(jsonStr),
			mockresponse: models.ResultResponse{
				Success: true,
			},
			expectedCode: 200,
			expectedBody: "{\"success\":true}\n",
		},
		{
			name:        "retieve failed by CreateFriendConnection error",
			bodyRequest: bytes.NewBuffer(jsonStr),
			mockresponse: models.ResultResponse{
				Success: false,
			},
			mockError:    errors.New("db err"),
			expectedCode: 500,
			expectedBody: "{\"statusCode\":500,\"message\":\"db err\"}\n",
		},
		{
			name:         "retieve failed by incorrect input",
			bodyRequest:  bytes.NewBuffer([]byte("")),
			expectedCode: 400,
			expectedBody: "{\"statusCode\":400,\"message\":\"Bad request\"}\n",
		},
	}
	for _, tt := range testCases {
		req, err := http.NewRequest(http.MethodPost, "/api/friendConnection", tt.bodyRequest)
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")
		require.NoError(t, err)
		router := chi.NewRouter()
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, router))
		rr := httptest.NewRecorder()
		serviceMock := new(mocks.StoreMock)
		serviceMock.On("CreateFriendConnection", mock.Anything, mock.Anything).Return(tt.mockresponse, tt.mockError)
		handler := createFriendConnection(serviceMock)
		handler.ServeHTTP(rr, req)
		require.Equal(t, tt.expectedCode, rr.Code)
		require.Equal(t, tt.expectedBody, rr.Body.String())
	}
}

func TestGetFriendList(t *testing.T) {
	var jsonStr = []byte(`{"email":"andy@example"}`)
	testCases := []struct {
		name         string
		bodyRequest  *bytes.Buffer
		expectedCode int
		expectedBody string
		mockresponse models.FriendListResponse
		mockError    error
	}{
		{
			name:         "retieve success",
			bodyRequest:  bytes.NewBuffer(jsonStr),
			expectedCode: 200,
			expectedBody: "{\"success\":true,\"friends\":[\"john@example\",\"tom@example\"],\"count\":1}\n",
			mockresponse: models.FriendListResponse{
				Success: true,
				Friends: []string{"john@example", "tom@example"},
				Count:   1,
			},
		},
		{
			name:        "retieve failed by CreateFriendConnection error",
			bodyRequest: bytes.NewBuffer(jsonStr),
			mockresponse: models.FriendListResponse{
				Success: false,
			},
			mockError:    errors.New("db err"),
			expectedCode: 500,
			expectedBody: "{\"statusCode\":500,\"message\":\"db err\"}\n",
		},
		{
			name:         "retieve failed by incorrect input",
			bodyRequest:  bytes.NewBuffer([]byte("")),
			expectedCode: 400,
			expectedBody: "{\"statusCode\":400,\"message\":\"Bad request\"}\n",
		},
	}
	for _, tt := range testCases {
		req, err := http.NewRequest(http.MethodPost, "/api/friendList", tt.bodyRequest)
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")
		require.NoError(t, err)
		router := chi.NewRouter()
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, router))
		rr := httptest.NewRecorder()
		serviceMock := new(mocks.StoreMock)
		serviceMock.On("GetFriendList", mock.Anything, mock.Anything).Return(tt.mockresponse, tt.mockError)
		handler := getFriendList(serviceMock)
		handler.ServeHTTP(rr, req)
		require.Equal(t, tt.expectedCode, rr.Code)
		require.Equal(t, tt.expectedBody, rr.Body.String())
	}
}

func TestGetCommonFriendsList(t *testing.T) {
	var jsonStr = []byte(`{"friends":["andy@example","john@example.com"]}`)
	testCases := []struct {
		name         string
		bodyRequest  *bytes.Buffer
		expectedCode int
		expectedBody string
		mockresponse models.FriendListResponse
		mockError    error
	}{
		{
			name:         "retieve success",
			bodyRequest:  bytes.NewBuffer(jsonStr),
			expectedCode: 200,
			expectedBody: "{\"success\":true,\"friends\":[\"lisa@example\",\"tom@example\"],\"count\":1}\n",
			mockresponse: models.FriendListResponse{
				Success: true,
				Friends: []string{"lisa@example", "tom@example"},
				Count:   1,
			},
		},
		{
			name:        "retieve failed by GetCommonFriendsList error",
			bodyRequest: bytes.NewBuffer(jsonStr),
			mockresponse: models.FriendListResponse{
				Success: false,
			},
			mockError:    errors.New("db err"),
			expectedCode: 500,
			expectedBody: "{\"statusCode\":500,\"message\":\"db err\"}\n",
		},
		{
			name:         "retieve failed by incorrect input",
			bodyRequest:  bytes.NewBuffer([]byte("")),
			expectedCode: 400,
			expectedBody: "{\"statusCode\":400,\"message\":\"Bad request\"}\n",
		},
	}

	for _, tt := range testCases {
		req, err := http.NewRequest(http.MethodPost, "/api/commonFriend", tt.bodyRequest)
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")
		require.NoError(t, err)
		router := chi.NewRouter()
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, router))
		rr := httptest.NewRecorder()
		serviceMock := new(mocks.StoreMock)
		serviceMock.On("GetCommonFriendsList", mock.Anything, mock.Anything).Return(tt.mockresponse, tt.mockError)
		handler := getCommonFriendsList(serviceMock)
		handler.ServeHTTP(rr, req)
		require.Equal(t, tt.expectedCode, rr.Code)
		require.Equal(t, tt.expectedBody, rr.Body.String())
	}
}

func TestCreateSubscribeFriend(t *testing.T) {
	var jsonStr = []byte(`{"requestor":"andy@example","target":"john@example"}`)
	testCases := []struct {
		name         string
		bodyRequest  *bytes.Buffer
		expectedCode int
		expectedBody string
		mockresponse models.ResultResponse
		mockError    error
	}{
		{
			name:         "retieve success",
			bodyRequest:  bytes.NewBuffer(jsonStr),
			expectedCode: 200,
			expectedBody: "{\"success\":true}\n",
			mockresponse: models.ResultResponse{
				Success: true,
			},
		},
		{
			name:         "retieve failed by CreateSubscribeFriend error",
			bodyRequest:  bytes.NewBuffer(jsonStr),
			expectedCode: 500,
			expectedBody: "{\"statusCode\":500,\"message\":\"db err\"}\n",
			mockError:    errors.New("db err"),
		},
		{
			name:         "retieve failed by incorrect input",
			bodyRequest:  bytes.NewBuffer([]byte("")),
			expectedCode: 400,
			expectedBody: "{\"statusCode\":400,\"message\":\"Bad request\"}\n",
		},
	}

	for _, tt := range testCases {
		req, err := http.NewRequest(http.MethodPost, "/api/subscribeFriend", tt.bodyRequest)
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")
		require.NoError(t, err)
		router := chi.NewRouter()
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, router))
		rr := httptest.NewRecorder()
		serviceMock := new(mocks.StoreMock)
		serviceMock.On("CreateSubscribeFriend", mock.Anything, mock.Anything).Return(tt.mockresponse, tt.mockError)
		handler := createSubscribeFriend(serviceMock)
		handler.ServeHTTP(rr, req)
		require.Equal(t, tt.expectedCode, rr.Code)
		require.Equal(t, tt.expectedBody, rr.Body.String())
	}
}

func TestCreateBlockFriend(t *testing.T) {
	var jsonStr = []byte(`{"requestor":"andy@example","target":"john@example"}`)
	testCases := []struct {
		name         string
		bodyRequest  *bytes.Buffer
		expectedCode int
		expectedBody string
		mockresponse models.ResultResponse
		mockError    error
	}{
		{
			name:         "retieve success",
			bodyRequest:  bytes.NewBuffer(jsonStr),
			expectedCode: 200,
			expectedBody: "{\"success\":true}\n",
			mockresponse: models.ResultResponse{
				Success: true,
			},
		},
		{
			name:         "retieve failed by CreateBlockFriend error",
			bodyRequest:  bytes.NewBuffer(jsonStr),
			expectedCode: 500,
			expectedBody: "{\"statusCode\":500,\"message\":\"db err\"}\n",
			mockError:    errors.New("db err"),
		},
		{
			name:         "retieve failed by incorrect input",
			bodyRequest:  bytes.NewBuffer([]byte("")),
			expectedCode: 400,
			expectedBody: "{\"statusCode\":400,\"message\":\"Bad request\"}\n",
		},
	}

	for _, tt := range testCases {
		req, err := http.NewRequest(http.MethodPost, "/api/blockFriend", tt.bodyRequest)
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")
		require.NoError(t, err)
		router := chi.NewRouter()
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, router))
		rr := httptest.NewRecorder()
		serviceMock := new(mocks.StoreMock)
		serviceMock.On("CreateBlockFriend", mock.Anything, mock.Anything).Return(tt.mockresponse, tt.mockError)
		handler := createBlockFriend(serviceMock)
		handler.ServeHTTP(rr, req)
		require.Equal(t, tt.expectedCode, rr.Code)
		require.Equal(t, tt.expectedBody, rr.Body.String())
	}
}

func TestReceiveUpdate(t *testing.T) {
	var jsonStr = []byte(`{"sender":"andy@example","text":"Hello World! kate@example.com"}`)
	testCases := []struct {
		name         string
		bodyRequest  *bytes.Buffer
		expectedCode int
		expectedBody string
		mockresponse models.SendUpdateEmailResponse
		mockError    error
	}{
		{
			name:         "retieve success",
			bodyRequest:  bytes.NewBuffer(jsonStr),
			expectedCode: 200,
			expectedBody: "{\"success\":true,\"recipients\":[\"kate@example.com\"]}\n",
			mockresponse: models.SendUpdateEmailResponse{
				Success:    true,
				Recipients: []string{"kate@example.com"},
			},
		},
		{
			name:         "retieve failed by CreateReceiveUpdate error",
			bodyRequest:  bytes.NewBuffer(jsonStr),
			expectedCode: 500,
			expectedBody: "{\"statusCode\":500,\"message\":\"db err\"}\n",
			mockError:    errors.New("db err"),
		},
		{
			name:         "retieve failed by incorrect input",
			bodyRequest:  bytes.NewBuffer([]byte("")),
			expectedCode: 400,
			expectedBody: "{\"statusCode\":400,\"message\":\"Bad request\"}\n",
		},
	}

	for _, tt := range testCases {
		req, err := http.NewRequest(http.MethodPost, "/api/blockFriend", tt.bodyRequest)
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")
		require.NoError(t, err)
		router := chi.NewRouter()
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, router))
		rr := httptest.NewRecorder()
		serviceMock := new(mocks.StoreMock)
		serviceMock.On("CreateReceiveUpdate", mock.Anything, mock.Anything).Return(tt.mockresponse, tt.mockError)
		handler := receiveUpdate(serviceMock)
		handler.ServeHTTP(rr, req)
		require.Equal(t, tt.expectedCode, rr.Code)
		require.Equal(t, tt.expectedBody, rr.Body.String())
	}
}
