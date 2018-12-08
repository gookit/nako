package internal

import "strings"

type EventHandler func(...interface{}) error

// SimpleEvent
type SimpleEvent struct {
	events map[string][]EventHandler
}

// NewSimpleEvent create
func NewSimpleEvent() *SimpleEvent {
	return &SimpleEvent{
		events: map[string][]EventHandler{},
	}
}

// On register a event handler
func (se *SimpleEvent) On(name string, handler EventHandler) {
	name = strings.TrimSpace(name)

	if name == "" {
		panic("event name cannot be empty")
	}

	if ls, ok := se.events[name]; ok {
		se.events[name] = append(ls, handler)
	} else {
		se.events[name] = []EventHandler{handler}
	}
}

// MustFire fire event by name
func (se *SimpleEvent) MustFire(name string, args ...interface{}) {
	err := se.Fire(name, args...)
	if err != nil {
		panic(err)
	}
}

// Fire event by name
func (se *SimpleEvent) Fire(name string, args ...interface{}) (err error) {
	ls, ok := se.events[name]
	if !ok {
		return
	}

	for _, fn := range ls {
		if err = fn(args...); err != nil {
			return
		}
	}

	return
}

// Clear all events info.
func (se *SimpleEvent) Clear() {
	se.events = map[string][]EventHandler{}
}
