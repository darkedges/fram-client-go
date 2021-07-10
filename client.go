package fram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const HostURL string = "http://fram.example.com:8080/openam"

// NewClient -
func NewClient(host, username, password *string, realm *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	if (username != nil) && (password != nil) {
		// authenticate
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/json/realms%s/authenticate", c.HostURL, *realm), nil)
		if err != nil {
			return nil, err
		}
		req.Header.Add("X-OpenAM-Password", *password)
		req.Header.Add("X-OpenAM-Username", *username)
		req.Header.Add("Accept-API-Version", "resource=2.1")

		body, err := c.doRequest(req)

		// parse response body
		ar := AuthResponse{}
		err = json.Unmarshal(body, &ar)
		if err != nil {
			return nil, err
		}

		c.Token = ar.TokenId
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Add("Cookie", fmt.Sprintf("iPlanetDirectoryPro=%s", c.Token))
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
