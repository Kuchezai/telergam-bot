package usecase

import (
	"telegram-bot/entity"
	"time"
)

type VKUserUsecase struct {
	userRepo    UserRepo
	requestRepo RequestsAdderRepo
	stateRepo   StateChangerRepo
}

func NewVKUserUsecase(ur UserRepo, rr RequestsAdderRepo, sr StateChangerRepo) *VKUserUsecase {
	return &VKUserUsecase{ur, rr, sr}
}

func (u *VKUserUsecase) GetUser(userID int, chatID int) (entity.User, error) {
	u.requestRepo.AddInHistory(entity.Request{ChatID: chatID, Command: entity.InfoAboutUser, Time: time.Now()})
	return u.userRepo.User(userID)
}
func (u *VKUserUsecase) GetInfoAboutGetUser() string {
	return "Это команда выводит имя и фамилию пользователя VK"
}

func (u *VKUserUsecase) GetUserFriends(userID int, chatID int) ([]entity.User, error) {
	u.requestRepo.AddInHistory(entity.Request{ChatID: chatID, Command: entity.InfoAboutUserFriends, Time: time.Now()})
	return u.userRepo.FriendsByID(userID)
}

func (u *VKUserUsecase) GetInfoAboutGetUserFriends() string {
	return "Это команда выводит 100 друзей публичных пользователей VK"
}