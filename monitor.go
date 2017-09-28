package zoneminder

import (
	"fmt"
	"log"
	"reflect"
	"time"
)

const monitorURL = "/api/monitors"

type Monitors []Monitor

type monitors struct {
	Monitors []monitorPlaceHolder
}

type monitorPlaceHolder struct {
	Monitor Monitor
}

// Monitor is the monitor object from zoneminder
type Monitor struct {
	*Client
	ID                  int `json:",string"`
	Name                string
	ServerID            int `json:",string"`
	Type                string
	Function            string
	Enabled             int `json:",string"`
	LinkedMonitors      string
	Triggers            string
	Device              string
	Channel             int `json:",string"`
	Format              int `json:",string"`
	V4LMultiBuffer      bool
	V4LCapturesPerFrame int `json:",string"`
	Protocol            string
	Method              string
	Host                string
	Port                int `json:",string"`
	SubPath             string
	Path                string
	Options             string
	User                string
	Pass                string
	Width               int `json:",string"`
	Height              int `json:",string"`
	Colours             int `json:",string"`
	Palette             int `json:",string"`
	Orientation         int `json:",string"`
	Deinterlacing       int `json:",string"`
	RTSPDescribe        bool
	Brightness          int `json:",string"`
	Contrast            int `json:",string"`
	Hue                 int `json:",string"`
	Colour              int `json:",string"`
	EventPrefix         string
	LabelFormat         string
	LabelX              int     `json:",string"`
	LabelY              int     `json:",string"`
	LabelSize           int     `json:",string"`
	ImageBufferCount    int     `json:",string"`
	WarmupCount         int     `json:",string"`
	PreEventCount       int     `json:",string"`
	PostEventCount      int     `json:",string"`
	StreamReplayBuffer  int     `json:",string"`
	AlarmFrameCount     int     `json:",string"`
	SectionLength       int     `json:",string"`
	FrameSkip           int     `json:",string"`
	MotionFrameSkip     int     `json:",string"`
	AnalysisFPS         float64 `json:",string"`
	AnalysisUpdateDelay int     `json:",string"`
	MaxFPS              float64 `json:",string"`
	AlarmMaxFPS         float64 `json:",string"`
	FPSReportInterval   int     `json:",string"`
	RefBlendPerc        int     `json:",string"`
	AlarmRefBlendPerc   int     `json:",string"`
	Controllable        int     `json:",string"`
	ControlID           int     `json:",string"`
	ControlDevice       string
	ControlAddress      string
	AutoStopTimeout     string
	TrackMotion         int `json:",string"`
	TrackDelay          int `json:",string"`
	ReturnLocation      int `json:",string"`
	ReturnDelay         int `json:",string"`
	DefaultView         string
	DefaultRate         int `json:",string"`
	DefaultScale        int `json:",string"`
	SignalCheckColour   string
	WebColour           string
	Exif                bool
	Sequence            int `json:",string"`
	eventMonitorT       *time.Ticker
}

type MonitorOpts struct {
	Name     string
	Function string
	Protocol string
	Method   string
	Host     string
	Port     int
	Path     string
	Width    int
	Height   int
	Colours  int
}

func (m *Monitors) GetByID(id int) Monitor {
	for _, mon := range *m {
		if mon.ID == id {
			return mon
		}
	}

	return Monitor{}
}

func (m *Monitors) GetByName(name string) Monitor {
	for _, mon := range *m {
		if mon.Name == name {
			return mon
		}
	}

	return Monitor{}
}

func (c *Client) GetMonitors() (Monitors, error) {
	monitorResponse, err := c.httpGet(fmt.Sprintf("%s.json", monitorURL), new(monitors))
	if err != nil {
		return Monitors{}, err
	}

	m := make(Monitors, len(monitorResponse.(*monitors).Monitors))
	for _, monitor := range monitorResponse.(*monitors).Monitors {
		monitor.Monitor.Client = c
		m = append(m, monitor.Monitor)
	}

	return m, nil
}

func (c *Client) GetMonitorById(ID int) (Monitor, error) {
	monitor, err := c.httpGet(fmt.Sprintf("%s/%d.json", monitorURL, ID), new(Monitor))
	if err != nil {
		return Monitor{}, err
	}

	return monitor.(Monitor), nil
}

func (m *Monitor) Refresh() (Monitor, error) {
	return m.Client.GetMonitorById(m.ID)
}

func (m *Monitor) Enable() error {
	postData := map[string]string{
		"Monitor[Enabled]": "1",
	}

	_, err := m.Client.httpPost(fmt.Sprintf("%s/%d.json", monitorURL, m.ID), postData, nil)
	return err
}

func (m *Monitor) Disable() error {
	postData := map[string]string{
		"Monitor[Enabled]": "0",
	}

	_, err := m.Client.httpPost(fmt.Sprintf("%s/%d.json", monitorURL, m.ID), postData, nil)
	return err
}

func (m *Monitor) SetFunction(f string) error {
	postData := map[string]string{
		"Monitor[Function]": f,
	}

	_, err := m.Client.httpPost(fmt.Sprintf("%s/%d.json", monitorURL, m.ID), postData, nil)
	return err
}

func (c *Client) AddMonitor(opts MonitorOpts) (Monitor, error) {
	postData := make(map[string]string)

	v := reflect.ValueOf(opts).Elem()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		if f.Interface() == nil {
			continue
		}

		postData["Monitor["+v.Type().Field(i).Name+"]"] = f.Interface().(string)
	}

	_, err := c.httpPost(fmt.Sprintf("%s.json", monitorURL), postData, nil)

	return Monitor{}, err
}

func (m *Monitor) Edit(opts MonitorOpts) error {
	postData := make(map[string]string)

	v := reflect.ValueOf(opts).Elem()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		if f.Interface() == nil {
			continue
		}

		postData["Monitor["+v.Type().Field(i).Name+"]"] = f.Interface().(string)
	}

	_, err := m.Client.httpPut(fmt.Sprintf("%s/%d.json", monitorURL, m.ID), postData, nil)
	return err
}

func (m *Monitor) ForceAlarm() error {
	_, err := m.Client.httpGet(fmt.Sprintf("%s/alarm/id:%d/command:on.json", monitorURL, m.ID), nil)
	return err
}

func (m *Monitor) StopAlarm() error {
	_, err := m.Client.httpGet(fmt.Sprintf("%s/alarm/id:%d/command:off.json", monitorURL, m.ID), nil)
	return err
}

func (m *Monitor) AlarmStatus() (int, error) {
	type status struct {
		Status int `json:"status,string"`
	}
	s, err := m.Client.httpGet(fmt.Sprintf("%s/alarm/id:%d/command:status.json", monitorURL, m.ID), new(status))
	return s.(*status).Status, err
}

func (m *Monitor) GetEvents(opts *EventOpts) (Events, error) {
	if opts == nil {
		opts = &EventOpts{}
	}

	opts.MonitorID = m.ID
	return m.Client.GetEvents(opts)
}

func (m *Monitor) MonitorForEvents(eventType string, interval int, eventChan chan Events) {
	m.eventMonitorT = time.NewTicker(time.Duration(interval) * time.Second)
	go func() {
		for range m.eventMonitorT.C {
			events, err := m.GetEvents(&EventOpts{
				Cause:              eventType,
				StartTime:          time.Now().Add(-time.Duration(interval) * time.Second),
				StartTimeOperation: ">=",
			})
			if err != nil {
				log.Println("Error getting events for monitor.", err)
				continue
			}

			eventChan <- events
		}
	}()
}

func (m *Monitor) StopEventMonitoring() {
	if m.eventMonitorT != nil {
		m.eventMonitorT.Stop()
	} else {
		log.Println("No event monitor currently running")
	}
}
