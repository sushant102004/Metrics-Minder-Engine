package services

import (
	"metric-minder-engine/db"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/rs/zerolog/log"
)

type Engine struct {
	dataExtractSvc DataExtracter
	usersDB        db.UsersDB
}

func NewEngine(svc DataExtracter, db db.UsersDB) Engine {
	return Engine{
		dataExtractSvc: svc,
		usersDB:        db,
	}
}

func (e Engine) Start() {
	e.StartScheduler(time.Second * 10)
}

func (e Engine) StartScheduler(interval time.Duration) {
	s, err := gocron.NewScheduler()
	checkPanic(err)

	s.NewJob(gocron.DurationJob(interval), gocron.NewTask(e.SaveStats))
	s.Start()

	select {}
}

func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func (e Engine) SaveStats() error {
	emails, err := e.usersDB.GetAllEmails()
	if err != nil {
		// Save this to error logs
		// notify to me.
		log.Error().Msg("error: " + err.Error())
		return err
	}

	log.Info().Msgf("Emails: %v", emails)

	for _, email := range emails {
		accessToken, err := e.usersDB.GetAccessToken(email)
		if err != nil {
			log.Error().Msg("error: " + err.Error())
			continue
		}

		propertyID, err := e.usersDB.GetPropertyID(email)
		if err != nil {
			log.Error().Msg("error: " + err.Error())
			continue
		}

		err = e.dataExtractSvc.GetAndSaveQuickStats(accessToken, propertyID)
		if err != nil {
			log.Error().Msg("error: " + err.Error())
			continue
		}
	}

	return nil
}
