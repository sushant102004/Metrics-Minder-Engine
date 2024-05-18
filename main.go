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
	usersDB := db.NewUsersDB()
	dataExtSvc := services.NewDataExtracter(statsDB)

	engine := services.NewEngine(dataExtSvc, usersDB)
	engine.Start()
}
