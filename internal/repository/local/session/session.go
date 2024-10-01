package session

import (
	"context"
	"sync"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/rand"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/resterr"
)

type Storage struct {
	mu       sync.RWMutex
	sessions map[domain.SessionID]domain.UserID
}

func NewStorage() *Storage {
	return &Storage{sessions: make(map[domain.SessionID]domain.UserID)}
}

func (s *Storage) Create(ctx context.Context, userID domain.UserID) (domain.SessionID, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	randS, err := rand.String(64)
	if err != nil {
		return "", resterr.NewInternalServerError(err)
	}

	sid := domain.SessionID(randS)
	s.sessions[sid] = userID

	return sid, nil
}

func (s *Storage) Delete(ctx context.Context, sessionID domain.SessionID) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.sessions, sessionID)

	return nil
}

func (s *Storage) GetUserID(ctx context.Context, sessionID domain.SessionID) (domain.UserID, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	userID, ok := s.sessions[sessionID]
	if !ok {
		return 0, resterr.NewNotFoundError("session not found")
	}

	return userID, nil
}
