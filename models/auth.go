package models

import (
  "gorm.io/gorm"
)

type GoogleUser struct {
	gorm.Model
	ID            uint
	Name          *string `json:"name"`
	GivenName     *string `json:"given_name"`
	Picture       *string `json:"picture"`
	Email         string  `json:"email" gorm:"uniqueIndex"`
	EmailVerified bool    `json:"email_verified"`
	AccessToken   string  `json:"access_token"`
	RefreshToken  string  `json:"refresh_token"`
	ExpireAt      *int64  `json:"expire"`
	OnboardStatus string  `json:"onboard_status"` // SetupPending, SetupCompleted
}
