package main

import (
	"runtime"
)

type EventsStorage struct {
	events map[string]string
}

func NewEventsStorage() *EventsStorage {
	return &EventsStorage{
		events: make(map[string]string),
	}
}

func (e *EventsStorage) Event(key string) (value string, ok bool) {
	value, ok = e.events[key]
	return
}

func (e *EventsStorage) AddEvent(key, value string) {
	e.events[key] = value
}

func main() {
	s := NewEventsStorage()
	for i := 0; i < runtime.NumCPU(); i++ {
		if i%2 == 0 {
			go func() {
				for {
					s.Event("Hello")
				}
			}()
		}

		go func() {
			for {
				s.AddEvent("Hello", "World")
			}
		}()
	}
}
