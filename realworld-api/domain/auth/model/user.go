package model

import "github.com/ryutah/realworld-echo/realworld-api/domain/premitive"

type UserID premitive.UID

func NewUserID(s string) (UserID, error) {
	uid, err := premitive.NewUID(s)
	if err != nil {
		return "", err
	}
	return UserID(uid), nil
}

func (u UserID) String() string {
	return premitive.UID(u).String()
}

type User struct {
	ID      UserID
	Account *Account
	Profile *Profile
}

type Account struct {
	Email premitive.Email
}

type Profile struct {
	Username premitive.Name
	Image    premitive.URL
}
