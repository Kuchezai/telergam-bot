package usecase

import (
	"telegram-bot/entity"
	"telegram-bot/entity/command"
	"time"
)

type AboutBotUsecase struct {
	requestRepo RequestsAdderRepo
	stateRepo   StateChangerRepo
}

func NewAboutBotUsecase(rp RequestsAdderRepo, sr StateChangerRepo) *AboutBotUsecase {
	return &AboutBotUsecase{rp, sr}
}

func (a *AboutBotUsecase) GetInfoAboutAuthor(chatID int) (string, error) {
	a.requestRepo.AddInHistory(entity.Request{chatID, command.InfoAboutAuthor, time.Now()})
	return "Kuchezai", nil
}

func (a *AboutBotUsecase) GetAuthorGitHub(chatID int) (string, error) {
	return "https://github.com/Kuchezai", nil
}
