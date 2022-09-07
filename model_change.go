package trac

type ChangeStatus struct {
	PendingBuildTypes int32 `json:"pendingBuildTypes,omitempty"`
	SuccessfulBuilds  int32 `json:"successfulBuilds,omitempty"`
	RunningBuilds     int32 `json:"runningBuilds,omitempty"`
	FailedBuilds      int32 `json:"failedBuilds,omitempty"`
	FinishedBuilds    int32 `json:"finishedBuilds,omitempty"`
}

// Represents a VCS change (commit).
type Change struct {
	Id                     int64                   `json:"id,omitempty"`
	Version                string                  `json:"version,omitempty"`
	InternalVersion        string                  `json:"internalVersion,omitempty"`
	Username               string                  `json:"username,omitempty"`
	Date                   string                  `json:"date,omitempty"`
	RegistrationDate       string                  `json:"registrationDate,omitempty"`
	Personal               bool                    `json:"personal,omitempty"`
	Href                   string                  `json:"href,omitempty"`
	WebUrl                 string                  `json:"webUrl,omitempty"`
	Comment                string                  `json:"comment,omitempty"`
	User                   *User                   `json:"user,omitempty"`
	Type                   string                  `json:"type,omitempty"`
	SnapshotDependencyLink *SnapshotDependencyLink `json:"snapshotDependencyLink,omitempty"`
	Files                  *FileChanges            `json:"files,omitempty"`
	VcsRootInstance        *VcsRootInstance        `json:"vcsRootInstance,omitempty"`
	ParentChanges          *Changes                `json:"parentChanges,omitempty"`
	ParentRevisions        *Items                  `json:"parentRevisions,omitempty"`
	Attributes             *Properties             `json:"attributes,omitempty"`
	StoresProjectSettings  bool                    `json:"storesProjectSettings,omitempty"`
	Status                 *ChangeStatus           `json:"status,omitempty"`
	CanEditComment         bool                    `json:"canEditComment,omitempty"`
	Locator                string                  `json:"locator,omitempty"`
}

// Represents a paginated list of Change entities.
type Changes struct {
	Count    int32    `json:"count,omitempty"`
	Change   []Change `json:"change,omitempty"`
	Href     string   `json:"href,omitempty"`
	NextHref string   `json:"nextHref,omitempty"`
	PrevHref string   `json:"prevHref,omitempty"`
}

// Represents links to the next or previous build.
type BuildChange struct {
	NextBuild *Build `json:"nextBuild,omitempty"`
	PrevBuild *Build `json:"prevBuild,omitempty"`
}

// Represents a list of BuildChange entities.
type BuildChanges struct {
	Count       int32         `json:"count,omitempty"`
	BuildChange []BuildChange `json:"buildChange,omitempty"`
}

// Represents a revision related to a VCS change.
type Revision struct {
	Version         string           `json:"version,omitempty"`
	InternalVersion string           `json:"internalVersion,omitempty"`
	VcsBranchName   string           `json:"vcsBranchName,omitempty"`
	VcsRootInstance *VcsRootInstance `json:"vcs-root-instance,omitempty"`
	CheckoutRules   string           `json:"checkout-rules,omitempty"`
}

// Represents a list of Revision entities.
type Revisions struct {
	Count    int32      `json:"count,omitempty"`
	Revision []Revision `json:"revision,omitempty"`
}

// Represents the last known repository check status.
type VcsCheckStatus struct {
	Status        string `json:"status,omitempty"`
	RequestorType string `json:"requestorType,omitempty"`
	Timestamp     string `json:"timestamp,omitempty"`
}

// Represents a relation between a VCS root and unique settings set for this root.
type VcsRootInstance struct {
	Id                        string           `json:"id,omitempty"`
	VcsRootId                 string           `json:"vcs-root-id,omitempty"`
	VcsRootInternalId         string           `json:"vcsRootInternalId,omitempty"`
	Name                      string           `json:"name,omitempty"`
	VcsName                   string           `json:"vcsName,omitempty"`
	ModificationCheckInterval int32            `json:"modificationCheckInterval,omitempty"`
	CommitHookMode            bool             `json:"commitHookMode,omitempty"`
	LastVersion               string           `json:"lastVersion,omitempty"`
	LastVersionInternal       string           `json:"lastVersionInternal,omitempty"`
	Href                      string           `json:"href,omitempty"`
	VcsRoot                   *VcsRoot         `json:"vcs-root,omitempty"`
	Status                    *VcsStatus       `json:"status,omitempty"`
	RepositoryState           *RepositoryState `json:"repositoryState,omitempty"`
	Properties                *Properties      `json:"properties,omitempty"`
	RepositoryIdStrings       *Items           `json:"repositoryIdStrings,omitempty"`
	ProjectLocator            string           `json:"projectLocator,omitempty"`
}

// Represents a VCS root.
type VcsRoot struct {
	Id                        string            `json:"id,omitempty"`
	InternalId                string            `json:"internalId,omitempty"`
	Uuid                      string            `json:"uuid,omitempty"`
	Name                      string            `json:"name,omitempty"`
	VcsName                   string            `json:"vcsName,omitempty"`
	ModificationCheckInterval int32             `json:"modificationCheckInterval,omitempty"`
	Href                      string            `json:"href,omitempty"`
	Project                   *Project          `json:"project,omitempty"`
	Properties                *Properties       `json:"properties,omitempty"`
	VcsRootInstances          *VcsRootInstances `json:"vcsRootInstances,omitempty"`
	RepositoryIdStrings       *Items            `json:"repositoryIdStrings,omitempty"`
	ProjectLocator            string            `json:"projectLocator,omitempty"`
	Locator                   string            `json:"locator,omitempty"`
}

// Represents links to the last or previous VCS root check.
type VcsStatus struct {
	Current  *VcsCheckStatus `json:"current,omitempty"`
	Previous *VcsCheckStatus `json:"previous,omitempty"`
}

type SnapshotDependencyLink struct {
	Build           *Build     `json:"build,omitempty"`
	BuildType       *BuildType `json:"buildType,omitempty"`
	BuildTypeBranch string     `json:"buildTypeBranch,omitempty"`
}

// Represents the specific file change (in the scope of the commit).
type FileChange struct {
	BeforeRevision    string `json:"before-revision,omitempty"`
	AfterRevision     string `json:"after-revision,omitempty"`
	ChangeType        string `json:"changeType,omitempty"`
	ChangeTypeComment string `json:"changeTypeComment,omitempty"`
	File              string `json:"file,omitempty"`
	RelativeFile      string `json:"relative-file,omitempty"`
	Directory         bool   `json:"directory,omitempty"`
}

// Represents a list of FileChange entities.
type FileChanges struct {
	Count int32        `json:"count,omitempty"`
	File  []FileChange `json:"file,omitempty"`
}

// Represents the list of the repository Branch entities with their recent revisions.
type RepositoryState struct {
	Timestamp string          `json:"timestamp,omitempty"`
	Count     int32           `json:"count,omitempty"`
	Branch    []BranchVersion `json:"branch,omitempty"`
}
