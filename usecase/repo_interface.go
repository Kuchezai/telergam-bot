package usecase

import (
	"telegram-bot/entity"
	"telegram-bot/entity/state"
)

type RequestsAdderRepo interface {
	AddInHistory(entity.Request) error
}

type RequestRepo interface {
	RequestsAdderRepo
	GetChatHistory(chatID int) ([]entity.Request, error)
}

type UserRepo interface {
	User(userID int) (entity.User, error)
	FriendsByID(userID int) ([]entity.User, error)
}

type StateChangerRepo interface {
	ChangeChatState(chatID int, newState state.State) error
}
