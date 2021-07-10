package fram

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) CreateBaseURLSource(bus BaseURLSource) (*Result, error) {
	rb, err := json.Marshal(bus)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/json/realm-config/services/baseurl", c.HostURL), strings.NewReader(string(rb)))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Requested-With", "SwaggerUI")

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	result := Result{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetBaseURLSource() (*BaseURLSource, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/json/realm-config/services/baseurl", c.HostURL), nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Requested-With", "SwaggerUI")

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	result := BaseURLSource{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteBaseURLSource() (*BaseURLSource, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/json/realm-config/services/baseurl", c.HostURL), nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Requested-With", "SwaggerUI")

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	result := BaseURLSource{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) BaseURLSourceId() string {
	id := fmt.Sprintf("%s/json/realm-config/services/baseurl", c.HostURL)
	return id
}
