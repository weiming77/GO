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

// A Thermostat measures and controls the temperature
// PS: value is not exported
type Thermostat struct {
	id    string
	value float64
}

// value return the current temperature in Celsius
func (t *Thermostat) Value() float64 {
	return t.value
}

// set tells the thermostat to set the temperature
func (t *Thermostat) Set(value float64) {
	t.value = value
}

func (t *Thermostat) ID() string {
	return t.id
}

// kind returns the device kind
func (*Thermostat) Kind() string {
	return "Thermostat"
}

// Camera is a security camera
type Camera struct {
	id string
}

// ID return the camera ID
func (c *Camera) ID() string {
	return c.id
}

func (*Camera) Kind() string {
	return "Camera"
}

// Interface in GO are small and defined at the point of using them
type Sensor interface {
	ID() string
	Kind() string
}

func printAll(sensor []Sensor) {
	for _, s := range sensor {
		fmt.Printf("%s <%s>\n", s.ID(), s.Kind())
	}
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

	// test Go methods
	t := Thermostat{"Living Room", 16.2}
	fmt.Printf("%s before: %.2f\n", t.ID, t.Value())

	t.Set(18)
	fmt.Printf("%s after: %.2f\n", t.ID, t.Value())

	c := Camera{"Baby room"}

	sensors := []Sensor{&t, &c}
	printAll(sensors)
}
