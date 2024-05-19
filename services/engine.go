package services

import (
	"metric-minder-engine/db"
	logger "metric-minder-engine/log"
	"metric-minder-engine/slack"
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

func (e Engine) SaveStats() error {
	emails, err := e.usersDB.GetAllEmails()
	if err != nil {
		log.Error().Msg("error: " + err.Error())
		slack.SendMessage(err.Error())
		log.Error().Msg("error: " + err.Error())
		return err
	}

	maxWorkers := 5
	workersChan := make(chan struct{}, maxWorkers)

	for _, email := range emails {
		if email == "" {
			continue
		}

		accessToken, err := e.usersDB.GetAccessToken(email)
		if err != nil {
			logger.Log.SaveNewLog(err.Error())
			slack.SendMessage(err.Error())
			log.Error().Msg("error: " + err.Error())
			continue
		}

		propertyID, err := e.usersDB.GetPropertyID(email)
		if err != nil {
			logger.Log.SaveNewLog(err.Error())
			slack.SendMessage(err.Error())
			log.Error().Msg("error: " + err.Error())
			continue
		}

		workersChan <- struct{}{} // This would block program if channel is full
		go func(email, accessToken, propertyID string) {
			err = e.dataExtractSvc.GetAndSaveQuickStats(accessToken, propertyID, email)
			if err != nil {
				logger.Log.SaveNewLog(err.Error())
				slack.SendMessage(err.Error())
				log.Error().Msg("error: " + err.Error())
			}
			<-workersChan
		}(email, accessToken, propertyID)
	}
	return nil
}

func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}
