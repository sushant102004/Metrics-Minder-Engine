package db

import (
	"metric-minder-engine/models"

	"gorm.io/gorm"
)

type StatsDB struct {
	conn *gorm.DB
}

func NewStatsDB() StatsDB {
	return StatsDB{
		conn: Conn,
	}
}

func (db *StatsDB) SaveStats(stats models.QuickStats) error {
	// TODO - Later find a way to get data multiple times and save in a single row.
	tx := db.conn.Save(&stats)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
