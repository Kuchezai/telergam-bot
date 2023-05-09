package handler

import (
	"fmt"
	"telegram-bot/entity"
	"telegram-bot/handler/helpers"
	"telegram-bot/usecase"
)

type RequestHandler struct {
	ru *usecase.RequestUsecase
}

func NewRequestHandler(ru *usecase.RequestUsecase) *RequestHandler {
	return &RequestHandler{ru}
}

func (h *RequestHandler) GetInfoAboutChatRequests(chatID int, msg string) entity.Response {
	info := h.ru.GetInfoAboutGetChatRequests()
	msgText := fmt.Sprintln(info, "\n\nВыберите следующее действие")
	return helpers.CreateResponseWithMainAndInfoButton(chatID, msgText)
}
