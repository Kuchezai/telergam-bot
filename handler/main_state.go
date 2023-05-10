package handler

import (
	"fmt"
	"strings"
	"telegram-bot/entity"
	"telegram-bot/entity/command"
	"telegram-bot/entity/msgtxt"
	"telegram-bot/entity/state"
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
	h.su.ChangeChatState(chatID, state.GetUser)

	return helpers.ResponseWithMainInfoBtn(chatID, msgtxt.InputVKIDOrChooseNextAction)
}

func (h *MainStateHandler) GoToGetUserFriends(chatID int, msg string) entity.Response {
	h.su.ChangeChatState(chatID, state.GetUserFriends)

	return helpers.ResponseWithMainInfoBtn(chatID, msgtxt.InputVKIDOrChooseNextAction)
}

func (h *MainStateHandler) GoToGetRequestHistory(chatID int, msg string) entity.Response {
	h.su.ChangeChatState(chatID, state.GetHistory)
	requests, err := h.ru.GetChatRequests(chatID)
	if err != nil {
		return entity.Response{}
	}

	var sb strings.Builder
	sb.WriteString(msgtxt.RequestHistory + "\n")
	for i, req := range requests {
		sb.WriteString(fmt.Sprintf("%d. %s (%s)\n", i+1, req.Command, req.Time.Format("2006-01-02 15:04:05")))
	}

	return helpers.ResponseWithMainInfoBtnAndChoicePrompt(chatID, sb.String())
}

func (h *MainStateHandler) GoToGetInfAboutAuthor(chatID int, msg string) entity.Response {
	h.su.ChangeChatState(chatID, state.GetInfoAboutAuthor)
	info, err := h.au.GetInfoAboutAuthor(chatID)
	if err != nil {
		return entity.Response{}
	}
	textMsg := fmt.Sprint(info, "\n\n", msgtxt.ChooseNextAction)
	return helpers.ResponseWithTwoBtn(chatID, textMsg, command.ToMain, command.InfoGitHub)
}
