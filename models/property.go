package models

import "gorm.io/gorm"

type PropertyToMonitor struct {
	gorm.Model
	Email      string `gorm:"primaryKey"`
	AccountID  string
	PropertyID string
}
