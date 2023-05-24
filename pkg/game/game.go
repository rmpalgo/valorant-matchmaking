package game

import (
	"fmt"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (m *Matchmaker) RunQueue(queueName string, playerCount int) {
	for {
		players := make([]Player, playerCount)
		for i := range players {
			player, ok := <-m.Queues[queueName]
			if !ok {
				// The queue has been closed, so exit the loop.
				return
			}
			players[i] = player
		}

		game := Game{
			ID:      fmt.Sprintf("%s-Game-%s", queueName, uuid.NewString()),
			Queue:   queueName,
			Players: players,
		}

		m.Logger.Info("Started a new match",
			zap.String("queue", game.Queue),
			zap.String("game_id", game.ID),
			zap.Int("num_players", len(game.Players)),
			zap.Any("players", game.Players))
	}
}

func (m *Matchmaker) AddPlayer(player Player) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	m.Queues[player.Queue] <- player
}
