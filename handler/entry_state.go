package handler

import (
	"telegram-bot/entity"
	"telegram-bot/entity/state"
	"telegram-bot/usecase"
)

type EntryStateHandler struct {
	su *usecase.StateRouterUsecase
}

func NewEntryStateHandler(su *usecase.StateRouterUsecase) *EntryStateHandler {
	return &EntryStateHandler{su}
}

func (h *EntryStateHandler) GoToMain(chatID int, msg string) entity.Response {
	h.su.ChangeChatState(chatID, state.Main)
	message := entity.Response{
		ChatID: chatID,
		Text:   "Пожалуйста, выберите действие:",
		ReplyMarkup: map[string]interface{}{
			"keyboard": [][]map[string]interface{}{
				{
					{"text": "Информация о пользователе"},
					{"text": "Информация о друзьях"},
				},
				{
					{"text": "История запросов"},
					{"text": "Информация об авторе"},
				},
			},
			"resize_keyboard": true,
		},
	}

	return message
}

func (h *EntryStateHandler) RestartBot(chatID int, msg string) entity.Response {
	h.su.ChangeChatState(chatID, state.EntryPoint)
	message := entity.Response{
		ChatID: chatID,
		Text:   "Добро пожаловать!:",
		ReplyMarkup: map[string]interface{}{
			"keyboard": [][]map[string]interface{}{
				{
					{"text": "В меню"},
				},
			},
			"resize_keyboard": true,
		},
	}
	return message
}

func (h *EntryStateHandler) UnknownCommand(chatID int, msg string) entity.Response {
	return entity.Response{
		ChatID: chatID,
		Text:   "Я не знаю такой команды",
	}
}
