package handler

import (
	"fmt"
	"strings"
	"telegram-bot/entity"
	"telegram-bot/entity/command"
	"telegram-bot/handler/helpers"
	"telegram-bot/usecase"
)

type MainStateHandler struct {
	su *usecase.StateRouterUsecase
	ru *usecase.RequestUsecase
	au *usecase.AboutBotUsecase
}

func NewMainStateHandler(su *usecase.StateRouterUsecase, ru *usecase.RequestUsecase, au *usecase.AboutBotUsecase) *MainStateHandler {
	return &MainStateHandler{su, ru, au}
}

func (h *MainStateHandler) GoToGetUser(chatID int, msg string) entity.Response {
	h.su.ChangeChatState(chatID, entity.GetUser)

	textMsg := "Пожалуйста, введите ID пользователя VK или выберите следующее действие"

	return helpers.CreateResponseWithMainAndInfoButton(chatID, textMsg)
}

func (h *MainStateHandler) GoToGetUserFriends(chatID int, msg string) entity.Response {
	h.su.ChangeChatState(chatID, entity.GetUserFriends)

	textMsg := "Пожалуйста, введите ID пользователя VK или выберите следующее действие"

	return helpers.CreateResponseWithMainAndInfoButton(chatID, textMsg)
}

func (h *MainStateHandler) GoToGetRequestHistory(chatID int, msg string) entity.Response {
	h.su.ChangeChatState(chatID, entity.GetHistory)
	requests, err := h.ru.GetChatRequests(chatID)
	if err != nil {
		return entity.Response{}
	}

	var sb strings.Builder
	sb.WriteString("История запросов:\n")
	for i, req := range requests {
		sb.WriteString(fmt.Sprintf("%d. %s (%s)\n", i+1, req.Command, req.Time.Format("2006-01-02 15:04:05")))
	}

	sb.WriteString("\n\nВыберите следующее действие")

	return helpers.CreateResponseWithMainAndInfoButton(chatID, sb.String())
}

func (h *MainStateHandler) GoToGetInfAboutAuthor(chatID int, msg string) entity.Response {
	h.su.ChangeChatState(chatID, entity.GetInfoAboutAuthor)
	info, err := h.au.GetInfoAboutAuthor(chatID)
	if err != nil {
		return entity.Response{}
	}
	textMsg := fmt.Sprintln(info, "\n\nВыберите следующее действие")

	return helpers.CreateResponseWithTwoButton(chatID, textMsg, command.ToMain, command.InfoGitHub)
}
