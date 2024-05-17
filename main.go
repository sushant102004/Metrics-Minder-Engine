package main

import (
	"metric-minder-engine/db"
	"metric-minder-engine/services"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
  zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
  log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123})
}

func main() {
	db.ConnectToDB()
	db.AutoMigrate()

  dataExtSvc := services.NewDataExtracter()
  dataExtSvc.GetAndSaveQuickStats("ya29.a0AXooCgvq1Cdnnmu8xFRcaY0VuJwFeISw-JYxvRhCBxE5hYsvzGb0-vPC9GizN0XghCJEdRPADz80yofSjwG6fv6ZdpNZdFlwfa5ETXi7zfLem6hzEl2J2Xkgh0nSJdA0ptfN1mLUX6Yhs-45DlOcSS0ktpCz6q9dDMrzaCgYKAUwSARMSFQHGX2MiHvfc2MAB9BEj1TmaFtI3Zw0171", "420272748")
}
