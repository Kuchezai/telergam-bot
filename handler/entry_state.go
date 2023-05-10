package handler

import (
	"telegram-bot/entity"
	"telegram-bot/entity/command"
	"telegram-bot/entity/msgtxt"
	"telegram-bot/entity/state"
	"telegram-bot/handler/helpers"
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

	return helpers.ResponseWithFourBtn(chatID, msgtxt.ChoseAction, command.InfoAboutUser,
		command.InfoAboutUserFriends, command.RequestHistory, command.InfoAboutAuthor)
}

func (h *EntryStateHandler) RestartBot(chatID int, msg string) entity.Response {
	h.su.ChangeChatState(chatID, state.EntryPoint)

	return helpers.ResponseWithOneBtn(chatID, msgtxt.Welcome, command.ToMain)
}

func (h *EntryStateHandler) UnknownCommand(chatID int, msg string) entity.Response {
	return helpers.ResponseWithText(chatID, msgtxt.UnknownCommand)
}
