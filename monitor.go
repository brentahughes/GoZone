package zoneminder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const monitorURL = "/api/monitors.json"

type monitors struct {
	Monitors []monitorPlaceHolder
}

type monitorPlaceHolder struct {
	Monitor Monitor
}

// Monitor is the monitor object from zoneminder
type Monitor struct {
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
}

func getMonitors(c *Client) ([]Monitor, error) {
	client := &http.Client{Jar: c.Cookies}
	resp, err := client.Get(fmt.Sprintf("%s%s", c.Host, monitorURL))
	if err != nil {
		return []Monitor{}, err
	}

	b, _ := ioutil.ReadAll(resp.Body)
	var monitorResponse monitors
	err = json.Unmarshal(b, &monitorResponse)
	if err != nil {
		return []Monitor{}, err
	}

	monitors := make([]Monitor, 0)

	for _, monitor := range monitorResponse.Monitors {
		monitors = append(monitors, monitor.Monitor)
	}

	return monitors, nil
}
