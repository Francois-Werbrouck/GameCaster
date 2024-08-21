package gameobject

import (
	"GameCaster/main/server/pubsub"
)

type Game struct {
	ID        int
	Combat    *Combat
	Players   []Player
	Caster    *Caster
	publisher pubsub.Publisher
}

func NewGame(id int) *Game {

	return &Game{
		ID:        id,
		publisher: pubsub.NewPublisher(),
	}
}

func (g *Game) AddUser(role string) {
	// Add a player to the game
	// This can be a player or a monster they must join the correct pub sub

	if role == "GM" {
		game_master := NewGM(g)
		g.publisher.AddSubscriber(game_master.sub)
	} else if role == "Player" {
		g.Players = append(g.Players, NewPlayer())
	} else if role == "Caster" {
		g.Caster = NewCaster()
	}
}

// GenerateCombat creates a new combat
func (g *Game) initiateCombat() {
	// Monster Initiative is rolled
	g.Combat.RollInitiatives()
	// Participants needs to be notified to join
}
