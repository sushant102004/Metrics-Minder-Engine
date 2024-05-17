package db

import (
	"metric-minder-engine/models"
	"os"

	"github.com/rs/zerolog/log"
)

func AutoMigrate() {
	err := Conn.AutoMigrate(models.QuickStats{})
	if err != nil {
		log.Error().Msg("err: " + err.Error())
		os.Exit(-1)
	}

	log.Info().Msg("Database migration completed ✅")
}
