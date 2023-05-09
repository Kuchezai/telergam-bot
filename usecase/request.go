package usecase

import (
	"telegram-bot/entity"
	"telegram-bot/entity/command"
	"time"
)

type RequestUsecase struct {
	requestRepo RequestRepo
	stateRepo   StateChangerRepo
}

func NewRequestUsecase(rr RequestRepo, sr StateChangerRepo) *RequestUsecase {
	return &RequestUsecase{rr, sr}
}

func (r *RequestUsecase) GetChatRequests(chatID int) ([]entity.Request, error) {
	r.requestRepo.AddInHistory(entity.Request{chatID, command.RequestHistory, time.Now()})
	return r.requestRepo.GetChatHistory(chatID)
}

func (r *RequestUsecase) GetInfoAboutGetChatRequests() string {
	return "Это команда показывает Вашу историю запросов к этому боту"
}
