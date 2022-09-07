package trac

// Represents a list of ArtifactDependency entities.
type ArtifactDependencies struct {
	Count              int32                `json:"count,omitempty"`
	ArtifactDependency []ArtifactDependency `json:"artifact-dependency,omitempty"`
	Replace            string               `json:"replace,omitempty"`
}

// Represents an artifact dependency relation.
type ArtifactDependency struct {
	Id              string      `json:"id,omitempty"`
	Name            string      `json:"name,omitempty"`
	Type            string      `json:"type,omitempty"`
	Disabled        bool        `json:"disabled,omitempty"`
	Inherited       bool        `json:"inherited,omitempty"`
	Href            string      `json:"href,omitempty"`
	Properties      *Properties `json:"properties,omitempty"`
	SourceBuildType *BuildType  `json:"source-buildType,omitempty"`
}

// Represents a requirement to agent parameters.
type AgentRequirement struct {
	Id         string      `json:"id,omitempty"`
	Name       string      `json:"name,omitempty"`
	Type       string      `json:"type,omitempty"`
	Disabled   bool        `json:"disabled,omitempty"`
	Inherited  bool        `json:"inherited,omitempty"`
	Href       string      `json:"href,omitempty"`
	Properties *Properties `json:"properties,omitempty"`
}

// Represents a list of AgentRequirement entities.
type AgentRequirements struct {
	Count            int32              `json:"count,omitempty"`
	AgentRequirement []AgentRequirement `json:"agent-requirement,omitempty"`
}

// Represents build customization settings of a trigger
type BuildTriggerCustomization struct {
	EnforceCleanCheckout                bool        `json:"enforceCleanCheckout,omitempty"`
	EnforceCleanCheckoutForDependencies bool        `json:"enforceCleanCheckoutForDependencies,omitempty"`
	Parameters                          *Properties `json:"parameters,omitempty"`
}

// Represents a build configuration.
type BuildType struct {
	Id                    string                `json:"id,omitempty"`
	InternalId            string                `json:"internalId,omitempty"`
	Name                  string                `json:"name,omitempty"`
	TemplateFlag          bool                  `json:"templateFlag,omitempty"`
	Type                  string                `json:"type,omitempty"`
	Paused                bool                  `json:"paused,omitempty"`
	Uuid                  string                `json:"uuid,omitempty"`
	Description           string                `json:"description,omitempty"`
	ProjectName           string                `json:"projectName,omitempty"`
	ProjectId             string                `json:"projectId,omitempty"`
	ProjectInternalId     string                `json:"projectInternalId,omitempty"`
	Href                  string                `json:"href,omitempty"`
	WebUrl                string                `json:"webUrl,omitempty"`
	Inherited             bool                  `json:"inherited,omitempty"`
	Links                 *Links                `json:"links,omitempty"`
	Project               *Project              `json:"project,omitempty"`
	Templates             *BuildTypes           `json:"templates,omitempty"`
	Template              *BuildType            `json:"template,omitempty"`
	VcsRootEntries        *VcsRootEntries       `json:"vcs-root-entries,omitempty"`
	Settings              *Properties           `json:"settings,omitempty"`
	Parameters            *Properties           `json:"parameters,omitempty"`
	Steps                 *Steps                `json:"steps,omitempty"`
	Features              *Features             `json:"features,omitempty"`
	Triggers              *Triggers             `json:"triggers,omitempty"`
	SnapshotDependencies  *SnapshotDependencies `json:"snapshot-dependencies,omitempty"`
	ArtifactDependencies  *ArtifactDependencies `json:"artifact-dependencies,omitempty"`
	AgentRequirements     *AgentRequirements    `json:"agent-requirements,omitempty"`
	Branches              *Branches             `json:"branches,omitempty"`
	Builds                *Builds               `json:"builds,omitempty"`
	Investigations        *Investigations       `json:"investigations,omitempty"`
	CompatibleAgents      *Agents               `json:"compatibleAgents,omitempty"`
	VcsRootInstances      *VcsRootInstances     `json:"vcsRootInstances,omitempty"`
	ExternalStatusAllowed bool                  `json:"externalStatusAllowed,omitempty"`
	PauseComment          *Comment              `json:"pauseComment,omitempty"`
	Locator               string                `json:"locator,omitempty"`
}

// Represents a paginated list of BuildType entities.
type BuildTypes struct {
	Count     int32       `json:"count,omitempty"`
	Href      string      `json:"href,omitempty"`
	NextHref  string      `json:"nextHref,omitempty"`
	PrevHref  string      `json:"prevHref,omitempty"`
	BuildType []BuildType `json:"buildType,omitempty"`
}

// Represents a build feature.
type Feature struct {
	Id         string      `json:"id,omitempty"`
	Name       string      `json:"name,omitempty"`
	Type       string      `json:"type,omitempty"`
	Disabled   bool        `json:"disabled,omitempty"`
	Inherited  bool        `json:"inherited,omitempty"`
	Href       string      `json:"href,omitempty"`
	Properties *Properties `json:"properties,omitempty"`
}

// Represents a list of Feature entities.
type Features struct {
	Count   int32     `json:"count,omitempty"`
	Feature []Feature `json:"feature,omitempty"`
}

// Represents a build configuration stub.
type NewBuildTypeDescription struct {
	CopyAllAssociatedSettings bool        `json:"copyAllAssociatedSettings,omitempty"`
	ProjectsIdsMap            *Properties `json:"projectsIdsMap,omitempty"`
	BuildTypesIdsMap          *Properties `json:"buildTypesIdsMap,omitempty"`
	VcsRootsIdsMap            *Properties `json:"vcsRootsIdsMap,omitempty"`
	Name                      string      `json:"name,omitempty"`
	Id                        string      `json:"id,omitempty"`
	SourceBuildTypeLocator    string      `json:"sourceBuildTypeLocator,omitempty"`
	SourceBuildType           *BuildType  `json:"sourceBuildType,omitempty"`
}

// Represents a list of SnapshotDependency entities.
type SnapshotDependencies struct {
	Count              int32                `json:"count,omitempty"`
	SnapshotDependency []SnapshotDependency `json:"snapshot-dependency,omitempty"`
}

// Represents a snapshot dependency relation.
type SnapshotDependency struct {
	Id              string      `json:"id,omitempty"`
	Name            string      `json:"name,omitempty"`
	Type            string      `json:"type,omitempty"`
	Disabled        bool        `json:"disabled,omitempty"`
	Inherited       bool        `json:"inherited,omitempty"`
	Href            string      `json:"href,omitempty"`
	Properties      *Properties `json:"properties,omitempty"`
	SourceBuildType *BuildType  `json:"source-buildType,omitempty"`
}

// Represents a build step.
type Step struct {
	Id         string      `json:"id,omitempty"`
	Name       string      `json:"name,omitempty"`
	Type       string      `json:"type,omitempty"`
	Disabled   bool        `json:"disabled,omitempty"`
	Inherited  bool        `json:"inherited,omitempty"`
	Href       string      `json:"href,omitempty"`
	Properties *Properties `json:"properties,omitempty"`
}

// Represents a list of Step entities.
type Steps struct {
	Count int32  `json:"count,omitempty"`
	Step  []Step `json:"step,omitempty"`
}

// Represents a build trigger.
type Trigger struct {
	Id                 string                     `json:"id,omitempty"`
	Name               string                     `json:"name,omitempty"`
	Type               string                     `json:"type,omitempty"`
	Disabled           bool                       `json:"disabled,omitempty"`
	Inherited          bool                       `json:"inherited,omitempty"`
	Href               string                     `json:"href,omitempty"`
	Properties         *Properties                `json:"properties,omitempty"`
	BuildCustomization *BuildTriggerCustomization `json:"buildCustomization,omitempty"`
}

// Represents a list of Trigger entities.
type Triggers struct {
	Count   int32     `json:"count,omitempty"`
	Trigger []Trigger `json:"trigger,omitempty"`
}

// Represents a list of VcsRootEntry entities.
type VcsRootEntries struct {
	Count        int32          `json:"count,omitempty"`
	VcsRootEntry []VcsRootEntry `json:"vcs-root-entry,omitempty"`
}

// Represents a VCS root assigned to this build configuration.
type VcsRootEntry struct {
	Id            string   `json:"id,omitempty"`
	Inherited     bool     `json:"inherited,omitempty"`
	VcsRoot       *VcsRoot `json:"vcs-root,omitempty"`
	CheckoutRules string   `json:"checkout-rules,omitempty"`
}

// Represents a paginated list of VcsRootInstance entities.
type VcsRootInstances struct {
	Count           int32             `json:"count,omitempty"`
	Href            string            `json:"href,omitempty"`
	NextHref        string            `json:"nextHref,omitempty"`
	PrevHref        string            `json:"prevHref,omitempty"`
	VcsRootInstance []VcsRootInstance `json:"vcs-root-instance,omitempty"`
}

// Represents a paginated list of VcsRoot entities.
type VcsRoots struct {
	Count    int32     `json:"count,omitempty"`
	Href     string    `json:"href,omitempty"`
	NextHref string    `json:"nextHref,omitempty"`
	PrevHref string    `json:"prevHref,omitempty"`
	VcsRoot  []VcsRoot `json:"vcs-root,omitempty"`
}

// Represents an investigation target.
type ProblemTarget struct {
	AnyProblem bool      `json:"anyProblem,omitempty"`
	Tests      *Tests    `json:"tests,omitempty"`
	Problems   *Problems `json:"problems,omitempty"`
}

// Represents an investigation scope.
type ProblemScope struct {
	Project    *Project    `json:"project,omitempty"`
	BuildTypes *BuildTypes `json:"buildTypes,omitempty"`
	BuildType  *BuildType  `json:"buildType,omitempty"`
}

// Represents an investigation of a build problem.
type Investigation struct {
	Id          string         `json:"id,omitempty"`
	State       string         `json:"state,omitempty"`
	Href        string         `json:"href,omitempty"`
	Assignee    *User          `json:"assignee,omitempty"`
	Assignment  *Comment       `json:"assignment,omitempty"`
	Scope       *ProblemScope  `json:"scope,omitempty"`
	Target      *ProblemTarget `json:"target,omitempty"`
	Resolution  *Resolution    `json:"resolution,omitempty"`
	Responsible *User          `json:"responsible,omitempty"`
}

// Represents a paginated list of Investigation entities.
type Investigations struct {
	Count         int32           `json:"count,omitempty"`
	NextHref      string          `json:"nextHref,omitempty"`
	PrevHref      string          `json:"prevHref,omitempty"`
	Href          string          `json:"href,omitempty"`
	Investigation []Investigation `json:"investigation,omitempty"`
}
