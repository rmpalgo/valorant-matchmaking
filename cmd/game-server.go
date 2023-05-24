package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/rmpalgo/valorant-multiplayer/pkg/game"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Unable to initialize logger: %v", err)
	}

	defer func() {
		if err := logger.Sync(); err != nil {
			log.Fatalf("Cannot sync logger: %v", err)
		}
	}()

	m := game.NewMatchmaker(logger)

	var wg sync.WaitGroup

	// Start the matchmaking goroutines.
	wg.Add(3)

	go func() {
		defer wg.Done()
		m.RunQueue(game.Competitive, 10)
	}()

	go func() {
		defer wg.Done()
		m.RunQueue(game.Unrated, 10)
	}()

	go func() {
		defer wg.Done()
		m.RunQueue(game.Deathmatch, 20)
	}()

	// Simulate adding players to the queues.
	for i := 1; i <= 60; i++ {
		player := game.Player{ID: fmt.Sprintf("Player%d", i)}

		switch {
		case i%3 == 0:
			player.Queue = game.Deathmatch
		case i%3 == 1:
			player.Queue = game.Unrated
		case i%3 == 2:
			player.Queue = game.Competitive
		}

		m.AddPlayer(player)
	}

	m.Mutex.Lock()
	for _, queue := range m.Queues {
		close(queue)
	}
	m.Mutex.Unlock()

	// Wait for matches to be formed before ending the program.
	wg.Wait()
}
