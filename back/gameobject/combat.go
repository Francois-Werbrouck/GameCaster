package gameobject

import (
	"GameCaster/main/server/pubsub"
)

type Combat struct {
	ID           int
	Participants []string
	TurnOrder    []int  // Participant IDs in order of initiative
	Turn         int    // Index of the current participant in `TurnOrder`
	State        string // Can be "ongoing", "paused", "completed"
	publisher    pubsub.Publisher
}

func NewCombat(id int) *Combat {
	// Needsw to set the pub sub system
	return &Combat{
		ID:        id,
		State:     "ongoing",
		publisher: pubsub.NewPublisher(),
	}
}

func (c *Combat) AddParticipant() {
	// Add a participant to the combat with their initiative
	// This can be a player or a monster they must join the correct pub sub
}

func (c *Combat) StartCombat() {
	// Roll initiative and sort participants
	c.SortTurnOrder()
}

func (c *Combat) RollInitiatives() {
	// monster are asked to roll initiative
}

func (c *Combat) SortTurnOrder() {
	// Sort participants by initiative
	// Implement a sorting algorithm here to sort `c.Participants` by `Initiative`
	// and update `c.TurnOrder` with their IDs
}

func (c *Combat) SaveState() {
	// Save the combat state to resume later
}
