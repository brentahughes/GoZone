package zoneminder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (c *Client) httpGet(path string, i interface{}) (interface{}, error) {
	client := &http.Client{Jar: c.Cookies}
	resp, err := client.Get(fmt.Sprintf("%s%s", c.Host, path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(b, &i)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (c *Client) httpPost(path string, postData map[string]string, i interface{}) (interface{}, error) {
	client := &http.Client{Jar: c.Cookies}
	post := make(url.Values)
	for k, v := range postData {
		post.Add(k, v)
	}

	requestURL := fmt.Sprintf("%s%s.json", c.Host, path)

	resp, err := client.PostForm(requestURL, post)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(b, &i)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (c *Client) httpPut(path string, postData map[string]string, i interface{}) (interface{}, error) {
	client := &http.Client{Jar: c.Cookies}
	post := make(url.Values)
	for k, v := range postData {
		post.Add(k, v)
	}

	requestURL := fmt.Sprintf("%s%s.json", c.Host, path)
	req, err := http.NewRequest("PUT", requestURL, bytes.NewBufferString(post.Encode()))
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)

	if i != nil {
		err = json.Unmarshal(b, &i)
		if err != nil {
			return nil, err
		}

		return i, nil
	}

	return b, nil
}
