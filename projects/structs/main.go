//When you need to define your own data types, we will use Structs

package main

import (
	"fmt"
	"log"
	"time"
)

type Event struct {
	ID   string
	time time.Time
}

// Event specific
//PS: Use capital therefore it could be exported and accessible from event
type DoorEvent struct {
	Event
	Action string // open, close
}

// Event specific
//PS: Use capital therefore it could be exported and accessible from event
type TemperatureEvent struct {
	Event
	Value float64
}

func NewDoorEvent(id string, time time.Time, action string) (*DoorEvent, error) {
	if id == "" {
		return nil, fmt.Errorf("empty ID")
	}

	evt := DoorEvent{
		Event:  Event{id, time},
		Action: action,
	}
	// return the pointer to the event
	return &evt, nil
}

func main() {
	evt, err := NewDoorEvent("Open door", time.Now(), "Open")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", evt)

}
