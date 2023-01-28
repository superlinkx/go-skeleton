package messages

import (
	"database/sql"
	"fmt"
)

type Queryer interface {
	Query(string, ...any) (*sql.Rows, error)
}

type Message struct {
	Message string `json:"message"`
}

func GetHelloMessage() Message {
	return Message{Message: "Welcome to the API"}
}

func GetDatabaseMessage(queryer Queryer) (Message, error) {
	if rows, err := queryer.Query("SELECT message FROM messages LIMIT(1);"); err != nil {
		return Message{}, fmt.Errorf("failed to execute query: %w", err)
	} else {
		var message string
		defer rows.Close()
		rows.Next()
		if err := rows.Scan(&message); err != nil {
			return Message{}, fmt.Errorf("error while scanning message: %w", err)
		} else {
			return Message{
				Message: message,
			}, nil
		}
	}
}
