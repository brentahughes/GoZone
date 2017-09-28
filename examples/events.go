package main

import (
	"fmt"
	"log"

	zoneminder "github.com/bah2830/GoZone"
)

func main() {
	client, err := zoneminder.NewClient("http://192.168.1.10:8910", "username", "password")
	if err != nil {
		log.Fatal(err)
	}

	events, err := client.GetEvents(&zoneminder.EventOpts{Cause: "Motion"})
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range events {
		fmt.Printf("%s - %s: %s\n", e.StartTime.String(), e.Name, e.Cause)
	}

}
