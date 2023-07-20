//go:build integration
// +build integration

package messages_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/superlinkx/go-skeleton/harnesses/integration"
	"github.com/superlinkx/go-skeleton/models"
	"github.com/superlinkx/go-skeleton/postgres"
	"github.com/superlinkx/go-skeleton/services/messages"
)

func TestGetDatabaseMessage_HappyPath(t *testing.T) {
	var (
		testCases = []struct {
			Name          string
			Input         int64
			Expected      models.Message
			ErrorContains string
		}{
			{
				Name:     "Get Message ID 1",
				Input:    1,
				Expected: models.Message{Message: "Hello Database 1!"},
			},
			{
				Name:     "Get Message ID 2",
				Input:    2,
				Expected: models.Message{Message: "Hello Database 2!"},
			},
			{
				Name:     "Get Message ID 3",
				Input:    3,
				Expected: models.Message{Message: "Hello Database 3!"},
			},
			{
				Name:     "Get Message ID 4",
				Input:    4,
				Expected: models.Message{Message: "Hello Database 4!"},
			},
			{
				Name:     "Get Message ID 5",
				Input:    5,
				Expected: models.Message{Message: "Hello Database 5!"},
			},
			{
				Name:     "Get Message ID 6",
				Input:    6,
				Expected: models.Message{Message: "Hello Database 6!"},
			},
			{
				Name:          "Get Invalid Message ID",
				Input:         7,
				Expected:      models.Message{},
				ErrorContains: "failed to get message (id=7): ",
			},
		}
	)

	db, err := integration.ConnectDatabase()
	require.NoError(t, err)
	defer db.Close()

	err = integration.MigrateDatabase(db)
	require.NoError(t, err)

	queries := postgres.New(db)
	messageService := messages.NewMessageService(queries)

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			message, err := messageService.GetDatabaseMessage(context.Background(), tc.Input)
			if tc.ErrorContains != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.ErrorContains)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.Expected, message)
			}
		})
	}
}

func TestGetOddDatabaseMessages(t *testing.T) {
	var (
		expected = []models.Message{
			{Message: "Hello Database 1!"},
			{Message: "Hello Database 3!"},
			{Message: "Hello Database 5!"},
		}
	)

	db, err := integration.ConnectDatabase()
	require.NoError(t, err)
	defer db.Close()

	err = integration.MigrateDatabase(db)
	require.NoError(t, err)

	queries := postgres.New(db)
	messageService := messages.NewMessageService(queries)

	messages, err := messageService.GetOddDatabaseMessages(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expected, messages)
}

func BenchmarkIntegrationTests(b *testing.B) {
	for i := 0; i < b.N; i++ {
		db, err := integration.ConnectDatabase()
		require.NoError(b, err)
		defer db.Close()

		err = integration.MigrateDatabase(db)
		require.NoError(b, err)

		queries := postgres.New(db)
		messageService := messages.NewMessageService(queries)

		_, err = messageService.GetOddDatabaseMessages(context.Background())
		require.NoError(b, err)
	}
}
