package entity

import (
	"time"

	"telegram-bot/entity/command"
)

type Request struct {
	ChatID  int
	Command command.Command
	Time    time.Time
}
