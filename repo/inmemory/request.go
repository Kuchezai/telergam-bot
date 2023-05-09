package inmemory

import (
	"fmt"
	"sync"
	"telegram-bot/entity"
)

type InMemoryRequestStorage struct {
	sync.RWMutex
	data map[int]*[]entity.Request
}

func NewInMemoryRequestRepo() *InMemoryRequestStorage {
	storage := InMemoryRequestStorage{data: make(map[int]*[]entity.Request)}
	return &storage
}

func (s *InMemoryRequestStorage) AddInHistory(req entity.Request) error {
	s.Lock()
	defer s.Unlock()
	id := req.ChatID
	if _, ok := s.data[id]; !ok {
		slice := make([]entity.Request, 0)
		s.data[id] = &slice
	}
	*s.data[id] = append(*s.data[id], req)
	return nil
}

func (s *InMemoryRequestStorage) GetChatHistory(chatID int) ([]entity.Request, error) {
	s.RLock()
	defer s.RUnlock()
	if _, ok := s.data[chatID]; !ok {
		return nil, fmt.Errorf("chat history not found")
	}
	return *s.data[chatID], nil

}
