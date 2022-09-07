package trac

import (
	"encoding/json"
	"encoding/xml"
	"io"
)

type Coder interface {
	String() string
	Encode(interface{}, io.Writer) error // writes Object values to an output stream.
	Decode(io.Reader, interface{}) error // reads and decodes Object values from an input stream.
}

var Coders = struct {
	JSON   Coder
	XML    Coder
	String Coder
	Stream Coder
}{
	JSON:   jsonCoder{},
	XML:    xmlCoder{},
	String: stringCoder{},
	Stream: streamCoder{},
}

type jsonCoder struct{}
type xmlCoder struct{}
type stringCoder struct{}
type streamCoder struct {
	stringCoder
}

func (c jsonCoder) String() string {
	return "application/json"
}

func (c jsonCoder) Decode(reader io.Reader, value interface{}) error {
	return json.NewDecoder(reader).Decode(value)
}

func (c jsonCoder) Encode(value interface{}, w io.Writer) error {
	return json.NewEncoder(w).Encode(value)
}

func (c xmlCoder) String() string {
	return "application/xml"
}

func (c xmlCoder) Decode(r io.Reader, v interface{}) error {
	return xml.NewDecoder(r).Decode(v)
}

func (c xmlCoder) Encode(v interface{}, w io.Writer) error {
	return xml.NewEncoder(w).Encode(v)
}

func (c stringCoder) String() string {
	return "text/plain"
}

func (c streamCoder) String() string {
	return "application/octet-stream"
}

func (c stringCoder) Decode(r io.Reader, v interface{}) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	*v.(*string) = string(b)
	return nil
}

func (c stringCoder) Encode(v interface{}, w io.Writer) error {
	_, err := w.Write([]byte(v.(string)))
	return err
}
