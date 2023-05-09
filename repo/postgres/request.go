package request_history

import (
	"database/sql"
	"telegram-bot/entity"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(connStr string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) AddInHistory(req entity.Request) error {
	query := "INSERT INTO chat_history (chat_id, message, time) VALUES ($1, $2, $3)"
	_, err := s.db.Exec(query, req.ChatID, req.Command, req.Time)
	return err
}

func (s *PostgresStorage) GetChatHistory(chatID int) ([]entity.Request, error) {
	query := "SELECT message, time FROM chat_history WHERE chat_id = $1"
	rows, err := s.db.Query(query, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []entity.Request
	for rows.Next() {
		req := entity.Request{ChatID: chatID}
		err = rows.Scan(&req.Command, &req.Time)
		if err != nil {
			return nil, err
		}
		history = append(history, req)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return history, nil
}
