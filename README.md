# The Friend Management RESTFul Api
Introduction
` simple api built using Golang 1.15.2`

## Verify it
`go version`
## Build the project
`docker-compose build`

## Launch the project
`docker-compose up`

## RESTFul Api
```sh
1 User Registration: http://localhost:8080/api/registration
  Example Request
    {
	  "email" : "phong@s3corp.com.vn"
    }
  Response Example
    {
      "success":true
    }
2 Create a friend connection : http://localhost:8080/api/friendConnection
  Example Request
    {
      "friends":[
        "phong@s3corp.com.vn",
        "hien@s3corp.com.vn"
        ]
    }
  Response Example
    {
      "success":true
    }
3 Retrieve the friends list for an email address :  http://localhost:8080/api/friendList
  Example Request
    {
    "email":"phongg@s3corp.com.vn"
    }
  Response Example
    {
        "success": true,
        "friends": [
            "hien@s3corp.com.vn"
        ],
        "count": 1
    }
4 Retrieve the common friends list between two email addresses :  http://localhost:8080/api/commonFriend
  Example Request
      {
        "friends":[
          "phong@s3corp.com.vn",
          "hien@s3corp.com.vn"
          ]
      }
  Response Example
      {
        "success": true,
        "friends": [
            "thinh@s3corp.com.vn"
        ],
        "count": 1
      }
5 Create subscribe to updates from an email address : http://localhost:8080/api/subscribeFriend
  Example Request
    {
      "requestor" : "phong@s3corp.com.vn",
      "target" : "hien@s3corp.com.vn"
    }
  Response Example
    {
      "success": true
    }
6 Block updates from an email address: http://localhost:8080/api/blockFriend
  Example Request
    {
      "requestor" : "phong@s3corp.com.vn",
      "target" : "thinh@s3corp.com.vn"
    }
  Response Example
    {
      "success": true
    }
7 Create API to retrieve all email addresses that can receive updates from an email address :  http://localhost:8080/api/receiveUpdates
  Example Request
    {
      "sender" :"phong@s3corp.com.vn",
      "text" : "Hello World! dat@s3corp.com.vn"
    }
   Response Example
    {
      "success": true,
      "recipients": [
          "dat",
          "thinh@s3corp.com.vn",
          "dat@s3corp.com.vn"
      ]
    }
```
## Unit Testing
### How to run

From the terminal, in the solution root, simply run:

### Go to service commands folder
cd service
## Run all tests
go test -v
### Go to handlers commands folder
cd handlers
## Run all tests
go test -v
