package main

import (
	"fmt"
	"runtime"
	"sync"
)

type EventsStorage struct {
	events map[string]string
	sync.Mutex
}

func NewEventsStorage() *EventsStorage {
	return &EventsStorage{
		events: make(map[string]string),
	}
}

func (e *EventsStorage) Event(key string) (value string, ok bool) {
	e.Lock()
	defer e.Unlock()
	value, ok = e.events[key]
	return
}

func (e *EventsStorage) AddEvent(key, value string) {
	e.Lock()
	e.events[key] = value
	e.Unlock()
}

func main() {
	s := NewEventsStorage()
	for i := 0; i < runtime.NumCPU(); i++ {
		if i%2 == 0 {
			go func() {
				for {
					s.Event("Hello")
					fmt.Println("get Hello")
				}
			}()
		}

		go func() {
			for {
				s.AddEvent("Hello", "World")
				fmt.Println("save Hello World")
			}
		}()
	}

	for {
	}
}
