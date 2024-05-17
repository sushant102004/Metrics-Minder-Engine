package db

import (
	"metric-minder-engine/models"
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct{}

var conn *gorm.DB

func ConnectToDB() {
	if conn == nil {
		connStr := "host=localhost user=sushant password=Sushant@8813! dbname=metrics_minder port=5432 sslmode=disable TimeZone=Asia/Kolkata"
		db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err != nil {
			log.Error().Msg("err: " + err.Error())
			os.Exit(-1)
		}

		conn = db
		log.Info().Msg("Connected to database ✅")
	} else {
		log.Info().Msg("Already connected to database ✅")
	}
}

func (db *DB) GetAllEmails() ([]string, error) {
	tx := conn.Find(&models.GoogleUser{})
	if tx.Error != nil {
		return nil, tx.Error
	}

	rows, err := tx.Rows()
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		rows.Scan()
	}

	return nil, nil
}

func (db *DB) GetAccessToken(email string) (string, error) {
  user := models.GoogleUser{}

  tx := conn.First(&user, "email = " + email)
  if tx.Error != nil {
    return "", tx.Error
  }
  
  return user.AccessToken, nil
}







