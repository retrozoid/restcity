package trac

/* Golang Teamcity Rest Api Client (TRAC) */

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var defaultCoder = Coders.JSON

type requestCustomizer interface {
	Do(r *http.Request)
}

type Authorization interface {
	Prefix() string
	requestCustomizer
}

type Configuration struct {
	ServerURL string
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
	// if strings.HasPrefix(r.RequestURI,
	r.Header.Set("Authorization", "Bearer "+string(t))
}

type roundTripper struct {
	baseURL string
	auth    Authorization
	retry   int
}

func (c Configuration) NewHTTPClient() *http.Client {
	return &http.Client{
		Transport: roundTripper{
			baseURL: strings.TrimRight(c.ServerURL, "/") + c.auth.Prefix(),
			auth:    c.auth,
			retry:   5,
		},
	}
}

type ClientImpl struct {
	BuildAPI      BuildAPI
	BuildQueueAPI BuildQueueAPI
	BuildTypeAPI  BuildTypeAPI
	RootAPI       RootAPI
	ProjectAPI    ProjectAPI
}

func New(c Configuration) ClientImpl {
	httpClient := c.NewHTTPClient()
	return ClientImpl{
		BuildAPI:      BuildAPI{HTTPClient: httpClient},
		BuildQueueAPI: BuildQueueAPI{HTTPClient: httpClient},
		BuildTypeAPI:  BuildTypeAPI{HTTPClient: httpClient},
		RootAPI:       RootAPI{HTTPClient: httpClient},
		ProjectAPI:    ProjectAPI{HTTPClient: httpClient},
	}
}

func (r roundTripper) RoundTrip(req *http.Request) (response *http.Response, err error) {

	if !req.URL.IsAbs() {
		req.URL, err = url.Parse(r.baseURL + req.URL.RequestURI())
		if err != nil {
			return nil, err
		}
		r.auth.Do(req) // authorization only for baseURL
	}

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
	Consumer Coder
}

func (r Request) Get(client *http.Client, value interface{}) error {
	return r.Do(client, http.MethodGet, value, r.Consumer)
}

func (r Request) Post(client *http.Client, value interface{}) error {
	return r.Do(client, http.MethodPost, value, r.Consumer)
}

func (r Request) Put(client *http.Client, value interface{}) error {
	return r.Do(client, http.MethodPut, value, r.Consumer)
}

func (r Request) Delete(client *http.Client) error {
	return r.Do(client, http.MethodDelete, nil, r.Consumer)
}

func (r Request) Do(client *http.Client, method string, result interface{}, producer Coder) error {
	if r.Consumer == nil {
		r.Consumer = defaultCoder
	}

	var data = new(bytes.Buffer)
	if r.Data != nil {
		if err := r.Consumer.Encode(r.Data, data); err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, r.Path, data)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", r.Consumer.String())
	req.Header.Set("Content-Type", r.Consumer.String())
	if len(r.Values) > 0 {
		req.URL.RawQuery = r.Values.Encode()
	}

	response, err := client.Do(req)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		e := APIError{
			RequestURL: req.RequestURI,
			Method:     req.Method,
			StatusCode: response.StatusCode,
		}
		err = Coders.String.Decode(response.Body, &e.Message)
		if err != nil {
			e.Message = err.Error()
		}
		return e
	}

	if result != nil {
		if producer == nil {
			producer = r.Consumer
		}
		return producer.Decode(response.Body, result)
	}
	return nil
}

type APIError struct {
	Method     string
	RequestURL string
	StatusCode int
	Message    string
}

func (e APIError) Error() string {
	return e.Message
}

func IsHTTPStatusCode(err error, statusCode int) bool {
	if err == nil {
		return false
	}
	e, ok := err.(APIError)
	return ok && e.StatusCode == statusCode
}
