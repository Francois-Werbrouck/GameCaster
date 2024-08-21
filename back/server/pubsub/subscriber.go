package pubsub

import "fmt"

type ConcreteSubscriber struct {
	id string
}

func NewSubscriber(id string) *ConcreteSubscriber {
	return &ConcreteSubscriber{id: id}
}

func (s *ConcreteSubscriber) Notify(data interface{}) {
	// Handle the received data
	fmt.Printf("Subscriber %s received data: %v\n", s.id, data)
}

func (s *ConcreteSubscriber) GetID() string {
	return s.id
}
