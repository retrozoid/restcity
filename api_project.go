package trac

import (
	"fmt"
	"net/http"
)

type ProjectAPI struct {
	HTTPClient *http.Client
}

// GetAllBranches Get all branches of the matching project.
func (t ProjectAPI) GetAllBranches(projectLocator ProjectLocator, locator string, fields string) (value Branches, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/branches", locatorString(projectLocator)),
		Values: Values{"locator": locator, "fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetAllSubprojectsOrdered Get all subprojects ordered of the matching project.
func (t ProjectAPI) GetAllSubprojectsOrdered(projectLocator ProjectLocator, field string) (value Projects, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/order/projects", locatorString(projectLocator)),
		Values: Values{"field": field},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// SetSubprojectsOrder Update all subprojects order of the matching project.
func (t ProjectAPI) SetSubprojectsOrder(projectLocator ProjectLocator, data Projects, field string) (value Projects, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/order/projects", locatorString(projectLocator)),
		Values: Values{"fields": field},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetBuildParameterValue Get value of build parameter.
func (t ProjectAPI) GetBuildParameterValue(projectLocator ProjectLocator, name string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/projects/%s/parameters/%s/value", locatorString(projectLocator), name),
		Consumer: Coders.String,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// UpdateBuildParameterValue Update value of build parameter.
func (t ProjectAPI) UpdateBuildParameterValue(projectLocator ProjectLocator, name string, data string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/projects/%s/parameters/%s/value", locatorString(projectLocator), name),
		Data:     data,
		Consumer: Coders.String,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetAllProjects Get all projects.
func (t ProjectAPI) GetAllProjects(locator string, fields string) (value Projects, err error) {
	req := Request{
		Path:   "/app/rest/projects",
		Values: Values{"locator": locator, "fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// AddProject Create a new project.
func (t ProjectAPI) AddProject(data NewProjectDescription) (value Project, err error) {
	req := Request{
		Path: "/app/rest/projects",
		Data: data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// GetBuildParameters Get build parameters.
func (t ProjectAPI) GetBuildParameters(projectLocator ProjectLocator, locator string, fields string) (value Properties, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/parameters", locatorString(projectLocator)),
		Values: Values{"locator": locator, "fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// CreateBuildParameter Create a build parameter.
func (t ProjectAPI) CreateBuildParameter(projectLocator ProjectLocator, data Property, fields string) (value Property, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/parameters", locatorString(projectLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// UpdateBuildParameters Update build parameters.
func (t ProjectAPI) UpdateBuildParameters(projectLocator ProjectLocator, data Properties, fields string) (value Properties, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/parameters", locatorString(projectLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// DeleteBuildParameters Delete all build parameters.
func (t ProjectAPI) DeleteBuildParameters(projectLocator ProjectLocator) (err error) {
	req := Request{Path: fmt.Sprintf("/app/rest/projects/%s/parameters", locatorString(projectLocator))}
	err = req.Delete(t.HTTPClient)
	return
}

// GetBuildParameter Get build parameter.
func (t ProjectAPI) GetBuildParameter(projectLocator ProjectLocator, name string, fields string) (value Property, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/parameters/%s", locatorString(projectLocator), name),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// UpdateBuildParameter Update build parameter.
func (t ProjectAPI) UpdateBuildParameter(projectLocator ProjectLocator, name string, data Property, fields string) (value Property, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/parameters/%s", locatorString(projectLocator), name),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// DeleteBuildParameter Delete build parameter.
func (t ProjectAPI) DeleteBuildParameter(projectLocator ProjectLocator, name string) (err error) {
	req := Request{Path: fmt.Sprintf("/app/rest/projects/%s/parameters/%s", locatorString(projectLocator), name)}
	err = req.Delete(t.HTTPClient)
	return
}

// GetDefaultTemplate Get the default template of the matching project.
func (t ProjectAPI) GetDefaultTemplate(projectLocator ProjectLocator, fields string) (value BuildType, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/defaultTemplate", locatorString(projectLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// SetDefaultTemplate Update the default template of the matching project.
func (t ProjectAPI) SetDefaultTemplate(projectLocator ProjectLocator, data BuildType, fields string) (value BuildType, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/defaultTemplate", locatorString(projectLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// RemoveDefaultTemplate Remove the default template from the matching project.
func (t ProjectAPI) RemoveDefaultTemplate(projectLocator ProjectLocator) (err error) {
	req := Request{Path: fmt.Sprintf("/app/rest/projects/%s/defaultTemplate", locatorString(projectLocator))}
	err = req.Delete(t.HTTPClient)
	return
}

// GetAgentPoolsProject Get agent pools appointed to the matching project.
func (t ProjectAPI) GetAgentPoolsProject(projectLocator ProjectLocator, fields string) (value AgentPools, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/agentPools", locatorString(projectLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// AddAgentPoolsProject Assign the matching project to the agent pool.
func (t ProjectAPI) AddAgentPoolsProject(projectLocator ProjectLocator, data AgentPool) (value AgentPool, err error) {
	req := Request{
		Path: fmt.Sprintf("/app/rest/projects/%s/agentPools", locatorString(projectLocator)),
		Data: data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// SetAgentPoolsProject Update agent pools apppointed to the matching project.
func (t ProjectAPI) SetAgentPoolsProject(projectLocator ProjectLocator, body AgentPools, fields string) (value AgentPools, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/agentPools", locatorString(projectLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetProjectSettingsFile Get the settings file of the matching build configuration.
func (t ProjectAPI) GetProjectSettingsFile(projectLocator ProjectLocator) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/projects/%s/settingsFile", locatorString(projectLocator)),
		Consumer: Coders.String,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetProjectField Get a field of the matching project.
func (t ProjectAPI) GetProjectField(projectLocator ProjectLocator, field string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/projects/%s/%s", locatorString(projectLocator), field),
		Consumer: Coders.String,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// SetProjectField Update a field of the matching project.
func (t ProjectAPI) SetProjectField(projectLocator ProjectLocator, field string, data string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/projects/%s/%s", locatorString(projectLocator), field),
		Data:     data,
		Consumer: Coders.String,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetBuildParameterSpecification Get build parameter specification.
func (t ProjectAPI) GetBuildParameterSpecification(projectLocator ProjectLocator, name string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/projects/%s/parameters/%s/type/rawValue", locatorString(projectLocator), name),
		Consumer: Coders.String,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// UpdateBuildParameterSpecification Update build parameter specification.
func (t ProjectAPI) UpdateBuildParameterSpecification(projectLocator ProjectLocator, name string, data string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/projects/%s/parameters/%s/type/rawValue", locatorString(projectLocator), name),
		Data:     data,
		Consumer: Coders.String,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetBuildParameterType Get type of build parameter.
func (t ProjectAPI) GetBuildParameterType(projectLocator ProjectLocator, name string) (value ModelType, err error) {
	req := Request{
		Path: fmt.Sprintf("/app/rest/projects/%s/parameters/%s/type", locatorString(projectLocator), name),
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// UpdateBuildParameterType Update type of build parameter.
func (t ProjectAPI) UpdateBuildParameterType(projectLocator ProjectLocator, name string, data ModelType) (value ModelType, err error) {
	req := Request{
		Path: fmt.Sprintf("/app/rest/projects/%s/parameters/%s/type", locatorString(projectLocator), name),
		Data: data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetSecureValue Get a secure token of the matching project.
func (t ProjectAPI) GetSecureValue(projectLocator ProjectLocator, token string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/projects/%s/secure/values/%s", locatorString(projectLocator), token),
		Consumer: Coders.String,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// RemoveProjectFromAgentPool Unassign a project from the matching agent pool.
func (t ProjectAPI) RemoveProjectFromAgentPool(projectLocator ProjectLocator, agentPoolLocator string) (err error) {
	req := Request{Path: fmt.Sprintf("/app/rest/projects/%s/agentPools/%s", locatorString(projectLocator), agentPoolLocator)}
	err = req.Delete(t.HTTPClient)
	return
}

// UpdateFeatures Update all features.
func (t ProjectAPI) UpdateFeatures(projectLocator ProjectLocator, data ProjectFeatures, fields string) (value ProjectFeatures, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/projectFeatures", locatorString(projectLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetFeatures Get all features.
func (t ProjectAPI) GetFeatures(projectLocator ProjectLocator, locator string, fields string) (value ProjectFeatures, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/projectFeatures", locatorString(projectLocator)),
		Values: Values{"locator": locator, "fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// AddFeature Add a feature.
func (t ProjectAPI) AddFeature(projectLocator ProjectLocator, data ProjectFeature, fields string) (value ProjectFeature, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/projectFeatures", locatorString(projectLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// AddBuildType Add a build configuration to the matching project.
func (t ProjectAPI) AddBuildType(projectLocator ProjectLocator, data NewBuildTypeDescription, fields string) (value BuildType, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/buildTypes", locatorString(projectLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// AddSecureToken Create a new secure token for the matching project.
func (t ProjectAPI) AddSecureToken(projectLocator ProjectLocator, data string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/projects/%s/secure/tokens", locatorString(projectLocator)),
		Data:     data,
		Consumer: Coders.String,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// AddTemplate Add a build configuration template to the matching project.
func (t ProjectAPI) AddTemplate(projectLocator ProjectLocator, data NewBuildTypeDescription, fields string) (value BuildType, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/templates", locatorString(projectLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// GetProjectTemplates Get all templates of the matching project.
func (t ProjectAPI) GetProjectTemplates(projectLocator ProjectLocator, fields string) (value BuildTypes, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/templates", locatorString(projectLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// SetBuildTypesOrder Update all build configurations order of the matching project.
func (t ProjectAPI) SetBuildTypesOrder(projectLocator ProjectLocator, data BuildTypes, field string) (value BuildTypes, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/order/buildTypes", locatorString(projectLocator)),
		Values: Values{"field": field},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetProject Get project matching the locator.
func (t ProjectAPI) GetProject(projectLocator ProjectLocator, fields string) (value Project, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s", locatorString(projectLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// DeleteProject Delete project matching the locator.
func (t ProjectAPI) DeleteProject(projectLocator ProjectLocator) (err error) {
	req := Request{Path: fmt.Sprintf("/app/rest/projects/%s", locatorString(projectLocator))}
	err = req.Delete(t.HTTPClient)
	return
}

// GetProjectParentProject Get the parent project of the matching project.
func (t ProjectAPI) GetProjectParentProject(projectLocator ProjectLocator, fields string) (value Project, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/parentProject", locatorString(projectLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// SetParentProject Update the parent project of the matching project.
func (t ProjectAPI) SetParentProject(projectLocator ProjectLocator, data Project, fields string) (value Project, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/parentProject", locatorString(projectLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetFeature Get a matching feature.
func (t ProjectAPI) GetFeature(projectLocator ProjectLocator, featureLocator string, fields string) (value ProjectFeature, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/projectFeatures/%s", locatorString(projectLocator), featureLocator),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// UpdateFeature Update a matching feature.
func (t ProjectAPI) UpdateFeature(projectLocator ProjectLocator, featureLocator string, data ProjectFeature, fields string) (value ProjectFeature, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/projects/%s/projectFeatures/%s", locatorString(projectLocator), featureLocator),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// DeleteFeature Delete a matching feature.
func (t ProjectAPI) DeleteFeature(projectLocator ProjectLocator, featureLocator string) (err error) {
	req := Request{Path: fmt.Sprintf("/app/rest/projects/%s/projectFeatures/%s", locatorString(projectLocator), featureLocator)}
	err = req.Delete(t.HTTPClient)
	return
}
