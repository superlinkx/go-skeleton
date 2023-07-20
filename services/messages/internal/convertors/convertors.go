package convertors

import (
	"github.com/superlinkx/go-skeleton/models"
	"github.com/superlinkx/go-skeleton/postgres"
)

func ConvertPostgresMessage(messages []postgres.Message) []models.Message {
	resultMessages := make([]models.Message, 0, len(messages))
	for _, message := range messages {
		resultMessages = append(resultMessages, models.Message{Message: message.Message})
	}
	return resultMessages
}
