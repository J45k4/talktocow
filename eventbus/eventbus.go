package eventbus

import (
	"sync"

	"github.com/sashabaranov/go-openai"
)

type Event struct {
	ChatGPTRes *openai.ChatCompletionResponse
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
	ch := make(chan Event)
	e.subscribers = append(e.subscribers, ch)
	return ch
}

func (e *Eventbus) Publish(event Event) {
	e.lock.RLock()
	defer e.lock.RUnlock()
	for _, ch := range e.subscribers {
		ch <- event
	}
}
