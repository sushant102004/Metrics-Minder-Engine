package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog/log"
)

var Log Logger

type Logger struct {
	file *os.File
}

func NewLog() {
	workingDir, err := os.Getwd()
	checkPanic(err)

	f, err := os.OpenFile(workingDir+"/error_logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	checkPanic(err)
	log.Info().Msg("Log file created ✅")

	Log = Logger{file: f}
}

func (l Logger) SaveNewLog(data string) {
	layout := "2006-01-02 15:04-05"
	time := time.Now().Format(layout)

	msg := fmt.Sprintf("%s - %s\n", time, data)

	Log.file.WriteString(msg)
}

func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}
