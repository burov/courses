package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type EventsStorage struct {
	events map[string]string
	sync.RWMutex
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
					s.RLock()
					time.Sleep(100 * time.Millisecond)
					s.Event("Hello")
					fmt.Println("Read Hello")
					s.RUnlock()
				}
			}()
		}

		go func() {
			for {
				s.Lock()
				time.Sleep(1 * time.Second)
				s.AddEvent("Hello", "World")
				fmt.Println("Add Hello")
				s.Unlock()
			}
		}()
	}

	for {
	}
}
