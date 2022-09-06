package trac

// Represents a project.
type Project struct {
	Id                      string           `json:"id,omitempty"`
	InternalId              string           `json:"internalId,omitempty"`
	Uuid                    string           `json:"uuid,omitempty"`
	Name                    string           `json:"name,omitempty"`
	ParentProjectId         string           `json:"parentProjectId,omitempty"`
	ParentProjectInternalId string           `json:"parentProjectInternalId,omitempty"`
	ParentProjectName       string           `json:"parentProjectName,omitempty"`
	Archived                bool             `json:"archived,omitempty"`
	Description             string           `json:"description,omitempty"`
	Href                    string           `json:"href,omitempty"`
	WebUrl                  string           `json:"webUrl,omitempty"`
	Links                   *Links           `json:"links,omitempty"`
	ParentProject           *Project         `json:"parentProject,omitempty"`
	ReadOnlyUI              *StateField      `json:"readOnlyUI,omitempty"`
	DefaultTemplate         *BuildType       `json:"defaultTemplate,omitempty"`
	BuildTypes              *BuildTypes      `json:"buildTypes,omitempty"`
	Templates               *BuildTypes      `json:"templates,omitempty"`
	Parameters              *Properties      `json:"parameters,omitempty"`
	VcsRoots                *VcsRoots        `json:"vcsRoots,omitempty"`
	ProjectFeatures         *ProjectFeatures `json:"projectFeatures,omitempty"`
	Projects                *Projects        `json:"projects,omitempty"`
	CloudProfiles           *CloudProfiles   `json:"cloudProfiles,omitempty"`
	Locator                 string           `json:"locator,omitempty"`
}

// Represents a project stub.
type NewProjectDescription struct {
	CopyAllAssociatedSettings bool        `json:"copyAllAssociatedSettings,omitempty"`
	ProjectsIdsMap            *Properties `json:"projectsIdsMap,omitempty"`
	BuildTypesIdsMap          *Properties `json:"buildTypesIdsMap,omitempty"`
	VcsRootsIdsMap            *Properties `json:"vcsRootsIdsMap,omitempty"`
	Name                      string      `json:"name,omitempty"`
	Id                        string      `json:"id,omitempty"`
	SourceProjectLocator      string      `json:"sourceProjectLocator,omitempty"`
	SourceProject             *Project    `json:"sourceProject,omitempty"`
	ParentProject             *Project    `json:"parentProject,omitempty"`
}

// Represents a paginated list of Project entities.
type Projects struct {
	Count    int32     `json:"count,omitempty"`
	Href     string    `json:"href,omitempty"`
	NextHref string    `json:"nextHref,omitempty"`
	PrevHref string    `json:"prevHref,omitempty"`
	Project  []Project `json:"project,omitempty"`
}

// Represents a list of ProjectFeature entities.
type ProjectFeatures struct {
	Count          int32            `json:"count,omitempty"`
	Href           string           `json:"href,omitempty"`
	ProjectFeature []ProjectFeature `json:"projectFeature,omitempty"`
}

// Represents a project feature.
type ProjectFeature struct {
	Id         string      `json:"id,omitempty"`
	Name       string      `json:"name,omitempty"`
	Type_      string      `json:"type,omitempty"`
	Disabled   bool        `json:"disabled,omitempty"`
	Inherited  bool        `json:"inherited,omitempty"`
	Href       string      `json:"href,omitempty"`
	Properties *Properties `json:"properties,omitempty"`
}

// Represents a project state field (as of now, limited to the read-only state of project).
type StateField struct {
	Value     bool `json:"value,omitempty"`
	Inherited bool `json:"inherited,omitempty"`
}
