package db

import (
	"metric-minder-engine/models"

	"github.com/rs/zerolog/log"
)

func AutoMigrate() {
	err := Conn.AutoMigrate(models.QuickStats{})
	if err != nil {
		log.Error().Msg("err: " + err.Error())
		panic(err)
	}

	log.Info().Msg("Database migration completed ✅")
}
