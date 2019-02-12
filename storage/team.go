package storage

import "sync"

// Team in memory storage for teams
type Team struct {
	mu   *sync.RWMutex
	data map[string]struct{}
}

// InitTeam initialize Team
func InitTeam(data map[string]struct{}) Team {
	return Team{mu: &sync.RWMutex{}, data: data}
}

// Remove remove team by key
func (t *Team) Remove(team string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if _, ok := t.data[team]; ok {
		delete(t.data, team)
	}
}

// Find find team by key
func (t *Team) Find(team string) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()

	_, ok := t.data[team]

	return ok
}

// Len length of teams
func (t *Team) Len() int {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return len(t.data)
}
