package zoneminder

import (
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	eventsURL     = "/api/events"
	eventCountURL = "/api/events/consoleEvents/"
)

type Events []Event

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

type EventOpts struct {
	MonitorID          int
	Cause              string
	StartTime          time.Time
	StartTimeOperation string
	EndTime            time.Time
	EndTimeOperation   string
}

func (events Events) GetByID(id int) Event {
	for _, e := range events {
		if e.ID == id {
			return e
		}
	}

	return Event{}
}

func (events Events) GetByMonitorID(id int) Events {
	es := make(Events, 0)
	for _, e := range events {
		if e.MonitorID == id {
			es = append(es, e)
		}
	}

	return es
}

func (c *Client) GetEvents(opts *EventOpts) (Events, error) {
	requestURL := eventsURL

	if opts != nil {
		requestURL += "/index"

		if opts.MonitorID != 0 {
			requestURL += fmt.Sprintf("/Monitor:%d", opts.MonitorID)
		}

		if opts.Cause != "" {
			requestURL += fmt.Sprintf("/Cause:%s", opts.Cause)
		}

		if opts.StartTime.String() != "0001-01-01 00:00:00 +0000 UTC" {
			requestURL += fmt.Sprintf("/StartTime%s:%s", opts.StartTimeOperation, opts.StartTime.Format("2006-01-02 15:04:05"))
		}

		if opts.EndTime.String() != "0001-01-01 00:00:00 +0000 UTC" {
			requestURL += fmt.Sprintf("/EndTime%s:%s", opts.EndTimeOperation, opts.EndTime.Format("2006-01-02 15:04:05"))
		}
	}

	requestURL += ".json"

	eventResponse, err := c.httpGet(requestURL, new(events))
	if err != nil {
		return Events{}, err
	}

	e := make(Events, 0)
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

func (c *Client) MonitorForEvents(opts *EventOpts, interval int, eventChan chan Events) {
	c.eventMonitorT = time.NewTicker(time.Duration(interval) * time.Second)

	if opts == nil {
		opts = &EventOpts{}
	}

	opts.StartTimeOperation = ">="

	go func() {
		for range c.eventMonitorT.C {
			opts.StartTime = time.Now().Add(-time.Duration(interval) * time.Second)

			events, err := c.GetEvents(opts)
			if err != nil {
				log.Println("Error getting events.", err)
				continue
			}

			eventChan <- events
		}
	}()
}

func (c *Client) StopEventMonitoring() {
	if c.eventMonitorT != nil {
		c.eventMonitorT.Stop()
	} else {
		log.Println("No event monitor currently running")
	}
}
