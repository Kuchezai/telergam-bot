package usecase

import "telegram-bot/entity/state"

type StateRouterUsecase struct {
	stateRepo StateChangerRepo
}

func NewStateRouterUsecase(sr StateChangerRepo) *StateRouterUsecase {
	return &StateRouterUsecase{sr}
}

func (s *StateRouterUsecase) ChangeChatState(chatID int, newState state.State) error {
	return s.stateRepo.ChangeChatState(chatID, newState)
}
