package database

import (
	"FriendManagementAPI/models"
)

func (db Database) GetTargetBlockedByRequest(requestor, target string) (int, error) {
	var countBlocked int
	query := `
	SELECT 
		COUNT(*) 
	FROM 
		block b
	WHERE 
		b.requestor = $1 AND b.target = $2;`
	row := db.Conn.QueryRow(query, requestor, target)
	err := row.Scan(&countBlocked)
	if err != nil {
		return countBlocked, err
	}
	return countBlocked, nil
}

func (db Database) GetTargetSubscribeByRequest(requestor, target string) (int, error) {
	var countSubscribe int
	query := `
	SELECT 
		COUNT(*) 
	FROM 
		subscription s
	WHERE 
		s.requestor = $1 AND s.target = $2;`
	row := db.Conn.QueryRow(query, requestor, target)
	err := row.Scan(&countSubscribe)
	if err != nil {
		return countSubscribe, err
	}
	return countSubscribe, nil
}

func (db Database) GetBlockListByRequest(req *models.FriendConnectionRequest) (int, error) {
	var countBlock int
	query := `
	SELECT 
		COUNT(*) 
	FROM 
		block b
	WHERE 
		b.requestor = $1 AND b.target = $2 
		OR 
		b.requestor = $2 AND b.target = $1;`
	row := db.Conn.QueryRow(query, req.Friends[0], req.Friends[1])
	err := row.Scan(&countBlock)
	if err != nil {
		return countBlock, err
	}
	return countBlock, nil
}

func (db Database) GetFriendListByRequest(req *models.FriendConnectionRequest) (int, error) {
	var count int
	query := `
	SELECT 
		COUNT(*) 
	FROM 
		friend f 
	WHERE 
		f.emailuserone = $1 AND f.emailusertwo = $2 
		OR 
		f.emailuserone = $2 AND f.emailusertwo = $1;`
	row := db.Conn.QueryRow(query, req.Friends[0], req.Friends[1])
	err := row.Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

func (db Database) GetUserByRequest(emailUserOne, emailUserTwo string) (int, error) {
	var count int
	query := `
	SELECT count(*) 
	FROM userprofile u 
	WHERE 
		u.email in($1,$2);`
	row := db.Conn.QueryRow(query, emailUserOne, emailUserTwo)
	err := row.Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

// CreateFriend executes create a friend connection between two email addresses
func (db Database) CreateFriend(item *models.FriendConnectionRequest) error {
	emailUserOne := item.Friends[0]
	emailUserTwo := item.Friends[1]
	query := `INSERT INTO friend (emailuserone, emailusertwo) VALUES ($1, $2);`
	_, err := db.Conn.Exec(query, emailUserOne, emailUserTwo)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByEmail executes get user from table 'userprofile' by email and return an email address
func (db Database) GetUserByEmail(email string) (int, error) {
	var count int
	query := `
	SELECT COUNT(*) 
	FROM userprofile u 
	WHERE 
		u.email = $1;`
	row := db.Conn.QueryRow(query, email)
	err := row.Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

// GetFriendListByEmail executes retrieve the friends list for an email address
func (db Database) GetFriendListByEmail(email string) (*models.User, error) {
	friendLst := &models.User{}
	query := `
	SELECT emailusertwo, id 
	FROM friend  
	WHERE 
		emailuserone = $1 
	UNION 
	SELECT emailuserone, id 
	FROM friend 
	WHERE 
		emailusertwo =  $1 
	ORDER BY id;`
	rows, err := db.Conn.Query(query, email)
	if err != nil {
		return friendLst, err
	}
	var id int
	for rows.Next() {
		var item models.User
		err := rows.Scan(&item.Email, &id)
		if err != nil {
			return friendLst, err
		}
		friendLst.Friends = append(friendLst.Friends, item.Email)
	}
	return friendLst, nil
}

// CreateSubscribeFriendByRequestorAndTarget executes subscribe to updates from an email address
func (db Database) CreateUserByEmail(email string) error {
	query := `INSERT INTO userprofile (email) VALUES ($1);`
	_, err := db.Conn.Exec(query, email)
	if err != nil {
		return err
	}
	return nil
}

// CreateSubscribeFriendByRequestorAndTarget executes subscribe to updates from an email address
func (db Database) CreateSubscribeFriendByRequestorAndTarget(requestor, target string) error {
	query := `INSERT INTO subscription (requestor, target) VALUES ($1, $2);`
	_, err := db.Conn.Exec(query, requestor, target)
	if err != nil {
		return err
	}
	return nil
}

// CreateBlockFriendByRequestorAndTarget is block updates from an email address
func (db Database) CreateBlockFriendByRequestorAndTarget(requestor, target string) error {
	query := `INSERT INTO block (requestor, target) VALUES ($1, $2);`
	_, err := db.Conn.Exec(query, requestor, target)
	if err != nil {
		return err
	}
	return nil
}

// GetAllUser executes get all user from table 'userprofile'
func (db Database) GetAllUser() ([]models.User, error) {
	allUser := []models.User{}
	query := `SELECT * FROM userprofile;`
	rows, err := db.Conn.Query(query)
	if err != nil {
		return allUser, err
	}
	for rows.Next() {
		var item models.User
		err := rows.Scan(&item.Email)
		if err != nil {
			return allUser, err
		}
		allUser = append(allUser, item)
	}
	return allUser, nil
}

// GetAllSubscriberByEmail executes get subscriber by email return subscriber list
func (db Database) GetAllSubscriberByEmail(requestor string) (*models.User, error) {
	targetLst := &models.User{}
	query := `SELECT s.requestor FROM subscription s WHERE s.target = $1;`
	rows, err := db.Conn.Query(query, requestor)
	if err != nil {
		return targetLst, err
	}
	for rows.Next() {
		var item models.SubscriptionRequest
		err := rows.Scan(&item.Requestor)
		if err != nil {
			return targetLst, err
		}
		targetLst.Subscription = append(targetLst.Subscription, item.Requestor)
	}
	return targetLst, nil
}

// GetAllBlockerByEmail executes get target by requestor return target
func (db Database) GetAllBlockerByEmail(requestor string) (*models.User, error) {
	targetLst := &models.User{}
	query := `SELECT b.requestor FROM block b WHERE b.target = $1;`
	rows, err := db.Conn.Query(query, requestor)
	if err != nil {
		return targetLst, err
	}
	for rows.Next() {
		var item models.BlockRequest
		err := rows.Scan(&item.Requestor)
		if err != nil {
			return targetLst, err
		}
		targetLst.Blocked = append(targetLst.Blocked, item.Requestor)
	}
	return targetLst, nil
}
