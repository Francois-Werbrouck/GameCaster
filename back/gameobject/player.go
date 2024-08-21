package gameobject

import "GameCaster/main/server/pubsub"

type Player struct {
	ID int
	//User user
	//Perso Perso
	sub pubsub.Subscriber
}

func NewPlayer(id int) {
	// This needs to set the pub-sub system
}

func (p *Player) JoinCombat(combat *Combat) {
	combat.AddParticipant(Participant{})
}
