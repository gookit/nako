package internal

import (
	"strings"
	"sync"
)

const wildcard = "*"

// EventHandler func define
type EventHandler func(e *EventData) error

/*************************************************************
 * Event Data
 *************************************************************/

// EventData struct
type EventData struct {
	abort bool
	// event name
	name string
	// user data.
	Data []interface{}
}

// Name get
func (e *EventData) Name() string {
	return e.name
}

// Aborted exec.
func (e *EventData) Aborted() {
	e.abort = true
}

func (e *EventData) init(s string, args []interface{}) {
	e.name = s
	e.Data = args
}

func (e *EventData) reset() {
	e.name = ""
	e.Data = make([]interface{}, 0)
}

/*************************************************************
 * Event Manager
 *************************************************************/

// EventManager struct
type EventManager struct {
	names map[string]int
	events map[string][]EventHandler
	pool sync.Pool
}

// NewEventManager create EventManager instance
func NewEventManager() *EventManager {
	em := &EventManager{
		names: make(map[string]int),
		events: make(map[string][]EventHandler),
	}

	// set pool creator
	em.pool.New = func() interface{} {
		return &EventData{}
	}

	return em
}

// On register a event handler
func (em *EventManager) On(name string, handler EventHandler) {
	name = strings.TrimSpace(name)
	if name == "" {
		panic("event name cannot be empty")
	}

	if ls, ok := em.events[name]; ok {
		em.events[name] = append(ls, handler)
	} else { // first add.
		em.events[name] = []EventHandler{handler}
	}
}

// MustFire fire event by name
func (em *EventManager) MustFire(name string, args ...interface{}) {
	err := em.Fire(name, args...)
	if err != nil {
		panic(err)
	}
}

// Fire event by name
func (em *EventManager) Fire(name string, args ...interface{}) (err error) {
	ls, ok := em.events[name]
	if !ok {
		return
	}

	e := em.pool.Get().(*EventData)
	e.init(name, args)

	defer func() {
		e.reset()
	}()

	for _, fn := range ls {
		if err = fn(e); err != nil {
			return
		}
	}

	if wildcard {
		
	}



	return
}

// EventsByName get events and handlers by name
func (em *EventManager) EventsByName(name string) (es []EventHandler) {
	es, _ = em.events[name]
	return
}

// Events get all events and handlers
func (em *EventManager) Events() map[string][]EventHandler {
	return em.events
}

// Names get all event names
func (em *EventManager) Names() map[string]int {
	return em.names
}

// Clear all events info.
func (em *EventManager) Clear() {
	em.events = map[string][]EventHandler{}
}
