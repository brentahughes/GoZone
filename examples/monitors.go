package main

import (
	"log"
	"time"

	zoneminder "github.com/bah2830/GoZone"
)

func main() {
	client, err := zoneminder.NewClient("http://192.168.1.10:8910", "username", "password")
	if err != nil {
		log.Fatal(err)
	}

	monitors, err := client.GetMonitors()
	if err != nil {
		log.Fatal(err)
	}

	eventChan := make(chan zoneminder.Events)
	m := monitors.GetByID(1)
	m.MonitorForEvents("Motion", 5, eventChan)

	for {
		select {
		case e := <-eventChan:
			log.Printf("\n%+v", e)
		default:
			time.Sleep(30 * time.Second)
			m.StopEventMonitoring()
			return
		}
	}

}
