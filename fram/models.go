package fram

import "net/http"

// Client -
type Client struct {
	HostURL    string
	Realm      string
	HTTPClient *http.Client
	Token      string
}

// AuthResponse -
type AuthResponse struct {
	TokenId    string `json:"tokenId"`
	SuccessUrl string `json:"successUrl"`
	Realm      string `json:"realm"`
}

type BaseURLSource struct {
	Contextpath        string `json:"contextPath,omitempty"`
	FixedValue         string `json:"fixedValue,omitempty"`
	Source             string `json:"source,omitempty"`
	ExtensionClassName string `json:"extensionClassName,omitempty"`
}

type Result struct {
	Contextpath        string     `json:"contextPath"`
	FixedValue         string     `json:"fixedValue"`
	Source             string     `json:"source"`
	ExtensionClassName string     `json:"extensionClassName"`
	Type               ResultType `json:"_type"`
}

type ResultType struct {
	Id         string `json:"_id"`
	Name       string `json:"name"`
	Collection bool   `json:"collection"`
}
