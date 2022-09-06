package trac

import (
	"fmt"
	"net/http"
)

type BuildTypeAPI struct {
	HTTPClient *http.Client
}

// GetBuildType Get build configuration matching the locator.
func (t BuildTypeAPI) GetBuildType(btLocator BuildTypeLocator, fields string) (value BuildTypes, err error) {
	req := Request{
		Path:   "/app/rest/buildTypes/" + locatorString(btLocator),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// DeleteBuildType Delete build configuration matching the locator.
func (t BuildTypeAPI) DeleteBuildType(btLocator BuildTypeLocator) (err error) {
	err = Request{Path: fmt.Sprintf("/app/rest/buildTypes/", locatorString(btLocator))}.Delete(t.HTTPClient)
	return
}

// GetBuildTypeBuildTags Get tags of builds of the matching build configuration.
func (t BuildTypeAPI) GetBuildTypeBuildTags(btLocator BuildTypeLocator, field string) (value Tags, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/buildTags", locatorString(btLocator)),
		Values: Values{"field": field},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetBuildParameterValueOfBuildType Get value of build parameter.
func (t BuildTypeAPI) GetBuildParameterValueOfBuildType(btLocator BuildTypeLocator, name string) (value string, err error) {
	req := Request{Path: fmt.Sprintf("/app/rest/buildTypes/%s/parameters/%s/value", locatorString(btLocator), name)}
	err = req.Get(t.HTTPClient, &value)
	return
}

// UpdateBuildParameterValueOfBuildType Update value of build parameter.
func (t BuildTypeAPI) UpdateBuildParameterValueOfBuildType(btLocator BuildTypeLocator, name string, data string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/buildTypes/%s/parameters/%s/value", locatorString(btLocator), name),
		Data:     data,
		Consumes: TextPlain,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetAllVcsRootsOfBuildType Get all VCS roots of the matching build configuration.
func (t BuildTypeAPI) GetAllVcsRootsOfBuildType(btLocator BuildTypeLocator, fields string) (value VcsRootEntries, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/vcs-root-entries", locatorString(btLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// AddVcsRootToBuildType Add a VCS root to the matching build.
func (t BuildTypeAPI) AddVcsRootToBuildType(btLocator BuildTypeLocator, data VcsRootEntry, fields string) (value VcsRootEntry, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/vcs-root-entries", locatorString(btLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// ReplaceAllVcsRoots Update all VCS roots of the matching build configuration.
func (t BuildTypeAPI) ReplaceAllVcsRoots(btLocator BuildTypeLocator, data VcsRootEntries, fields string) (value VcsRootEntries, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/vcs-root-entries", locatorString(btLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetArtifactDependency Get an artifact dependency of the matching build configuration.
func (t BuildTypeAPI) GetArtifactDependency(btLocator BuildTypeLocator, artifactDepLocator string, fields string) (value ArtifactDependency, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/artifact-dependencies/%s", locatorString(btLocator), artifactDepLocator),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// ReplaceArtifactDependency Update an artifact dependency of the matching build configuration.
func (t BuildTypeAPI) ReplaceArtifactDependency(btLocator BuildTypeLocator, artifactDepLocator string, fields string, data ArtifactDependency) (value ArtifactDependency, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/artifact-dependencies/%s", locatorString(btLocator), artifactDepLocator),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// DeleteArtifactDependency Remove an artifact dependency from the matching build configuration.
func (t BuildTypeAPI) DeleteArtifactDependency(btLocator BuildTypeLocator, artifactDepLocator string) (err error) {
	req := Request{Path: fmt.Sprintf("/app/rest/buildTypes/%s/artifact-dependencies/%s", locatorString(btLocator), artifactDepLocator)}
	err = req.Delete(t.HTTPClient)
	return
}

// GetFilesListForSubpathOfBuildType List files under this path.
func (t BuildTypeAPI) GetFilesListForSubpathOfBuildType(
	path string,
	basePath string,
	locator string,
	fields string,
	btLocator BuildTypeLocator,
	resolveParameters bool,
) (value Files, err error) {
	req := Request{
		Path: fmt.Sprintf("/app/rest/buildTypes/%s/vcs/files/latest/{path}", locatorString(btLocator), path),
		Values: Values{
			"basePath":          basePath,
			"locator":           locator,
			"fields":            fields,
			"resolveParameters": resolveParameters,
		},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetBuildStepSetting Get the setting of a build step of the matching build configuration.
func (t BuildTypeAPI) GetBuildStepSetting(btLocator BuildTypeLocator, stepId string, fieldName string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/buildTypes/%s/steps/%s/%s", locatorString(btLocator), stepId, fieldName),
		Consumes: TextPlain,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// SetBuildStepParameter Update a parameter of a build step of the matching build configuration.
func (t BuildTypeAPI) SetBuildStepParameter(btLocator BuildTypeLocator, stepId string, fieldName string, data string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/buildTypes/%s/steps/%s/%s", locatorString(btLocator), stepId, fieldName),
		Consumes: TextPlain,
		Data:     data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetFileMetadataOfBuildType Get metadata of specific file.
func (t BuildTypeAPI) GetFileMetadataOfBuildType(path, fields string, btLocator BuildTypeLocator, resolveParameters bool) (value File, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/vcs/files/latest/metadata%s", locatorString(btLocator), path),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// ReplaceBuildStep Replace a build step of the matching build configuration.
func (t BuildTypeAPI) ReplaceBuildStep(btLocator BuildTypeLocator, stepId string, fields string, data Step) (value Step, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/steps/%s", locatorString(btLocator), stepId),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// DeleteBuildStep Delete a build step of the matching build configuration.
func (t BuildTypeAPI) DeleteBuildStep(btLocator BuildTypeLocator, stepId string) (err error) {
	req := Request{Path: fmt.Sprintf("/app/rest/buildTypes/%s/steps/%s", locatorString(btLocator), stepId)}
	err = req.Delete(t.HTTPClient)
	return
}

// GetBuildStep Get a build step of the matching build configuration.
func (t BuildTypeAPI) GetBuildStep(btLocator BuildTypeLocator, stepId string, fields string) (value Step, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/steps/%s", locatorString(btLocator), stepId),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetAllBuildStepParameters Get all parameters of a build step of the matching build configuration.
func (t BuildTypeAPI) GetAllBuildStepParameters(btLocator BuildTypeLocator, stepId string, fields string) (value Properties, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/steps/%s/parameters", locatorString(btLocator), stepId),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// DeleteBuildStepParameters Update a parameter of a build step of the matching build configuration.
func (t BuildTypeAPI) DeleteBuildStepParameters(btLocator BuildTypeLocator, stepId string, data Properties, fields string) (value Properties, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/steps/%s/parameters", locatorString(btLocator), stepId),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetTrigger Get a trigger of the matching build configuration.
func (t BuildTypeAPI) GetTrigger(btLocator BuildTypeLocator, triggerLocator string, fields string) (value Trigger, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/triggers/%s", locatorString(btLocator), triggerLocator),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// ReplaceTrigger Update a trigger of the matching build configuration.
func (t BuildTypeAPI) ReplaceTrigger(btLocator BuildTypeLocator, triggerLocator string, fields string, data Trigger) (value Trigger, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/triggers/%s", locatorString(btLocator), triggerLocator),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// DeleteTrigger Delete a trigger of the matching build configuration.
func (t BuildTypeAPI) DeleteTrigger(btLocator BuildTypeLocator, triggerLocator string) (err error) {
	req := Request{Path: fmt.Sprintf("/app/rest/buildTypes/%s/triggers/%s", locatorString(btLocator), triggerLocator)}
	err = req.Delete(t.HTTPClient)
	return
}

// GetAgentRequirementParameter Get a setting of an agent requirement of the matching build configuration.
func (t BuildTypeAPI) GetAgentRequirementParameter(btLocator BuildTypeLocator, agentRequirementLocator, fieldName string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/buildTypes/%s/agent-requirements/%s/%s", locatorString(btLocator), agentRequirementLocator, fieldName),
		Consumes: TextPlain,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// SetAgentRequirementParameter Update a parameter of an agent requirement of the matching build configuration.
func (t BuildTypeAPI) SetAgentRequirementParameter(btLocator BuildTypeLocator, agentRequirementLocator string, fieldName string, data string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/buildTypes/%s/agent-requirements/%s/%s", locatorString(btLocator), agentRequirementLocator, fieldName),
		Consumes: TextPlain,
		Data:     data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// DeleteAgentRequirement Remove an agent requirement of the matching build configuration.
func (t BuildTypeAPI) DeleteAgentRequirement(btLocator BuildTypeLocator, agentRequirementLocator string) (err error) {
	req := Request{Path: fmt.Sprintf("/app/rest/buildTypes/%s/agent-requirements/%s", locatorString(btLocator), agentRequirementLocator)}
	err = req.Delete(t.HTTPClient)
	return
}

// GetAgentRequirement Get an agent requirement of the matching build configuration.
func (t BuildTypeAPI) GetAgentRequirement(btLocator BuildTypeLocator, agentRequirementLocator string, fields string) (value AgentRequirement, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/agent-requirements/%s", locatorString(btLocator), agentRequirementLocator),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// ReplaceAgentRequirement Update an agent requirement of the matching build configuration.
func (t BuildTypeAPI) ReplaceAgentRequirement(btLocator BuildTypeLocator, agentRequirementLocator string, fields string, data AgentRequirement) (value AgentRequirement, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/agent-requirements/%s", locatorString(btLocator), agentRequirementLocator),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetAllInvestigationsOfBuildType Get all investigations of the matching build configuration.
func (t BuildTypeAPI) GetAllInvestigationsOfBuildType(btLocator BuildTypeLocator, fields string) (value Investigations, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/investigations", locatorString(btLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetBuildParameterOfBuildType Get build parameter.
func (t BuildTypeAPI) GetBuildParameterOfBuildType(btLocator BuildTypeLocator, name string, fields string) (value Property, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/parameters/%s", locatorString(btLocator), name),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// UpdateBuildParameterOfBuildType Update build parameter.
func (t BuildTypeAPI) UpdateBuildParameterOfBuildType(btLocator BuildTypeLocator, name string, data Property, fields string) (value Property, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/parameters/%s", locatorString(btLocator), name),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// DeleteBuildParameterOfBuildType Delete build parameter.
func (t BuildTypeAPI) DeleteBuildParameterOfBuildType(btLocator BuildTypeLocator, name string) (err error) {
	req := Request{Path: fmt.Sprintf("/app/rest/buildTypes/%s/parameters/%s", locatorString(btLocator), name)}
	err = req.Delete(t.HTTPClient)
	return
}

// GetAllBuildTypes Get all build configurations.
func (t BuildTypeAPI) GetAllBuildTypes(locator BuildTypeLocator, fields string) (value BuildTypes, err error) {
	req := Request{
		Path: "/app/rest/buildTypes",
		Values: Values{
			"locator": locatorString(locator),
			"fields":  fields,
		},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// CreateBuildType Create a new build configuration.
func (t BuildTypeAPI) CreateBuildType(data BuildType, fields string) (value BuildType, err error) {
	req := Request{
		Path:   "/app/rest/buildTypes",
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// AddArtifactDependencyToBuildType Add an artifact dependency to the matching build configuration.
func (t BuildTypeAPI) AddArtifactDependencyToBuildType(btLocator BuildTypeLocator, fields string, data ArtifactDependency) (value ArtifactDependency, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/artifact-dependencies", locatorString(btLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// ReplaceAllArtifactDependencies Update all artifact dependencies of the matching build configuration.
func (t BuildTypeAPI) ReplaceAllArtifactDependencies(btLocator BuildTypeLocator, fields string, data ArtifactDependency) (value ArtifactDependency, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/artifact-dependencies", locatorString(btLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetAllArtifactDependencies Get all artifact dependencies of the matching build configuration.
func (t BuildTypeAPI) GetAllArtifactDependencies(btLocator BuildTypeLocator, fields string) (value ArtifactDependencies, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/artifact-dependencies", locatorString(btLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// UpdateBuildParameterSpecificationOfBuildType Update build parameter specification.
func (t BuildTypeAPI) UpdateBuildParameterSpecificationOfBuildType(btLocator BuildTypeLocator, name string, data string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/buildTypes/%s/parameters/%s/type/rawValue", locatorString(btLocator), name),
		Consumes: TextPlain,
		Data:     data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetBuildParameterSpecificationOfBuildType Get build parameter specification.
func (t BuildTypeAPI) GetBuildParameterSpecificationOfBuildType(name string, btLocator BuildTypeLocator) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/buildTypes/%s/parameters/%s/type/rawValue", locatorString(btLocator), name),
		Consumes: TextPlain,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetVcsRootCheckoutRules Get checkout rules of a VCS root of the matching build configuration.
func (t BuildTypeAPI) GetVcsRootCheckoutRules(btLocator BuildTypeLocator, vcsRootLocator string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/buildTypes/%s/vcs-root-entries/%s/checkout-rules", locatorString(btLocator), vcsRootLocator),
		Consumes: TextPlain,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// UpdateBuildTypeVcsRootCheckoutRules Update checkout rules of a VCS root of the matching build configuration.
func (t BuildTypeAPI) UpdateBuildTypeVcsRootCheckoutRules(btLocator BuildTypeLocator, vcsRootLocator string, data string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/buildTypes/%s/vcs-root-entries/%s/checkout-rules", locatorString(btLocator), vcsRootLocator),
		Consumes: TextPlain,
		Data:     data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetAllBranchesOfBuildType Get all branches of the matching build configuration.
func (t BuildTypeAPI) GetAllBranchesOfBuildType(btLocator BuildTypeLocator, locator string, fields string) (value Branches, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/branches", locatorString(btLocator)),
		Values: Values{"fields": fields, "locator": locator},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetTriggerParameter Get a parameter of a trigger of the matching build configuration.
func (t BuildTypeAPI) GetTriggerParameter(btLocator BuildTypeLocator, triggerLocator string, fieldName string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/buildTypes/%s/triggers/%s/%s", locatorString(btLocator), triggerLocator, fieldName),
		Consumes: TextPlain,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// SetTriggerParameter Update a parameter of a trigger of the matching build configuration.
func (t BuildTypeAPI) SetTriggerParameter(btLocator BuildTypeLocator, triggerLocator string, fieldName string, data string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/buildTypes/%s/triggers/%s/%s", locatorString(btLocator), triggerLocator, fieldName),
		Consumes: TextPlain,
		Data:     data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetBuildTemplate Get a template of the matching build configuration.
func (t BuildTypeAPI) GetBuildTemplate(btLocator BuildTypeLocator, templateLocator string, fields string) (value BuildType, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/templates/%s", locatorString(btLocator), templateLocator),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// RemoveTemplate Detach a template from the matching build configuration.
func (t BuildTypeAPI) RemoveTemplate(btLocator BuildTypeLocator, templateLocator string, inlineSettings bool) (err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/templates/%s", locatorString(btLocator), templateLocator),
		Values: Values{"inlineSettings": inlineSettings},
	}
	err = req.Delete(t.HTTPClient)
	return
}

// GetVcsRootInstancesOfBuildType Get all VCS root instances of the matching build configuration.
func (t BuildTypeAPI) GetVcsRootInstancesOfBuildType(btLocator BuildTypeLocator, fields string) (value VcsRootInstances, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/vcsRootInstances", locatorString(btLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// DownloadFileOfBuildType Download specific file.
func (t BuildTypeAPI) DownloadFileOfBuildType(path string, btLocator BuildTypeLocator, resolveParameters bool) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/buildTypes/%s/vcs/files/latest/files%s", locatorString(btLocator), path),
		Values:   Values{"resolveParameters": resolveParameters},
		Consumes: TextPlain,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// ReplaceAllTriggers Update all triggers of the matching build configuration.
func (t BuildTypeAPI) ReplaceAllTriggers(btLocator BuildTypeLocator, fields string, data Triggers) (value Triggers, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/triggers", locatorString(btLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetAllTriggers Get all triggers of the matching build configuration.
func (t BuildTypeAPI) GetAllTriggers(btLocator BuildTypeLocator, fields string) (value Triggers, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/triggers", locatorString(btLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// AddTriggerToBuildType Add a trigger to the matching build configuration.
func (t BuildTypeAPI) AddTriggerToBuildType(btLocator BuildTypeLocator, fields string, data Trigger) (value Trigger, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/triggers", locatorString(btLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// GetBuildFeature Get a build feature of the matching build configuration.
func (t BuildTypeAPI) GetBuildFeature(btLocator BuildTypeLocator, featureId string, fields string) (value Feature, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/features/%s", locatorString(btLocator), featureId),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// ReplaceBuildFeature Update a build feature of the matching build configuration.
func (t BuildTypeAPI) ReplaceBuildFeature(btLocator BuildTypeLocator, featureId string, fields string, data Feature) (value Feature, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/features/%s", locatorString(btLocator), featureId),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// DeleteFeatureOfBuildType Remove a build feature of the matching build configuration.
func (t BuildTypeAPI) DeleteFeatureOfBuildType(btLocator BuildTypeLocator, featureId string) (err error) {
	req := Request{
		Path: fmt.Sprintf("/app/rest/buildTypes/%s/features/%s", locatorString(btLocator), featureId),
	}
	err = req.Delete(t.HTTPClient)
	return
}

// GetBuildParameterTypeOfBuildType Get type of build parameter.
func (t BuildTypeAPI) GetBuildParameterTypeOfBuildType(btLocator BuildTypeLocator, name string) (value ModelType, err error) {
	req := Request{
		Path: fmt.Sprintf("/app/rest/buildTypes/%s/parameters/%s/type", locatorString(btLocator), name),
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// UpdateBuildParameterTypeOfBuildType Update type of build parameter.
func (t BuildTypeAPI) UpdateBuildParameterTypeOfBuildType(btLocator BuildTypeLocator, name string, data ModelType) (value ModelType, err error) {
	req := Request{
		Path: fmt.Sprintf("/app/rest/buildTypes/%s/parameters/%s/type", locatorString(btLocator), name),
		Data: data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetBuildTypeBuilds Get builds of the matching build configuration.
func (t BuildTypeAPI) GetBuildTypeBuilds(btLocator BuildTypeLocator, fields string) (value Builds, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/builds", locatorString(btLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetBuildTypeSettingsFile Get the settings file of the matching build configuration.
func (t BuildTypeAPI) GetBuildTypeSettingsFile(btLocator BuildTypeLocator) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/buildTypes/%s/settingsFile", locatorString(btLocator)),
		Consumes: TextPlain,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetAllSnapshotDependencies Get all snapshot dependencies of the matching build configuration.
func (t BuildTypeAPI) GetAllSnapshotDependencies(btLocator BuildTypeLocator, fields string) (value SnapshotDependencies, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/snapshot-dependencies", locatorString(btLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// AddSnapshotDependencyToBuildType Add a snapshot dependency to the matching build configuration.
func (t BuildTypeAPI) AddSnapshotDependencyToBuildType(btLocator BuildTypeLocator, fields string, data SnapshotDependency) (value SnapshotDependency, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/snapshot-dependencies", locatorString(btLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// ReplaceAllSnapshotDependencies Update all snapshot dependencies of the matching build configuration.
func (t BuildTypeAPI) ReplaceAllSnapshotDependencies(btLocator BuildTypeLocator, fields string, data SnapshotDependencies) (value SnapshotDependencies, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/snapshot-dependencies", locatorString(btLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetAliases Get external IDs of the matching build configuration.
func (t BuildTypeAPI) GetAliases(btLocator BuildTypeLocator, field string) (value Items, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/aliases", locatorString(btLocator)),
		Values: Values{"field": field},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetZippedFileOfBuildType Get specific file zipped.
func (t BuildTypeAPI) GetZippedFileOfBuildType(path string, basePath string, locator string, name string, btLocator BuildTypeLocator, resolveParameters bool) (value string, err error) {
	req := Request{
		Path: fmt.Sprintf("/app/rest/buildTypes/%s/vcs/files/latest/archived%s", locatorString(btLocator), path),
		Values: Values{
			"basePath":          basePath,
			"locator":           locator,
			"resolveParameters": resolveParameters,
		},
		Consumes: TextPlain,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetBuildFeatureSetting Get the setting of a build feature of the matching build configuration.
func (t BuildTypeAPI) GetBuildFeatureSetting(btLocator BuildTypeLocator, featureId string, name string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/buildTypes/%s/features/%s/%s", locatorString(btLocator), featureId, name),
		Consumes: TextPlain,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// SetBuildFeatureParameter Update a parameter of a build feature of the matching build configuration.
func (t BuildTypeAPI) SetBuildFeatureParameter(btLocator BuildTypeLocator, featureId string, name string, data string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/buildTypes/%s/features/%s/%s", locatorString(btLocator), featureId, name),
		Consumes: TextPlain,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// // GetVcsRoot Get a VCS root of the matching build configuration.
// func (t BuildTypeAPI) GetVcsRoot(btLocator BuildTypeLocator, vcsRootLocator string, fields string) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/vcs-root-entries/{vcsRootLocator}",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Get(t.HTTPClient, &value)
// 	return
// }

// // UpdateBuildTypeVcsRoot Update a VCS root of the matching build configuration.
// func (t BuildTypeAPI) UpdateBuildTypeVcsRoot(btLocator BuildTypeLocator, vcsRootLocator string, body XXX, fields string) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/vcs-root-entries/{vcsRootLocator}",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Put(t.HTTPClient, &value)
// 	return
// }

// // DeleteVcsRootOfBuildType Remove a VCS root of the matching build configuration.
// func (t BuildTypeAPI) DeleteVcsRootOfBuildType(btLocator BuildTypeLocator, vcsRootLocator string) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/vcs-root-entries/{vcsRootLocator}",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Delete(t.HTTPClient, &value)
// 	return
// }

// // SetArtifactDependencyParameter Update a parameter of an artifact dependency of the matching build configuration.
// func (t BuildTypeAPI) SetArtifactDependencyParameter(btLocator BuildTypeLocator, artifactDepLocator string, fieldName string, body XXX) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/artifact-dependencies/{artifactDepLocator}/{fieldName}",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Put(t.HTTPClient, &value)
// 	return
// }

// // GetArtifactDependencyParameter Get a parameter of an artifact dependency of the matching build configuration.
// func (t BuildTypeAPI) GetArtifactDependencyParameter(btLocator BuildTypeLocator, artifactDepLocator string, fieldName string) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/artifact-dependencies/{artifactDepLocator}/{fieldName}",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Get(t.HTTPClient, &value)
// 	return
// }

// // GetAllBuildTemplates Get all build templates of the matching build configuration.
// func (t BuildTypeAPI) GetAllBuildTemplates(btLocator BuildTypeLocator, fields string) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/templates",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Get(t.HTTPClient, &value)
// 	return
// }

// // AddBuildTemplate Add a build template to the matching build configuration.
// func (t BuildTypeAPI) AddBuildTemplate(btLocator BuildTypeLocator, body XXX, optimizeSettings bool, fields string) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/templates",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Post(t.HTTPClient, &value)
// 	return
// }

// // SetBuildTypeTemplates Update all templates of the matching build configuration.
// func (t BuildTypeAPI) SetBuildTypeTemplates(btLocator BuildTypeLocator, body XXX, optimizeSettings bool, fields string) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/templates",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Put(t.HTTPClient, &value)
// 	return
// }

// // RemoveAllTemplates Detach all templates from the matching build configuration.
// func (t BuildTypeAPI) RemoveAllTemplates(btLocator BuildTypeLocator, inlineSettings bool) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/templates",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Delete(t.HTTPClient, &value)
// 	return
// }

// // AddBuildFeatureToBuildType Add build feature to the matching build configuration.
// func (t BuildTypeAPI) AddBuildFeatureToBuildType(btLocator BuildTypeLocator, fields string, body XXX) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/features",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Post(t.HTTPClient, &value)
// 	return
// }

// // ReplaceAllBuildFeatures Update all build features of the matching build configuration.
// func (t BuildTypeAPI) ReplaceAllBuildFeatures(btLocator BuildTypeLocator, fields string, body XXX) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/features",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Put(t.HTTPClient, &value)
// 	return
// }

// // GetAllBuildFeatures Get all build features of the matching build configuration.
// func (t BuildTypeAPI) GetAllBuildFeatures(btLocator BuildTypeLocator, fields string) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/features",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Get(t.HTTPClient, &value)
// 	return
// }

// // GetBuildFeatureParameter Get a parameter of a build feature of the matching build configuration.
// func (t BuildTypeAPI) GetBuildFeatureParameter(btLocator BuildTypeLocator, featureId string, parameterName string) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/features/{featureId}/parameters/{parameterName}",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Get(t.HTTPClient, &value)
// 	return
// }

// // AddParameterToBuildFeature Update build feature parameter for the matching build configuration.
// func (t BuildTypeAPI) AddParameterToBuildFeature(btLocator BuildTypeLocator, featureId string, parameterName string, body XXX) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/features/{featureId}/parameters/{parameterName}",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Put(t.HTTPClient, &value)
// 	return
// }

// // GetBuildParametersOfBuildType Get build parameters.
// func (t BuildTypeAPI) GetBuildParametersOfBuildType(locator string, fields string, btLocator BuildTypeLocator) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/parameters",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Get(t.HTTPClient, &value)
// 	return
// }

// // CreateBuildParameterOfBuildType Create a build parameter.
// func (t BuildTypeAPI) CreateBuildParameterOfBuildType(body XXX, fields string, btLocator BuildTypeLocator) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/parameters",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Post(t.HTTPClient, &value)
// 	return
// }

// // UpdateBuildParametersOfBuildType Update build parameters.
// func (t BuildTypeAPI) UpdateBuildParametersOfBuildType(body XXX, fields string, btLocator BuildTypeLocator) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/parameters",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Put(t.HTTPClient, &value)
// 	return
// }

// // DeleteBuildParametersOfBuildType Delete all build parameters.
// func (t BuildTypeAPI) DeleteBuildParametersOfBuildType(btLocator BuildTypeLocator) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/parameters",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Delete(t.HTTPClient, &value)
// 	return
// }

// // GetSnapshotDependency Get a snapshot dependency of the matching build configuration.
// func (t BuildTypeAPI) GetSnapshotDependency(btLocator BuildTypeLocator, snapshotDepLocator string, fields string) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/snapshot-dependencies/{snapshotDepLocator}",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Get(t.HTTPClient, &value)
// 	return
// }

// // ReplaceSnapshotDependency Update a snapshot dependency of the matching build configuration.
// func (t BuildTypeAPI) ReplaceSnapshotDependency(btLocator BuildTypeLocator, snapshotDepLocator string, fields string, body XXX) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/snapshot-dependencies/{snapshotDepLocator}",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Put(t.HTTPClient, &value)
// 	return
// }

// // DeleteSnapshotDependency Delete a snapshot dependency of the matching build configuration.
// func (t BuildTypeAPI) DeleteSnapshotDependency(btLocator BuildTypeLocator, snapshotDepLocator string) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/snapshot-dependencies/{snapshotDepLocator}",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Delete(t.HTTPClient, &value)
// 	return
// }

// // GetAllAgentRequirements Get all agent requirements of the matching build configuration.
// func (t BuildTypeAPI) GetAllAgentRequirements(btLocator BuildTypeLocator, fields string) (value AgentRequirements, err error) {
// 	req := Request{
// 		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/agent-requirements", locatorString(btLocator)),
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Get(t.HTTPClient, &value)
// 	return
// }

// // AddAgentRequirementToBuildType Add an agent requirement to the matching build configuration.
// func (t BuildTypeAPI) AddAgentRequirementToBuildType(btLocator BuildTypeLocator, fields string, data AgentRequirement) (value AgentRequirement, err error) {
// 	req := Request{
// 		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/agent-requirements", locatorString(btLocator)),
// 		Values: Values{"fields": fields},
// 		Data:   data,
// 	}
// 	err = req.Post(t.HTTPClient, &value)
// 	return
// }

// // ReplaceAllAgentRequirements Update all agent requirements of the matching build configuration.
// func (t BuildTypeAPI) ReplaceAllAgentRequirements(btLocator BuildTypeLocator, fields string, data AgentRequirements) (value AgentRequirements, err error) {
// 	req := Request{
// 		Path:   fmt.Sprintf("/app/rest/buildTypes/%s/agent-requirements", locatorString(btLocator)),
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Put(t.HTTPClient, &value)
// 	return
// }

// // AddBuildStepToBuildType Add a build step to the matching build configuration.
// func (t BuildTypeAPI) AddBuildStepToBuildType(btLocator BuildTypeLocator, fields string, body XXX) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/steps",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Post(t.HTTPClient, &value)
// 	return
// }

// // ReplaceAllBuildSteps Update all build steps of the matching build configuration.
// func (t BuildTypeAPI) ReplaceAllBuildSteps(btLocator BuildTypeLocator, fields string, body XXX) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/steps",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Put(t.HTTPClient, &value)
// 	return
// }

// // GetAllBuildSteps Get all build steps of the matching build configuration.
// func (t BuildTypeAPI) GetAllBuildSteps(btLocator BuildTypeLocator, fields string) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/steps",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Get(t.HTTPClient, &value)
// 	return
// }

// // GetBuildStepParameter Get a parameter of a build step of the matching build configuration.
// func (t BuildTypeAPI) GetBuildStepParameter(btLocator BuildTypeLocator, stepId string, parameterName string) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/steps/{stepId}/parameters/{parameterName}",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Get(t.HTTPClient, &value)
// 	return
// }

// // AddParameterToBuildStep Add a parameter to a build step of the matching build configuration.
// func (t BuildTypeAPI) AddParameterToBuildStep(btLocator BuildTypeLocator, stepId string, parameterName string, body XXX) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/steps/{stepId}/parameters/{parameterName}",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Put(t.HTTPClient, &value)
// 	return
// }

// // GetBuildTypeField Get a field of the matching build configuration.
// func (t BuildTypeAPI) GetBuildTypeField(btLocator BuildTypeLocator, field string) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/{field}",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Get(t.HTTPClient, &value)
// 	return
// }

// // SetBuildTypeField Update a field of the matching build configuration.
// func (t BuildTypeAPI) SetBuildTypeField(btLocator BuildTypeLocator, field string, body XXX) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/{field}",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Put(t.HTTPClient, &value)
// 	return
// }

// // GetAllBuildFeatureParameters Get all parameters of a build feature of the matching build configuration.
// func (t BuildTypeAPI) GetAllBuildFeatureParameters(btLocator BuildTypeLocator, featureId string, fields string) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/features/{featureId}/parameters",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Get(t.HTTPClient, &value)
// 	return
// }

// // ReplaceBuildFeatureParameters Update a parameter of a build feature of the matching build configuration.
// func (t BuildTypeAPI) ReplaceBuildFeatureParameters(btLocator BuildTypeLocator, featureId string, body XXX, fields string) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/features/{featureId}/parameters",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Put(t.HTTPClient, &value)
// 	return
// }

// // GetFilesListOfBuildType List all files.
// func (t BuildTypeAPI) GetFilesListOfBuildType(basePath string, locator string, fields string, btLocator BuildTypeLocator, resolveParameters bool) (value XXX, err error) {
// 	req := Request{
// 		Path:   "/app/rest/buildTypes/{btLocator}/vcs/files/latest",
// 		Values: Values{"fields": fields},
// 	}
// 	err = req.Get(t.HTTPClient, &value)
// 	return
// }
