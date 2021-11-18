package aggregator

import (
	"sync"

	"BigID/domain"
)

type Aggregator interface {
	Aggregate(key string, match domain.Position)
	GetAggregation() map[string][]domain.Position
}

type aggregator struct {
	data map[string][]domain.Position
	lock *sync.Mutex
}

func NewAggregator() *aggregator {
	return &aggregator{lock: &sync.Mutex{}, data: make(map[string][]domain.Position)}
}

func (s *aggregator) Aggregate(key string, match domain.Position) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.data[key] = append(s.data[key], match)
}

func (s *aggregator) GetAggregation() map[string][]domain.Position {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.data
}
