package entity

import (
	"time"

	"telegram-bot/entity/command"
)

// entity for requests stored in history
type Request struct {
	ChatID  int
	Command command.Command
	Time    time.Time
}
