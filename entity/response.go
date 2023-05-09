package entity

type Response struct {
	ChatID      int         `json:"chat_id"`
	Text        string      `json:"text"`
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
