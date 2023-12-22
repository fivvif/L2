package main

import (
	"encoding/json"
	"time"
)

type Event struct {
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

type Calendar struct {
	Events []Event
}

func NewCalendar() *Calendar {
	return &Calendar{}
}

func (c *Calendar) NewEvent(event Event) {
	c.Events = append(c.Events, event)
}

func (c *Calendar) GetEvents() []Event {
	return c.Events
}

func serializeCalendar(calendar *Calendar) ([]byte, error) {
	return json.Marshal(calendar)
}

func unserializeCalendar(data []byte) (Calendar, error) {
	var calendar Calendar
	err := json.Unmarshal(data, &calendar)
	if err != nil {
		return Calendar{}, err
	}
	return calendar, nil
}
