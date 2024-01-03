package models

import "errors"

// Ideally you should only have one type per file in this folder.
// You can also make a constructor function for the type, which is what I have done here.
// This is a good place to add validation logic for the type.
// You can also add methods to the type, which is a good place to add business logic,
// such as calculating a total price for an order, or calculating the age of a user.

type User struct {
	UserID   string `json:"userid" bson:"userid"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
}

func NewUser(userid, username, password, email string) (*User, error) {

	if userid == "" || username == "" || password == "" || email == "" {
		return &User{}, errors.New("Missing required fields")
	}

	if len(userid) < 5 {
		return &User{}, errors.New("Userid must be at least 5 characters")
	}

	return &User{
		UserID:   userid,
		Username: username,
		Password: password,
		Email:    email,
	}, nil
}

func (u *User) Validate() error {
	// Add validation logic here
	return nil
}
