package eventbus

import (
	"encoding/json"
	"log"
	"sync"
)

type Event struct {
	ChatroomMessage *ChatroomMessage
}

type Eventbus struct {
	subscribers []chan Event
	lock        sync.RWMutex
}

func New() *Eventbus {
	return &Eventbus{
		subscribers: make([]chan Event, 0),
	}
}

func (e *Eventbus) Subscribe() chan Event {
	e.lock.Lock()
	defer e.lock.Unlock()
	ch := make(chan Event, 10)
	e.subscribers = append(e.subscribers, ch)
	return ch
}

func (e *Eventbus) Publish(event Event) {
	res2B, _ := json.Marshal(event)
	log.Printf("Eventbus publish event: %v", string(res2B))

	e.lock.RLock()
	defer e.lock.RUnlock()
	for _, ch := range e.subscribers {
		select {
		case ch <- event:
		default:
		}
	}
}
