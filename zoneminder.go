package zoneminder

import (
	"fmt"
	"net/http/cookiejar"
)

// Version of the package
const Version = 0.1

// Client is the api client used for communicating with the zoneminder API
type Client struct {
	Host    string
	Cookies *cookiejar.Jar
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

// GetMonitors returns a list of monitors from zoneminder
func (c *Client) GetMonitors() ([]Monitor, error) {
	return getMonitors(c)
}

// GetEvents returns a list of events from zoneminder
func (c *Client) GetEvents() ([]Event, error) {
	return getEvents(c)
}
