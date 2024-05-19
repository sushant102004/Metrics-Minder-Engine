package main

import (
	"metric-minder-engine/db"
	logger "metric-minder-engine/log"
	"metric-minder-engine/services"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123})

	err := godotenv.Load()
	checkPanic(err)
}

func main() {
	logger.NewLog()

	db.ConnectToDB()
	db.AutoMigrate()

	statsDB := db.NewStatsDB()
	usersDB := db.NewUsersDB()
	dataExtSvc := services.NewDataExtracter(statsDB)

	engine := services.NewEngine(dataExtSvc, usersDB)
	engine.Start()
}

func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}
