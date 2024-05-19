package models

import (
	"gorm.io/gorm"
)

type QuickStats struct {
	gorm.Model
	Visitors       int
	PageViews      int
	NewVisitors    int
	TotalVisits    int
	PagesPerVisit  float64
	OrganicTraffic int
	User           string
}
