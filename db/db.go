package db

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct{}

var Conn *gorm.DB

func ConnectToDB() {
	if Conn == nil {
		connStr := "host=localhost user=sushant password=Sushant@8813! dbname=metrics_minder port=5432 sslmode=disable TimeZone=Asia/Kolkata"
		db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err != nil {
			log.Error().Msg("err: " + err.Error())
			panic(err)
		}

		Conn = db
		log.Info().Msg("Connected to database ✅")
	} else {
		log.Info().Msg("Already connected to database ✅")
	}
}
