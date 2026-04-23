package model

import "time"

type UserModel struct {
	ID              string
	Email           string
	Username        string
	Password        string
	PasswordConfirm string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
