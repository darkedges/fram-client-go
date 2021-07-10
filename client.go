package fram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const HostURL string = "https://fram.darkedges.com"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

// AuthResponse -
type AuthResponse struct {
	TokenId    string `json:"tokenId"`
	SuccessUrl string `json:"successUrl"`
	Realm      string `json:"realm"`
}

// NewClient -
func NewClient(host, username, password *string, realm *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Hashicups URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	if (username != nil) && (password != nil) {

		// authenticate
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/openam/json/realms/root/authenticate", c.HostURL), nil)
		if err != nil {
			return nil, err
		}

		req.Header.Add("X-OpenAM-Password", *username)
		req.Header.Add("X-OpenAM-Username", *password)

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
	req.Header.Set("Authorization", c.Token)

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
