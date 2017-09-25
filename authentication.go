package zoneminder

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const loginURL = "/index.php"

func login(host, username, password string) (*cookiejar.Jar, error) {
	cookies, err := cookiejar.New(nil)
	if err != nil {
		log.Printf("Error creating cookie jar for zoneminder. %s", err)
	}

	requestURL := fmt.Sprintf("%s%s", host, loginURL)
	postData := url.Values{
		"username": {username},
		"password": {password},
		"action":   {"login"},
		"view":     {"console"},
	}

	resp, requestError := http.PostForm(requestURL, postData)
	if requestError != nil {
		return cookies, fmt.Errorf("Error logging into zineminder. %s", requestError)
	}
	defer resp.Body.Close()

	u, urlParseErr := url.Parse(requestURL)
	if urlParseErr != nil {
		return cookies, fmt.Errorf("Error parsing login info from zoneminder. %s", urlParseErr)
	}

	cookies.SetCookies(u, resp.Cookies())

	return cookies, nil
}
