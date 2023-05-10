package handler

import (
	"fmt"
	"telegram-bot/entity"
	"telegram-bot/entity/command"
	"telegram-bot/entity/msgtxt"
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
	textMsg := fmt.Sprint(gitHub, "\n\n", msgtxt.ChooseNextAction)

	return helpers.ResponseWithTwoBtn(chatID, textMsg, command.ToMain, command.InfoGitHub)
}
