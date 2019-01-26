package storage

import (
	"sort"
	"strings"
	"sync"
)

const (
	teamDelim      = ", "
	playerOptDelim = "; "
)

// Storage memory storage.
type Storage struct {
	mu   *sync.Mutex
	data map[string]Player
}

// Init initialize memory storage.
func Init(data map[string]Player) Storage {
	return Storage{mu: &sync.Mutex{}, data: data}
}

// Upsert insert or update player in memory.
func (p *Storage) Upsert(sp Player) {
	p.mu.Lock()
	defer p.mu.Unlock()

	key := sp.FullName + sp.BirthDate
	player, ok := p.data[key]
	if ok {
		player.Teams = append(player.Teams, sp.Teams...)
		p.data[key] = player
		return
	}

	p.data[key] = sp
}

// SortedByNameList  create sorted by name list of players.
func (p *Storage) SortedByNameList() []Player {
	p.mu.Lock()
	defer p.mu.Unlock()

	res := make([]Player, len(p.data))

	i := 0
	for _, player := range p.data {
		res[i] = player
		i++
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].FullName < res[j].FullName
	})

	return res
}

// Player player instance.
type Player struct {
	FullName  string   `json:"fullName"`
	Age       string   `json:"age"`
	Teams     []string `json:"teams"`
	BirthDate string   `json:"-"`
}

// String return player as string.
func (p Player) String() string {
	return p.FullName + playerOptDelim + p.Age + playerOptDelim + strings.Join(p.Teams, teamDelim)
}
