package models

import "time"

type (
	UserModel struct {
		UserID           int64
		Username         string
		Password         string
		RegistrationDate time.Time
	}

	//RefreshTokenModel struct {
	//	ID           int64
	//	UserID       int64
	//	RefreshToken string
	//	ExpiredAt    string
	//	CreatedAt    string
	//	UpdatedAt    string
	//}
)
