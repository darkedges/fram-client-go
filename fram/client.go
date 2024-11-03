package fram

import (
	"fmt"
	"github.com/darkedges/go-frodo-lib"
	"io"
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
	if realm != nil {
		c.Realm = *realm
	}

	if (username != nil) && (password != nil) {
		f, _ := frodo.CreateInstanceWithAdminAccount(frodo.Params{
			Host:  *host,
			User:  *username,
			Pass:  *password,
			Realm: "/root",
		})
		f.Login()
		info := f.GetInfo()
		c.Token = info.SessionToken
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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
