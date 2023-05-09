package entity

import "time"

type Request struct {
	ChatID  int
	Command Command
	Time    time.Time
}
