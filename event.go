package zoneminder

import (
	"fmt"
	"strings"
	"time"
)

const (
	eventsURL     = "/api/events"
	eventCountURL = "/api/events/consoleEvents/"
)

type events struct {
	Events []eventsPlaceHolder
}

type eventsPlaceHolder struct {
	Event Event
}

type date struct {
	time.Time
}

func (d *date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}

	d.Time = t
	return nil
}

// Event is an event sent back from zoneminder
type Event struct {
	*Client
	ID          int `json:",string"`
	MonitorID   int `json:",string"`
	Name        string
	Cause       string
	StartTime   date
	EndTime     date
	Width       int     `json:",string"`
	Height      int     `json:",string"`
	Length      float64 `json:",string"`
	Frames      int     `json:",string"`
	AlarmFrames int     `json:",string"`
	TotScore    int     `json:",string"`
	AvgScore    int     `json:",string"`
	MaxScore    int     `json:",string"`
	Archived    int     `json:",string"`
	Videoed     int     `json:",string"`
	Uploaded    int     `json:",string"`
	Emailed     int     `json:",string"`
	Message     int     `json:",string"`
	Executed    int     `json:",string"`
	Notes       string
}

func (c *Client) GetEvents() ([]Event, error) {
	eventResponse, err := c.httpGet(fmt.Sprintf("%s.json", eventsURL), new(events))
	if err != nil {
		return []Event{}, err
	}

	e := make([]Event, 0)
	for _, event := range eventResponse.(*events).Events {
		event.Event.Client = c
		e = append(e, event.Event)
	}

	return e, nil
}

func (c *Client) GetEventById(ID int) (Event, error) {
	event, err := c.httpGet(fmt.Sprintf("%s%d.json", eventsURL, ID), new(Event))
	if err != nil {
		return Event{}, err
	}

	return event.(Event), nil
}

func (e Event) Refresh() (Event, error) {
	return e.Client.GetEventById(e.ID)
}
