package game

import (
	"sync"

	"go.uber.org/zap"
)

type Player struct {
	ID    string
	Queue string
}

type Game struct {
	ID      string
	Queue   string
	Players []Player
}

type Matchmaker struct {
	Queues map[string]chan Player
	Mutex  sync.Mutex
	Logger *zap.Logger
}

func NewMatchmaker(logger *zap.Logger) *Matchmaker {
	return &Matchmaker{
		Queues: map[string]chan Player{
			Competitive: make(chan Player, 10),
			Unrated:     make(chan Player, 10),
			Deathmatch:  make(chan Player, 20),
		},
		Logger: logger,
		Mutex:  sync.Mutex{},
	}
}
