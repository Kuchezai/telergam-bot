package handler

import (
	"fmt"
	"telegram-bot/entity"
	"telegram-bot/handler/helpers"
	"telegram-bot/usecase"
)

type AboutBotHandler struct {
	au *usecase.AboutBotUsecase
}

func NewAboutBotHandler(au *usecase.AboutBotUsecase) *AboutBotHandler {
	return &AboutBotHandler{au}
}

func (h *AboutBotHandler) GetAuthorGitHub(chatID int, msg string) entity.Response {
	gitHub, err := h.au.GetAuthorGitHub(chatID)
	if err != nil {
		return entity.Response{}
	}
	textMsg := fmt.Sprintln(gitHub, "\n\nВыберите следующее действие")

	return helpers.CreateResponseWithTwoButton(chatID, textMsg, entity.ToMain, entity.InfoGitHub)
}
