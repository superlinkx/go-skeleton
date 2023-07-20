package convertors_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/superlinkx/go-skeleton/models"
	"github.com/superlinkx/go-skeleton/postgres"
	"github.com/superlinkx/go-skeleton/services/messages/internal/convertors"
)

func TestConvertPostgresMessage(t *testing.T) {
	var (
		testCases = []struct {
			Name  string
			Input []postgres.Message
			Want  []models.Message
		}{
			{
				Name:  "empty list",
				Input: []postgres.Message{},
				Want:  []models.Message{},
			},
			{
				Name: "single message",
				Input: []postgres.Message{
					{
						ID:      0,
						Message: "Hello world!",
					},
				},
				Want: []models.Message{
					{
						Message: "Hello world!",
					},
				},
			},
			{
				Name: "multiple messages",
				Input: []postgres.Message{
					{
						ID:      0,
						Message: "Hello world!",
					},
					{
						ID:      1,
						Message: "Goodbye world!",
					},
				},
				Want: []models.Message{
					{
						Message: "Hello world!",
					},
					{
						Message: "Goodbye world!",
					},
				},
			},
		}
	)

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			got := convertors.ConvertPostgresMessage(tc.Input)
			assert.Equal(t, tc.Want, got)
		})
	}
}
