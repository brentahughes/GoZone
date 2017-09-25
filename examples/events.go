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

	events, err := client.GetEvents()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%+v\n", events)
}
