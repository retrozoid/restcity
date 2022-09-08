package trac

import (
	"fmt"
	"net/http"
)

type BuildAPI struct {
	HTTPClient *http.Client
}

// GetAggregatedBuildStatus Get the build status of aggregated matching builds.
func (t BuildAPI) GetAggregatedBuildStatus(buildLocator BuildLocator) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/builds/aggregated/%s/status", EncodeLocator(buildLocator)),
		Consumer: Coders.String,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// AddLogMessageToBuild Adds a message to the build log. Service messages are accepted.
func (t BuildAPI) AddLogMessageToBuild(buildLocator BuildLocator, data string, fields string) (err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/builds/%s/log", EncodeLocator(buildLocator)),
		Values:   Values{"fields": fields},
		Data:     data,
		Consumer: Coders.String,
	}
	err = req.Post(t.HTTPClient, nil)
	return
}

// SetMultipleBuildComments Update comments in multiple matching builds.
func (t BuildAPI) SetMultipleBuildComments(buildLocator BuildLocator, data string, fields string) (value MultipleOperationResult, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/multiple/%s/comment", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// DeleteMultipleBuildComments Delete comments of multiple matching builds.
func (t BuildAPI) DeleteMultipleBuildComments(buildLocator BuildLocator, fields string) (value MultipleOperationResult, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/multiple/%s/comment", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Do(t.HTTPClient, http.MethodDelete, &value, nil)
	return
}

// GetBuildStatusIcon Get the status icon (in specified format) of the matching build.
func (t BuildAPI) GetBuildStatusIcon(buildLocator BuildLocator, suffix string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/builds/%s/statusIcon/%s", EncodeLocator(buildLocator), suffix),
		Consumer: Coders.String,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetZippedFileOfBuild Get specific file zipped.
func (t BuildAPI) GetZippedFileOfBuild(buildLocator BuildLocator, path string, basePath string, locator string, name string, resolveParameters bool, logBuildUsage bool) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/builds/%s/artifacts/archived/%s", EncodeLocator(buildLocator), path),
		Values:   Values{"basePath": basePath, "locator": locator},
		Consumer: Coders.String,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// ResetBuildFinishProperties Remove build parameters from the matching build.
func (t BuildAPI) ResetBuildFinishProperties(buildLocator BuildLocator) (err error) {
	req := Request{Path: fmt.Sprintf("/app/rest/builds/%s/caches/finishProperties", EncodeLocator(buildLocator))}
	err = req.Delete(t.HTTPClient)
	return
}

// GetBuildStatisticValues Get all statistical values of the matching build.
func (t BuildAPI) GetBuildStatisticValues(buildLocator BuildLocator, fields string) (value Properties, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/%s/statistics", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetBuildVcsLabels Get VCS labels of the matching build.
func (t BuildAPI) GetBuildVcsLabels(buildLocator BuildLocator, fields string) (value VcsLabels, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/%s/vcsLabels", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// AddBuildVcsLabel Add a VCS label to the matching build.
func (t BuildAPI) AddBuildVcsLabel(buildLocator BuildLocator, locator VcsRootInstanceLocator, fields string, data string) (value VcsLabels, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/%s/vcsLabels", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// PinMultipleBuilds Pin multiple matching builds.
func (t BuildAPI) PinMultipleBuilds(buildLocator BuildLocator, data PinInfo, fields string) (value MultipleOperationResult, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/multiple/%s/pinInfo", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetFilesListOfBuild List all files.
func (t BuildAPI) GetFilesListOfBuild(basePath string, locator string, fields string, buildLocator BuildLocator, resolveParameters bool, logBuildUsage bool) (value Files, err error) {
	req := Request{
		Path: fmt.Sprintf("/app/rest/builds/%s/artifacts", EncodeLocator(buildLocator)),
		Values: Values{
			"basePath":          basePath,
			"locator":           locator,
			"fields":            fields,
			"resolveParameters": resolveParameters,
			"logBuildUsage":     logBuildUsage,
		},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetCanceledInfo Check if the matching build is canceled.
func (t BuildAPI) GetCanceledInfo(buildLocator BuildLocator, fields string) (value Comment, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/%s/canceledInfo", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// SetBuildNumber Update the number of the matching build.
func (t BuildAPI) SetBuildNumber(buildLocator BuildLocator, data string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/builds/%s/number", EncodeLocator(buildLocator)),
		Data:     data,
		Consumer: Coders.String,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetBuildNumber Get the number of the matching build.
func (t BuildAPI) GetBuildNumber(buildLocator BuildLocator) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/builds/%s/number", EncodeLocator(buildLocator)),
		Consumer: Coders.String,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetBuildTags Get tags of the matching build.
func (t BuildAPI) GetBuildTags(buildLocator BuildLocator, locator TagLocator, fields string) (value Tags, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/%s/tags", EncodeLocator(buildLocator)),
		Values: Values{"locator": EncodeLocator(locator), "fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// AddTagsToBuild Add tags to the matching build.
func (t BuildAPI) AddTagsToBuild(buildLocator BuildLocator, data Tags, fields string) (value Tags, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/%s/tags", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// SetBuildTags Update tags of the matching build.
func (t BuildAPI) SetBuildTags(buildLocator BuildLocator, locator TagLocator, data Tags, fields string) (value Tags, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/%s/tags", EncodeLocator(buildLocator)),
		Values: Values{"locator": EncodeLocator(locator), "fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetFileMetadataOfBuild Get metadata of specific file.
func (t BuildAPI) GetFileMetadataOfBuild(buildLocator BuildLocator, path string, fields string, resolveParameters bool, logBuildUsage bool) (value File, err error) {
	req := Request{
		Path: fmt.Sprintf("/app/rest/builds/%s/artifacts/metadata/%s", EncodeLocator(buildLocator), path),
		Values: Values{
			"resolveParameters": resolveParameters,
			"logBuildUsage":     logBuildUsage,
			"fields":            fields,
		},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetBuildResultingProperties Update a build parameter of the matching build.
func (t BuildAPI) GetBuildResultingProperties(buildLocator BuildLocator, propertyName string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/builds/%s/resulting-properties/%s", EncodeLocator(buildLocator), propertyName),
		Consumer: Coders.String,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// AddTagsToMultipleBuilds Add tags to multiple matching builds.
func (t BuildAPI) AddTagsToMultipleBuilds(buildLocator BuildLocator, data Tags, fields string) (value MultipleOperationResult, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/multiple/%s/tags", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// RemoveMultipleBuildTags Remove tags from multiple matching builds.
func (t BuildAPI) RemoveMultipleBuildTags(buildLocator BuildLocator, data Tags, fields string) (value MultipleOperationResult, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/multiple/%s/tags", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Do(t.HTTPClient, http.MethodDelete, &value, nil)
	return
}

// SetBuildPinInfo Update the pin info of the matching build.
func (t BuildAPI) SetBuildPinInfo(buildLocator BuildLocator, data PinInfo, fields string) (value PinInfo, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/%s/pinInfo", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetBuildPinInfo Check if the matching build is pinned.
func (t BuildAPI) GetBuildPinInfo(buildLocator BuildLocator, fields string) (value PinInfo, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/%s/pinInfo", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetBuildSourceFile Get a source file of the matching build.
func (t BuildAPI) GetBuildSourceFile(buildLocator BuildLocator, fileName string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/builds/%s/sources/files/%s", EncodeLocator(buildLocator), fileName),
		Consumer: Coders.Stream,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetAggregatedBuildStatusIcon Get the status icon (in specified format) of aggregated matching builds.
func (t BuildAPI) GetAggregatedBuildStatusIcon(buildLocator BuildLocator, suffix string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/builds/aggregated/%s/statusIcon/%s", EncodeLocator(buildLocator), suffix),
		Consumer: Coders.String,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetBuildProblems Get build problems of the matching build.
func (t BuildAPI) GetBuildProblems(buildLocator BuildLocator, fields string) (value ProblemOccurrences, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/%s/problemOccurrences", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// AddProblemToBuild Add a build problem to the matching build.
func (t BuildAPI) AddProblemToBuild(buildLocator BuildLocator, data string, fields string) (value ProblemOccurrence, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/%s/problemOccurrences", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// GetBuildField Get a field of the matching build.
func (t BuildAPI) GetBuildField(buildLocator BuildLocator, field string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/builds/%s/%s", EncodeLocator(buildLocator), field),
		Consumer: Coders.String,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// SetFinishedTime Marks the running build as finished by passing agent the current time of the build to finish.
func (t BuildAPI) SetFinishedTime(buildLocator BuildLocator) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/builds/%s/finish", EncodeLocator(buildLocator)),
		Consumer: Coders.String,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetArtifactsDirectory Get the artifacts' directory of the matching build.
func (t BuildAPI) GetArtifactsDirectory(buildLocator BuildLocator) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/builds/%s/artifactsDirectory", EncodeLocator(buildLocator)),
		Consumer: Coders.String,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetBuildFinishDate Get the finish date of the matching build.
func (t BuildAPI) GetBuildFinishDate(buildLocator BuildLocator) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/builds/%s/finishDate", EncodeLocator(buildLocator)),
		Consumer: Coders.String,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// SetBuildFinishDate Marks the running build as finished by passing agent the current time of the build to finish.
func (t BuildAPI) SetBuildFinishDate(buildLocator BuildLocator, data string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/builds/%s/finishDate", EncodeLocator(buildLocator)),
		Consumer: Coders.String,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetAllBuilds Get all builds.
func (t BuildAPI) GetAllBuilds(buildQueueLocator BuildLocator, fields string) (value Builds, err error) {
	req := Request{
		Path:   "/app/rest/builds",
		Values: Values{"locator": EncodeLocator(buildQueueLocator), "fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetBuildStatisticValue Get a statistical value of the matching build.
func (t BuildAPI) GetBuildStatisticValue(buildLocator BuildLocator, name string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/builds/%s/statistics/%s", EncodeLocator(buildLocator), name),
		Consumer: Coders.String,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetBuildTestOccurrences Get test occurrences of the matching build.
func (t BuildAPI) GetBuildTestOccurrences(buildLocator BuildLocator, fields string) (value TestOccurrences, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/%s/testOccurrences", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetBuildRelatedIssues Get related issues of the matching build.
func (t BuildAPI) GetBuildRelatedIssues(buildLocator BuildLocator, fields string) (value IssuesUsages, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/%s/relatedIssues", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// SetBuildStatusText Update the build status of the matching build.
func (t BuildAPI) SetBuildStatusText(buildLocator BuildLocator, data string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/builds/%s/statusText", EncodeLocator(buildLocator)),
		Data:     data,
		Consumer: Coders.String,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetBuildStatusText Get the build status text of the matching build.
func (t BuildAPI) GetBuildStatusText(buildLocator BuildLocator) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/builds/%s/statusText", EncodeLocator(buildLocator)),
		Consumer: Coders.String,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetBuildResolved Get the resolvement status of the matching build.
func (t BuildAPI) GetBuildResolved(buildLocator BuildLocator, val string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/builds/%s/resolved/%s", EncodeLocator(buildLocator), val),
		Consumer: Coders.String,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetFilesListForSubpathOfBuild List files under this path.
func (t BuildAPI) GetFilesListForSubpathOfBuild(path string, basePath string, locator string, fields string, buildLocator BuildLocator, resolveParameters bool, logBuildUsage bool) (value Files, err error) {
	req := Request{
		Path: fmt.Sprintf("/app/rest/builds/%s/artifacts/%s", EncodeLocator(buildLocator), path),
		Values: Values{
			"basePath":          basePath,
			"locator":           locator,
			"fields":            fields,
			"resolveParameters": resolveParameters,
			"logBuildUsage":     logBuildUsage,
		},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetBuild Get build matching the locator.
func (t BuildAPI) GetBuild(buildLocator BuildLocator, fields string) (value Build, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/%s", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// CancelBuild cancelBuild
func (t BuildAPI) CancelBuild(buildLocator BuildLocator, data BuildCancelRequest, fields string) (value Build, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/%s", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// DeleteBuild Delete build matching the locator.
func (t BuildAPI) DeleteBuild(buildLocator BuildLocator) (err error) {
	req := Request{Path: fmt.Sprintf("/app/rest/builds/%s", EncodeLocator(buildLocator))}
	err = req.Delete(t.HTTPClient)
	return
}

// GetArtifactDependencyChanges Get artifact dependency changes of the matching build.
func (t BuildAPI) GetArtifactDependencyChanges(buildLocator BuildLocator, fields string) (value BuildChanges, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/%s/artifactDependencyChanges", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetBuildActualParameters Get actual build parameters of the matching build.
func (t BuildAPI) GetBuildActualParameters(buildLocator BuildLocator, fields string) (value Properties, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/%s/resulting-properties", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// DownloadFileOfBuild Download specific file.
func (t BuildAPI) DownloadFileOfBuild(buildLocator BuildLocator, path string, resolveParameters bool, logBuildUsage bool) (value string, err error) {
	req := Request{
		Path: fmt.Sprintf("/app/rest/builds/%s/artifacts/files/%s", EncodeLocator(buildLocator), path),
		Values: Values{
			"resolveParameters": resolveParameters,
			"logBuildUsage":     logBuildUsage,
		},
		Consumer: Coders.String,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// SetBuildComment Update the comment on the matching build.
func (t BuildAPI) SetBuildComment(buildLocator BuildLocator, data string) (err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/builds/%s/comment", EncodeLocator(buildLocator)),
		Data:     data,
		Consumer: Coders.String,
	}
	err = req.Put(t.HTTPClient, nil)
	return
}

// DeleteBuildComment Remove the build comment matching the locator.
func (t BuildAPI) DeleteBuildComment(buildLocator BuildLocator) (err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/builds/%s/comment", EncodeLocator(buildLocator)),
		Consumer: Coders.String,
	}
	err = req.Delete(t.HTTPClient)
	return
}

// MarkBuildAsRunning Starts the queued build as an agent-less build and returns the corresponding running build.
func (t BuildAPI) MarkBuildAsRunning(buildLocator BuildLocator, data string, fields string) (value Build, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/builds/%s/runningData", EncodeLocator(buildLocator)),
		Values:   Values{"fields": fields},
		Data:     data,
		Consumer: Coders.String,
	}
	err = req.Do(t.HTTPClient, http.MethodPut, &value, Coders.JSON)
	return
}

// GetMultipleBuilds Get multiple builds matching the locator.
func (t BuildAPI) GetMultipleBuilds(buildLocator BuildLocator, fields string) (value Builds, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/multiple/%s", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// CancelMultiple cancelMultipleBuilds
func (t BuildAPI) CancelMultiple(buildLocator BuildLocator, data BuildCancelRequest, fields string) (value MultipleOperationResult, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/multiple/%s", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// DeleteMultipleBuilds Delete multiple builds matching the locator.
func (t BuildAPI) DeleteMultipleBuilds(buildLocator BuildLocator, fields string) (value MultipleOperationResult, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/builds/multiple/%s", EncodeLocator(buildLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Do(t.HTTPClient, http.MethodDelete, &value, nil)
	return
}
