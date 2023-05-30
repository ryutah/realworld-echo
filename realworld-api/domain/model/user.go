package model

import "github.com/ryutah/realworld-echo/realworld-api/domain/premitive"

type UserID string

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
