package handler

import (
	"fmt"
	"strconv"
	"strings"
	"telegram-bot/entity"
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
		return helpers.CreateResponseWithMainAndInfoButton(chatID, "Пользователь не найден, скорее всего введный ID некорректен \n\nВыберите следующее действие")
	}
	ans, err := h.uc.GetUser(userID, chatID)
	if err != nil {
		if err.Error() == errorUserIsPrivate.Error() {
			return helpers.CreateResponseWithMainAndInfoButton(chatID, "Указаный пользователь имеет приватный профиль \n\nВыберите следующее действие")
		} else if err.Error() == errorUserNotFound.Error() || err.Error() == errorInvalidUser.Error() {
			return helpers.CreateResponseWithMainAndInfoButton(chatID, "Пользователь не найден, скорее всего введный ID некорректен \n\nВыберите следующее действие")
		} else {
			return helpers.CreateResponseWithMainAndInfoButton(chatID, "Что-то пошло не так. Попробуйте позже \n\nВыберите следующее действие")
		}
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("ID: %v\nИмя: %v\nФамилия: %v\n", ans.ID, ans.FirstName, ans.SecondName))
	sb.WriteString("\n\nВыберите следующее действие")

	return helpers.CreateResponseWithMainAndInfoButton(chatID, sb.String())
}

func (h *UserHandler) GetInfoAboutGeUser(chatID int, msg string) entity.Response {
	info := h.uc.GetInfoAboutGetUser()
	textMsg := fmt.Sprintln(info, "\n\nВыберите следующее действие")
	return helpers.CreateResponseWithMainAndInfoButton(chatID, textMsg)
}

func (h *UserHandler) GetUserFriends(chatID int, msg string) entity.Response {
	userID, err := strconv.Atoi(msg)
	if err != nil {
		return helpers.CreateResponseWithMainAndInfoButton(chatID, "Пользователь не найден, скорее всего введный ID некорректен \n\nВыберите следующее действие")
	}
	friends, err := h.uc.GetUserFriends(userID, chatID)
	if err != nil {
		if err.Error() == errorUserIsPrivate.Error() {
			return helpers.CreateResponseWithMainAndInfoButton(chatID, "Указаный пользователь имеет приватный профиль \n\nВыберите следующее действие")
		} else if err.Error() == errorUserNotFound.Error() || err.Error() == errorInvalidUser.Error() {
			return helpers.CreateResponseWithMainAndInfoButton(chatID, "Пользователь не найден, скорее всего введный ID некорректен \n\nВыберите следующее действие")
		} else {
			return helpers.CreateResponseWithMainAndInfoButton(chatID, "Что-то пошло не так. Попробуйте позже \n\nВыберите следующее действие")
		}
	}

	var sb strings.Builder
	sb.WriteString("Список друзей:\n")
	for i, friend := range friends {
		if i == maxFriendsNumber {
			break
		}
		sb.WriteString(fmt.Sprintf("%d. %s %s\n", i+1, friend.FirstName, friend.SecondName))
	}
	sb.WriteString("\n\nВыберите следующее действие")

	return helpers.CreateResponseWithMainAndInfoButton(chatID, sb.String())
}

func (h *UserHandler) GetInfoAboutGeUserFriends(chatID int, msg string) entity.Response {
	info := h.uc.GetInfoAboutGetUserFriends()

	textMsg := fmt.Sprintln(info, "\n\nВыберите следующее действие")
	return helpers.CreateResponseWithMainAndInfoButton(chatID, textMsg)
}
