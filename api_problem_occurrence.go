package trac

import (
	"fmt"
	"net/http"
)

type ProblemOccurrenceAPI struct {
	HTTPClient *http.Client
}

// GetAllBuildProblemOccurrences Get all build problem occurrences.
func (t ProblemOccurrenceAPI) GetAllBuildProblemOccurrences(locator ProblemOccurrenceLocator, fields string) (value ProblemOccurrences, err error) {
	req := Request{
		Path:   "/app/rest/problemOccurrences",
		Values: Values{"locator": EncodeLocator(locator), "fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetBuildProblemOccurrence Get a matching build problem occurrence.
func (t ProblemOccurrenceAPI) GetBuildProblemOccurrence(problemLocator ProblemOccurrenceLocator, fields string) (value ProblemOccurrence, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/problemOccurrences/%s", EncodeLocator(problemLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}
