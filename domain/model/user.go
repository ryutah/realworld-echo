package model

import "github.com/ryutah/realworld-echo/domain/premitive"

type User struct {
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
