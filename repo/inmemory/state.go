package inmemory

import (
	"sync"
	"telegram-bot/entity"
)

type InMemoryStateStorage struct {
	chatStates map[int]entity.State
	mutex      sync.RWMutex
}

func NewInMemoryStateRepo() *InMemoryStateStorage {
	return &InMemoryStateStorage{
		chatStates: make(map[int]entity.State),
	}
}

func (r *InMemoryStateStorage) ChatState(chatID int) entity.State {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.chatStates[chatID]
}

func (r *InMemoryStateStorage) ChangeChatState(chatID int, newState entity.State) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.chatStates[chatID] = newState
	return nil
}
