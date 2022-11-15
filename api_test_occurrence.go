package trac

import (
	"fmt"
	"net/http"
)

type TestOccurrenceAPI struct {
	HTTPClient *http.Client
}

// GetAllTestOccurrences Get all test occurrences.
func (t TestOccurrenceAPI) GetAllTestOccurrences(testLocator TestOccurrenceLocator, fields string) (value TestOccurrences, err error) {
	req := Request{
		Path:   "/app/rest/testOccurrences",
		Values: Values{"locator": EncodeLocator(testLocator), "fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetTestOccurrence Get a matching test occurrence.
func (t TestOccurrenceAPI) GetTestOccurrence(testLocator TestOccurrenceLocator, fields string) (value TestOccurrence, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/testOccurrences/%s", EncodeLocator(testLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}
