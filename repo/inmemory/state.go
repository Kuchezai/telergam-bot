package inmemory

import (
	"sync"
	"telegram-bot/entity/state"
)

type InMemoryStateStorage struct {
	chatStates map[int]state.State
	mutex      sync.RWMutex
}

func NewInMemoryStateRepo() *InMemoryStateStorage {
	return &InMemoryStateStorage{
		chatStates: make(map[int]state.State),
	}
}

func (r *InMemoryStateStorage) ChatState(chatID int) state.State {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.chatStates[chatID]
}

func (r *InMemoryStateStorage) ChangeChatState(chatID int, newState state.State) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.chatStates[chatID] = newState
	return nil
}
