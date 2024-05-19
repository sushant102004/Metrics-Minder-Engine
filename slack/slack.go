package slack

import (
	httpClient "metric-minder-engine/http"
	"os"

	"github.com/rs/zerolog/log"
)

func SendMessage(msgData string) {
	webhookURL := os.Getenv("SLACK_WEBHOOK")
	if webhookURL == "" {
		log.Error().Msg("Slack webhook not found")
		return
	}

	body := map[string]string{
		"text": msgData,
	}

	h := map[string]string{
		"Content-Type": "application/json",
	}

	httpClient.HTTPRequest("POST", webhookURL, body, nil, h)
}
