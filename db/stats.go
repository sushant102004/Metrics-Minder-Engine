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
	tx := db.conn.Save(&stats)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
