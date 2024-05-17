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


  statsDB := db.NewStatsDB()
  dataExtSvc := services.NewDataExtracter(statsDB)
  dataExtSvc.GetAndSaveQuickStats("ya29.a0AXooCgtOArZ0nJYwAH3trjOG3h1M4c_TiPpoiegqt-_TVEi72riZeky5qlYyWAq7MzjetIRMeSPvOxXeVaCap7OVnj_VI7nqAGPpEApLc0TI-brFvbLW3Pnjj2-dz4WCGnUQM0AwfBPbr1A36u2nUwRw0xYDJwrZ27znaCgYKAQ8SARMSFQHGX2Miy5iXvxwY5x8Loha8KfAFag0171", "420272748")
}
