package trac

import (
	"net/http"
)

type RootAPI struct {
	HTTPClient *http.Client
}

func (t RootAPI) GetRootEndpointsOfRoot() (value string, err error) {
	return value, Request{Path: "/app/rest", Consumer: Coders.String}.Get(t.HTTPClient, &value)
}
