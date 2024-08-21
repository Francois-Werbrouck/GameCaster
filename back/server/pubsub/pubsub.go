package pubsub

type Subscriber interface {
	Notify(data interface{})
	GetID() string
}

type Publisher interface {
	AddSubscriber(sub Subscriber)
	RemoveSubscriber(sub Subscriber)
	NotifySubscribers(data interface{})
}
