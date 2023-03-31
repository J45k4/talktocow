package bot

import "github.com/j45k4/talktocow/eventbus"

type CowGPT struct {
	eventbus *eventbus.Eventbus
}

func NewCowGPT(eb *eventbus.Eventbus) *CowGPT {
	return &CowGPT{
		eventbus: eb,
	}
}

func (c *CowGPT) Run() {
	channel := c.eventbus.Subscribe()

	for i := range channel {
		if i.ChatGPTRes != nil {
			// do something with i.ChatGPTRes
		}
	}
}
