package usecase

import (
	"telegram-bot/entity"
	"telegram-bot/entity/command"
	"telegram-bot/entity/msgtxt"
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
	u.requestRepo.AddInHistory(entity.Request{chatID, command.InfoAboutUser, time.Now()})
	return u.userRepo.User(userID)
}
func (u *VKUserUsecase) GetInfoAboutGetUser() string {
	return msgtxt.InfoAboutGetUser
}

func (u *VKUserUsecase) GetUserFriends(userID int, chatID int) ([]entity.User, error) {
	u.requestRepo.AddInHistory(entity.Request{chatID, command.InfoAboutUserFriends, time.Now()})
	return u.userRepo.FriendsByID(userID)
}

func (u *VKUserUsecase) GetInfoAboutGetUserFriends() string {
	return msgtxt.InfoAboutGetUserFriends
}
