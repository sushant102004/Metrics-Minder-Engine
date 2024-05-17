package models

import (
  "gorm.io/gorm"
  "time"
)

type QuickStats struct {
	gorm.Model
	Visitors       int
	PageViews      int
	NewVisitors    int
  TotalVisits    int
	PagesPerVisit  float64
	OrganicTraffic int
	DateRecorded   time.Time
	TimeRecorded   time.Time `gorm:"autoCreateTime"`
}
