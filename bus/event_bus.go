package bus

type Event interface {
	Data() interface{}
	Name() string
}

type Subscriber interface {
	Handle(e Event)
	Supports() Event
}

type EventBus struct {
	subscribers map[string][]Subscriber
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]Subscriber),
	}
}

func (b *EventBus) Register(s ...Subscriber) {
	for _, subscriber := range s {
		event := subscriber.Supports()
		b.subscribers[event.Name()] = append(b.subscribers[event.Name()], subscriber)
	}
}

func (b *EventBus) Dispatch(e Event) {
	for _, subscriber := range b.subscribers[e.Name()] {
		go subscriber.Handle(e)
	}
}
