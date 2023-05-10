package helpers

import (
	"telegram-bot/entity"
	"telegram-bot/entity/command"
	"telegram-bot/entity/msgtxt"
)

func ResponseWithText(chatID int, text string) entity.Response {
	return entity.Response{
		ChatID: chatID,
		Text:   text,
	}
}

func ResponseWithOneBtn(chatID int, text string, buttonText command.Command) entity.Response {
	replyMarkup := map[string]interface{}{
		"keyboard": [][]map[string]interface{}{
			{
				{"text": buttonText},
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

func ResponseWithTwoBtn(chatID int, text string, btnOneText, btnTwoText command.Command) entity.Response {
	replyMarkup := map[string]interface{}{
		"keyboard": [][]map[string]interface{}{
			{
				{"text": btnOneText},
				{"text": btnTwoText},
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

func ResponseWithFourBtn(chatID int, text string, btnOneText, btnTwoText, btnThreeText, btnFourText command.Command) entity.Response {
	replyMarkup := map[string]interface{}{
		"keyboard": [][]map[string]interface{}{
			{
				{"text": btnOneText},
				{"text": btnTwoText},
			},
			{
				{"text": btnThreeText},
				{"text": btnFourText},
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

func ResponseWithMainInfoBtn(chatID int, text string) entity.Response {
	return ResponseWithTwoBtn(chatID, text, command.ToMain, command.InfoAboutCommand)
}

func ResponseWithMainInfoBtnAndChoicePrompt(chatID int, text string) entity.Response {
	return ResponseWithMainInfoBtn(chatID, text+"\n\n"+msgtxt.ChooseNextAction)
}
