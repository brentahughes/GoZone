package zoneminder

import (
	"fmt"
	"net/http/cookiejar"
	"time"
)

// Version of the package
const Version = 0.7

// Client is the api client used for communicating with the zoneminder API
type Client struct {
	Host          string
	Cookies       *cookiejar.Jar
	eventMonitorT *time.Ticker
}

type Collection interface {
	GetByID(id int)
	GetByName(name string)
}

// NewClient creates the remote client for api communication to zoneminder.
func NewClient(host, username, password string) (*Client, error) {
	cookies, err := login(host, username, password)
	if err != nil {
		return nil, fmt.Errorf("Error creating new zoneminder client. %s", err)
	}

	client := &Client{
		Host:    host,
		Cookies: cookies,
	}
	return client, nil
}
