package pubsub

import "sync"

type ConcretePublisher struct {
	subscribers map[string]Subscriber
	mu          sync.Mutex
}

func NewPublisher() *ConcretePublisher {
	return &ConcretePublisher{
		subscribers: make(map[string]Subscriber),
	}
}

func (p *ConcretePublisher) AddSubscriber(sub Subscriber) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.subscribers[sub.GetID()] = sub
}

func (p *ConcretePublisher) RemoveSubscriber(sub Subscriber) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.subscribers, sub.GetID())
}

func (p *ConcretePublisher) NotifySubscribers(data interface{}) {
	p.mu.Lock()
	defer p.mu.Unlock()
	for _, sub := range p.subscribers {
		sub.Notify(data)
	}
}
