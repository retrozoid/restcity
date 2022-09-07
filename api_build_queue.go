package trac

import (
	"fmt"
	"net/http"
)

type BuildQueueAPI struct {
	HTTPClient *http.Client
}

// DeleteAllQueuedBuilds Delete all queued builds.
func (t BuildQueueAPI) DeleteAllQueuedBuilds(buildQueueLocator BuildQueueLocator, fields string) (err error) {
	req := Request{
		Path: "/app/rest/buildQueue",
		Values: Values{
			"locator": locatorString(buildQueueLocator),
			"fields":  fields,
		},
	}
	err = req.Delete(t.HTTPClient)
	return
}

// GetAllQueuedBuilds Get all queued builds.
func (t BuildQueueAPI) GetAllQueuedBuilds(buildQueueLocator BuildQueueLocator, fields string) (value Builds, err error) {
	req := Request{
		Path: "/app/rest/buildQueue",
		Values: Values{
			"locator": locatorString(buildQueueLocator),
			"fields":  fields,
		},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// AddBuildToQueue Add a new build to the queue.
func (t BuildQueueAPI) AddBuildToQueue(data Build, moveToTop bool) (value Build, err error) {
	req := Request{
		Path: "/app/rest/buildQueue",
		Values: Values{
			"moveToTop": moveToTop,
		},
		Data: data,
	}
	return value, req.Post(t.HTTPClient, &value)
}

// SetQueuedBuildsOrder Update the build queue order.
func (t BuildQueueAPI) SetQueuedBuildsOrder(data Builds, fields string) (value Builds, err error) {
	req := Request{
		Path:   "/app/rest/buildQueue/order",
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetQueuedBuildPosition Get the queue position of a queued matching build.
func (t BuildQueueAPI) GetQueuedBuildPosition(queuePosition string, fields string) (value Build, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildQueue/order/%s", queuePosition),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// SetQueuedBuildPosition Update the queue position of a queued matching build.
func (t BuildQueueAPI) SetQueuedBuildPosition(queuePosition string, data Build, fields string) (value Build, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildQueue/order/%s", queuePosition),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// getApprovalInfo
// approveQueuedBuild

// GetQueuedBuildTags Get tags of the queued matching build.
func (t BuildQueueAPI) GetQueuedBuildTags(buildLocator BuildQueueLocator, locator, fields string) (value Tags, err error) {
	req := Request{
		Path: fmt.Sprintf("/app/rest/buildQueue/%s/tags", locatorString(buildLocator)),
		Values: Values{
			"locator": locator,
			"fields":  fields,
		},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// AddTagsToBuildOfBuildQueue Add tags to the matching build.
func (t BuildQueueAPI) AddTagsToBuildOfBuildQueue(buildLocator BuildQueueLocator, data Tags) (value Tags, err error) {
	req := Request{
		Path: fmt.Sprintf("/app/rest/buildQueue/%s/tags", locatorString(buildLocator)),
		Data: data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// DeleteQueuedBuild Delete a queued matching build.
func (t BuildQueueAPI) DeleteQueuedBuild(buildQueueLocator BuildQueueLocator) (err error) {
	req := Request{Path: fmt.Sprintf("/app/rest/buildQueue/%s", locatorString(buildQueueLocator))}
	err = req.Delete(t.HTTPClient)
	return
}

// GetQueuedBuild Get a queued matching build.
func (t BuildQueueAPI) GetQueuedBuild(queuedBuildLocator BuildQueueLocator, fields string) (value *Build, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildQueue/%s", locatorString(queuedBuildLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, value)
	return
}

// CancelQueuedBuild Cancel a queued matching build.
func (t BuildQueueAPI) CancelQueuedBuild(queuedBuildLocator BuildQueueLocator, data BuildCancelRequest) (value Build, err error) {
	req := Request{
		Path: fmt.Sprintf("/app/rest/buildQueue/%s", locatorString(queuedBuildLocator)),
		Data: data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// GetCompatibleAgentsForBuild Get compatible agents for a queued matching build.
func (t BuildQueueAPI) GetCompatibleAgentsForBuild(queuedBuildLocator BuildQueueLocator, fields string) (value Agents, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildQueue/%s/compatibleAgents", locatorString(queuedBuildLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}
