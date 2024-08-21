package gameobject

import (
	"GameCaster/main/server/pubsub"
	"fmt"
	"time"
)

type GM struct {
	ID   int
	Name string
	Game *Game
	sub  pubsub.Subscriber
}

type Monster struct {
	ID         string
	Archetype  string
	HP         int
	Initiative int
}

type NPC struct {
	ID string
	HP int
}

func NewGM(g *Game) *GM {
	return &GM{
		Name: "dieu",
		Game: g,
		sub:  pubsub.NewSubscriber("DIEU"),
	}
}

func (gm *GM) CreateMonster(archetype string, hp int) Monster {
}

func (gm *GM) GenerateCombat(combatID int) {
	// Create Combat object
	c := NewCombat(combatID)
	c.publisher.AddSubscriber(gm.sub)
	gm.Game.Combat = c
	// Assign monsters to the combat
	// managge the pub sub of the combat and GM and Monster
	gm.Game.initiateCombat()
}

func (gm *GM) StartCombat() {
	// Start the combat
	gm.Game.Combat.StartCombat()
}

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
