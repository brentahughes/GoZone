package zoneminder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const eventsURL = "/api/events.json"
const eventCountURL = "/api/events/consoleEvents/"

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

func getEvents(c *Client) ([]Event, error) {
	client := &http.Client{Jar: c.Cookies}
	resp, err := client.Get(fmt.Sprintf("%s%s", c.Host, eventsURL))
	if err != nil {
		return []Event{}, err
	}

	b, _ := ioutil.ReadAll(resp.Body)
	var eventresponse events
	err = json.Unmarshal(b, &eventresponse)
	if err != nil {
		return []Event{}, err
	}

	events := make([]Event, 0)
	for _, event := range eventresponse.Events {
		events = append(events, event.Event)
	}

	return events, nil
}
