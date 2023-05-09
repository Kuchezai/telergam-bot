package helpers

import (
	"telegram-bot/entity"
	"telegram-bot/entity/command"
)

func CreateResponseWithTwoButton(chatID int, text string, buttonOneText, buttonTwoText command.Command) entity.Response {
	replyMarkup := map[string]interface{}{
		"keyboard": [][]map[string]interface{}{
			{
				{"text": buttonOneText},
				{"text": buttonTwoText},
			},
		},
		"resize_keyboard": true,
	}

	response := entity.Response{
		ChatID:      chatID,
		Text:        text,
		ReplyMarkup: replyMarkup,
	}

	return response
}

func CreateResponseWithMainAndInfoButton(chatID int, text string) entity.Response {
	return CreateResponseWithTwoButton(chatID, text, command.ToMain, command.InfoAboutCommand)
}
