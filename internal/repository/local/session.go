package local

import (
	"context"
	"sync"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/interr"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/rand"
)

const SessionIDLength = 64

type SessionRepository struct {
	mu       sync.RWMutex
	sessions map[domain.SessionID]domain.UserID
}

func NewSessionRepository() *SessionRepository {
	return &SessionRepository{sessions: make(map[domain.SessionID]domain.UserID)}
}

func (r *SessionRepository) CreateSession(ctx context.Context, userID domain.UserID) (domain.SessionID, error) {
	randS, err := rand.String(SessionIDLength)
	if err != nil {
		return "", interr.NewInternalError(err, "session SessionRepository.CreateUser rand.String")
	}

	sid := domain.SessionID(randS)

	r.mu.Lock()
	defer r.mu.Unlock()
	r.sessions[sid] = userID

	return sid, nil
}

func (r *SessionRepository) DeleteSession(ctx context.Context, sessionID domain.SessionID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.sessions, sessionID)

	return nil
}

func (r *SessionRepository) GetUserID(ctx context.Context, sessionID domain.SessionID) (domain.UserID, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	userID, ok := r.sessions[sessionID]
	if !ok {
		return 0, interr.NewNotFoundError("session SessionRepository.GetUserID r.sessions")
	}

	return userID, nil
}
