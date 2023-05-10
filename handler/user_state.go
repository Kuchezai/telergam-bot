package handler

import (
	"fmt"
	"strconv"
	"strings"
	"telegram-bot/entity"
	"telegram-bot/entity/msgtxt"
	"telegram-bot/handler/helpers"
	"telegram-bot/usecase"
)

const (
	maxFriendsNumber = 100
)

var (
	errorUserNotFound  = fmt.Errorf("err: user not found")
	errorUserIsPrivate = fmt.Errorf("err: user is private")
	errorInvalidUser   = fmt.Errorf("err: invalid user_id")
)

type UserHandler struct {
	uc *usecase.VKUserUsecase
}

func NewUserHandler(uc *usecase.VKUserUsecase) *UserHandler {
	return &UserHandler{uc}
}

func (h *UserHandler) GetUser(chatID int, msg string) entity.Response {
	userID, err := strconv.Atoi(msg)
	if err != nil {
		return helpers.ResponseWithMainInfoBtnAndChoicePrompt(chatID, msgtxt.NotFoundOrBadVKUserID)
	}
	ans, err := h.uc.GetUser(userID, chatID)
	if err != nil {
		if err.Error() == errorUserIsPrivate.Error() {
			return helpers.ResponseWithMainInfoBtnAndChoicePrompt(chatID, msgtxt.UserIsPrivat)
		} else if err.Error() == errorUserNotFound.Error() || err.Error() == errorInvalidUser.Error() {
			return helpers.ResponseWithMainInfoBtnAndChoicePrompt(chatID, msgtxt.NotFoundOrBadVKUserID)
		} else {
			return helpers.ResponseWithMainInfoBtnAndChoicePrompt(chatID, msgtxt.SomethingWrongTryLatter)
		}
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("ID: %v\nИмя: %v\nФамилия: %v\n", ans.ID, ans.FirstName, ans.SecondName))

	return helpers.ResponseWithMainInfoBtnAndChoicePrompt(chatID, sb.String())
}

func (h *UserHandler) GetInfoAboutGeUser(chatID int, msg string) entity.Response {
	info := h.uc.GetInfoAboutGetUser()
	return helpers.ResponseWithMainInfoBtnAndChoicePrompt(chatID, info)
}

func (h *UserHandler) GetUserFriends(chatID int, msg string) entity.Response {
	userID, err := strconv.Atoi(msg)
	if err != nil {
		return helpers.ResponseWithMainInfoBtnAndChoicePrompt(chatID, msgtxt.NotFoundOrBadVKUserID)
	}
	friends, err := h.uc.GetUserFriends(userID, chatID)
	if err != nil {
		if err.Error() == errorUserIsPrivate.Error() {
			return helpers.ResponseWithMainInfoBtnAndChoicePrompt(chatID, msgtxt.UserIsPrivat)
		} else if err.Error() == errorUserNotFound.Error() || err.Error() == errorInvalidUser.Error() {
			return helpers.ResponseWithMainInfoBtnAndChoicePrompt(chatID, msgtxt.NotFoundOrBadVKUserID)
		} else {
			return helpers.ResponseWithMainInfoBtnAndChoicePrompt(chatID, msgtxt.SomethingWrongTryLatter)
		}
	}

	if len(friends) == 0 {
		return helpers.ResponseWithMainInfoBtnAndChoicePrompt(chatID, msgtxt.FriendsNotFound)
	}

	var sb strings.Builder
	sb.WriteString("\n")
	for i, friend := range friends {
		if i == maxFriendsNumber {
			break
		}
		sb.WriteString(fmt.Sprintf("%d. %s %s\n", i+1, friend.FirstName, friend.SecondName))
	}

	return helpers.ResponseWithMainInfoBtnAndChoicePrompt(chatID, sb.String())
}

func (h *UserHandler) GetInfoAboutGeUserFriends(chatID int, msg string) entity.Response {
	info := h.uc.GetInfoAboutGetUserFriends()
	return helpers.ResponseWithMainInfoBtn(chatID, info)
}
