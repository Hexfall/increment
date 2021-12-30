package increment

import (
	"context"
	"sync"
)

type Server struct {
	number int64
	mutex  sync.Mutex
}

func NewServer() Server {
	return Server{
		number: -1,
	}
}

func (s *Server) Increment(ctx context.Context, void *VoidMessage) (*IncrementMessage, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.number++
	return &IncrementMessage{Number: s.number}, nil
}
