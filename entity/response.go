package entity

// entity for messages sent by the bot
type Response struct {
	ChatID      int
	Text        string
	ReplyMarkup interface{}
}
