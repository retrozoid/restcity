package trac

/* Golang Teamcity Rest Api Client (TRAC) */

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type requestCustomizer interface {
	Do(r *http.Request)
}

type Authorization interface {
	Prefix() string
	requestCustomizer
}

type Configuration struct {
	ServerURL string
	isXML     bool
	auth      Authorization
}

func NewHTTPAuthConfiguration(serverURL, username, password string) Configuration {
	return Configuration{
		ServerURL: serverURL,
		auth: BasicAuth{
			UserName: username,
			Password: password,
		},
	}
}

func NewTokenAuthConfiguration(serverURL, token string) Configuration {
	return Configuration{
		ServerURL: serverURL,
		auth:      TokenAuth(token),
		isXML:     false,
	}
}

func NewGuestAuthConfiguration(serverURL string) Configuration {
	return Configuration{
		ServerURL: serverURL,
		auth:      GuestAuth{},
	}
}

type GuestAuth struct{}

func (_ GuestAuth) Prefix() string {
	return "/guestAuth"
}

func (_ GuestAuth) Do(r *http.Request) {

}

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

func (b BasicAuth) Prefix() string {
	return "/httpAuth"
}

func (b BasicAuth) Do(r *http.Request) {
	r.SetBasicAuth(b.UserName, b.Password)
}

type TokenAuth string

func (t TokenAuth) Prefix() string {
	return ""
}

func (t TokenAuth) Do(r *http.Request) {
	r.Header.Set("Authorization", "Bearer "+string(t))
}

type ClientImpl struct {
	BuildAPI      BuildAPI
	BuildQueueAPI BuildQueueAPI
	RootAPI       RootAPI
}

type roundTripper struct {
	baseURL string
	auth    Authorization
	retry   int
}

func NewHTTPClient(c Configuration) *http.Client {
	u := strings.TrimRight(c.ServerURL, "/") + c.auth.Prefix()
	return &http.Client{
		Transport: roundTripper{
			baseURL: u,
			auth:    c.auth,
			retry:   5,
		},
	}
}

func New(c Configuration) ClientImpl {
	httpClient := NewHTTPClient(c)
	return ClientImpl{
		BuildAPI:      BuildAPI{HTTPClient: httpClient},
		BuildQueueAPI: BuildQueueAPI{HTTPClient: httpClient},
		RootAPI:       RootAPI{HTTPClient: httpClient},
	}
}

func (r roundTripper) RoundTrip(req *http.Request) (response *http.Response, err error) {
	r.auth.Do(req)

	u, err := url.Parse(r.baseURL + req.URL.Path)
	if err != nil {
		return nil, err
	}
	u.RawQuery = req.URL.RawQuery
	u.RawFragment = req.URL.RawFragment
	req.URL = u

	var buf []byte
	if req.GetBody != nil {
		buf, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
	}
	for i := 0; i < r.retry; i++ {
		if i > 0 {
			time.Sleep(time.Second)
		}
		req.Body = io.NopCloser(bytes.NewReader(buf))
		response, err = http.DefaultTransport.RoundTrip(req)
		if req.Method == http.MethodGet {
			if err != nil {
				continue
			}
			if response != nil {
				switch response.StatusCode {
				case http.StatusBadGateway,
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable,
					http.StatusInternalServerError:
					continue
				}
			}
		}
		return response, err
	}
	return
}

type Values map[string]interface{}

func (v Values) Encode() string {
	var values = url.Values{}
	for k, v := range v {
		str := fmt.Sprint(v)
		if str != "" {
			values.Set(k, str)
		}
	}
	return values.Encode()
}

type Request struct {
	Path     string
	Values   Values
	Data     interface{}
	Consumes Mime
}

func (r Request) Get(client *http.Client, value interface{}) error {
	return r.Do(client, http.MethodGet, value, r.Consumes)
}

func (r Request) Post(client *http.Client, value interface{}) error {
	return r.Do(client, http.MethodPost, value, r.Consumes)
}

func (r Request) Put(client *http.Client, value interface{}) error {
	return r.Do(client, http.MethodPut, value, r.Consumes)
}

func (r Request) Delete(client *http.Client) error {
	return r.Do(client, http.MethodDelete, nil, r.Consumes)
}

func (r Request) Do(client *http.Client, method string, result interface{}, produces Mime) error {
	if r.Consumes == nil {
		r.Consumes = ApplicationJson
	}
	var data io.Reader
	if r.Data != nil {
		if err := r.Consumes.Decode(r.Data, data); err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, r.Path, data)
	if err != nil {
		return err
	}
	r.Consumes.Do(req)
	if len(r.Values) > 0 {
		req.URL.RawQuery = r.Values.Encode()
	}
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		text, _ := nopCloserRead(response.Body)
		return errors.New(string(text))
	}

	if result != nil {
		if produces == nil {
			produces = r.Consumes
		}
		return produces.Encode(response.Body, result)
	}
	return nil
}
