package trac

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"strings"
)

type Mime interface {
	requestCustomizer
	Decode(interface{}, io.Reader) error
	Encode(io.ReadCloser, interface{}) error
}

var (
	ApplicationJson Mime = internalJson{}
	ApplicationXml  Mime = internalXml{}
	TextPlain       Mime = internalTextPlain{}
)

type internalJson struct{}
type internalXml struct{}
type internalTextPlain struct{}

func nopCloserRead(r io.ReadCloser) ([]byte, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	if err = r.Close(); err != nil {
		return nil, err
	}
	r = io.NopCloser(bytes.NewReader(b))
	return b, nil
}

func (c internalJson) Do(r *http.Request) {
	r.Header.Set("Accept", "application/json")
	r.Header.Set("Content-Type", "application/json")
}

func (c internalJson) Decode(value interface{}, r io.Reader) error {
	return json.NewDecoder(r).Decode(value)
}

func (c internalJson) Encode(r io.ReadCloser, value interface{}) error {
	data, err := nopCloserRead(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, value)
}

func (c internalXml) Do(r *http.Request) {
	r.Header.Set("Accept", "application/xml")
	r.Header.Set("Content-Type", "application/xml")
}

func (c internalXml) Decode(v interface{}, r io.Reader) error {
	return xml.NewDecoder(r).Decode(v)
}

func (c internalXml) Encode(r io.ReadCloser, value interface{}) error {
	data, err := nopCloserRead(r)
	if err != nil {
		return err
	}
	return xml.Unmarshal(data, value)
}

func (c internalTextPlain) Do(r *http.Request) {
	r.Header.Set("Accept", "text/plain")
	r.Header.Set("Content-Type", "text/plain")
}

func (c internalTextPlain) Decode(v interface{}, r io.Reader) error {
	r = strings.NewReader(v.(string))
	return nil
}

func (c internalTextPlain) Encode(r io.ReadCloser, value interface{}) error {
	data, err := nopCloserRead(r)
	if err != nil {
		return err
	}
	*value.(*string) = string(data)
	return nil
}
