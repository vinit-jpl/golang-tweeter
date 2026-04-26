package model

import "time"

type UserModel struct {
	ID              int64
	Email           string
	Username        string
	Password        string
	PasswordConfirm string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type RefreshTokenModel struct {
	ID           int64
	UserID       int64
	RefreshToken string
	ExpiredAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
