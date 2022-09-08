package trac

import "time"

// Represents a locator string for filtering Agent entities.
type AgentLocator struct {
	Authorized bool              `json:"authorized,omitempty"` // Is the agent authorized.
	Build      *BuildLocator     `json:"build,omitempty"`      // Build locator.
	Compatible *BuildTypeLocator `json:"compatible,omitempty"` // Compatible build types locator.
	Connected  bool              `json:"connected,omitempty"`  // Is the agent connected.
	Count      int32             `json:"count,omitempty"`      // For paginated calls, how many entities to return per page.
	Enabled    bool              `json:"enabled,omitempty"`    // Is the agent enabled.
	ID         int32             `json:"id,omitempty"`         // Entity ID.
	IP         string            `json:"ip,omitempty"`
	Item       *ItemLocator      `json:"item,omitempty"` // Supply multiple locators and return a union of the results.
	Name       string            `json:"name,omitempty"`
	Parameter  string            `json:"parameter,omitempty"`
	Pool       *AgentPoolLocator `json:"pool,omitempty"`  // Agent pool locator.
	Start      int32             `json:"start,omitempty"` // For paginated calls, from which entity to start rendering the page.
}

type AgentPoolLocator struct {
	Agent   *AgentLocator   `json:"agent,omitempty"` // Pool's agents locator.
	Id      string          `json:"id,omitempty"`
	Item    *ItemLocator    `json:"item,omitempty"` // Supply multiple locators and return a union of the results.
	Name    string          `json:"name,omitempty"`
	Project *ProjectLocator `json:"project,omitempty"` // Pool's associated projects locator.
}

// Represents a locator string for filtering AuditEvent entities.
type AuditLocator struct {
	Action          string            `json:"action,omitempty"`          // Use `$help` to get the full list of supported actions.
	AffectedProject *ProjectLocator   `json:"affectedProject,omitempty"` // Related project locator.
	BuildType       *BuildTypeLocator `json:"buildType,omitempty"`       // Related build type or template locator.
	Count           int32             `json:"count,omitempty"`           // For paginated calls, how many entities to return per page.
	Id              string            `json:"id,omitempty"`
	Item            *ItemLocator      `json:"item,omitempty"`        // Supply multiple locators and return a union of the results.
	LookupLimit     int32             `json:"lookupLimit,omitempty"` // Limit processing to the latest `<lookupLimit>` entities.
	Start           int32             `json:"start,omitempty"`       // For paginated calls, from which entity to start rendering the page.
	SystemAction    string            `json:"systemAction,omitempty"`
	User            string            `json:"user,omitempty"` // Locator of user who caused the audit event.
}

func BuildLocatorID(id int64) BuildLocator {
	return BuildLocator{Id: id}
}

// Represents a locator string for filtering Build entities.
type BuildLocator struct {
	AffectedProject    *ProjectLocator   `json:"affectedProject,omitempty"` // Project (direct or indirect parent) locator.
	Agent              *AgentLocator     `json:"agent,omitempty"`           // Agent locator.
	AgentTypeId        int32             `json:"agentTypeId,omitempty"`     // typeId of agent used to execute build.
	Any                bool              `json:"any,omitempty"`             // State can be any.
	ArtifactDependency string            `json:"artifactDependency,omitempty"`
	Branch             *BranchLocator    `json:"branch,omitempty"`          // Branch locator.
	BuildType          *BuildTypeLocator `json:"buildType,omitempty"`       // Build type locator.
	Canceled           bool              `json:"canceled,omitempty"`        // Is canceled.
	CompatibleAgent    *AgentLocator     `json:"compatibleAgent,omitempty"` // Agent locator.
	Composite          bool              `json:"composite,omitempty"`       // Is composite.
	Count              int32             `json:"count,omitempty"`           // For paginated calls, how many entities to return per page.
	DefaultFilter      bool              `json:"defaultFilter,omitempty"`   // If true, applies default filter which returns only \"normal\" builds (finished builds which are not canceled, not failed-to-start, not personal, and on default branch (in branched build configurations)).
	FailedToStart      bool              `json:"failedToStart,omitempty"`   // Is failed to start.
	FinishDate         string            `json:"finishDate,omitempty"`      // Requires either date or build dimension.
	Finished           bool              `json:"finished,omitempty"`        // Is finished.
	Hanging            bool              `json:"hanging,omitempty"`         // Is hanging.
	History            bool              `json:"history,omitempty"`         // Is history build.
	Id                 int64             `json:"id,omitempty"`              // Entity ID.
	Item               *ItemLocator      `json:"item,omitempty"`            // Supply multiple locators and return a union of the results.
	LookupLimit        int32             `json:"lookupLimit,omitempty"`     // Limit processing to the latest `<lookupLimit>` entities.
	Number             string            `json:"number,omitempty"`
	Personal           bool              `json:"personal,omitempty"` // Is a personal build.
	Pinned             bool              `json:"pinned,omitempty"`   // Is pinned.
	Project            *ProjectLocator   `json:"project,omitempty"`  // Project (direct parent) locator.
	Property           string            `json:"property,omitempty"`
	Queued             bool              `json:"queued,omitempty"`     // Is queued.
	QueuedDate         string            `json:"queuedDate,omitempty"` // Requires either date or build dimension.
	Running            bool              `json:"running,omitempty"`    // Is running.
	SnapshotDependency string            `json:"snapshotDependency,omitempty"`
	Start              int32             `json:"start,omitempty"`     // For paginated calls, from which entity to start rendering the page.
	StartDate          string            `json:"startDate,omitempty"` // Requires either date or build dimension.
	State              string            `json:"state,omitempty"`
	Status             string            `json:"status,omitempty"`
	Tag                *TagLocator       `json:"tag,omitempty"`    // Tag locator.
	TaskId             int32             `json:"taskId,omitempty"` // ID of a build or build promotion.
	User               *UserLocator      `json:"user,omitempty"`   // User locator.
}

// Represents a locator string for filtering Build entities.
type BuildQueueLocator struct {
	Agent     *AgentLocator     `json:"agent,omitempty"`     // Agent locator.
	BuildType *BuildTypeLocator `json:"buildType,omitempty"` // Build type locator.
	Count     int32             `json:"count,omitempty"`     // For paginated calls, how many entities to return per page.
	Id        int32             `json:"id,omitempty"`        // Entity ID.
	Item      *ItemLocator      `json:"item,omitempty"`      // Supply multiple locators and return a union of the results.
	Personal  bool              `json:"personal,omitempty"`  // Is personal.
	Pool      *AgentPoolLocator `json:"pool,omitempty"`      // Agent pool locator.
	Project   *ProjectLocator   `json:"project,omitempty"`   // Project locator.
	Start     int32             `json:"start,omitempty"`     // For paginated calls, from which entity to start rendering the page.
	TaskId    string            `json:"taskId,omitempty"`
	User      *UserLocator      `json:"user,omitempty"` // User locator.
}

func BuildTypeLocatorID(id string) BuildTypeLocator {
	return BuildTypeLocator{Id: id}
}

// Represents a locator string for filtering BuildType entities.
type BuildTypeLocator struct {
	AffectedProject *ProjectLocator         `json:"affectedProject,omitempty"` // Project (direct or indirect parent) locator.
	Build           *BuildLocator           `json:"build,omitempty"`           // Build locator.
	Count           int32                   `json:"count,omitempty"`           // For paginated calls, how many entities to return per page.
	Id              string                  `json:"id,omitempty"`
	InternalId      string                  `json:"internalId,omitempty"`
	Item            *ItemLocator            `json:"item,omitempty"` // Supply multiple locators and return a union of the results.
	Name            string                  `json:"name,omitempty"`
	Paused          bool                    `json:"paused,omitempty"`       // Is paused.
	Project         *ProjectLocator         `json:"project,omitempty"`      // Project (direct parent) locator.
	Start           int32                   `json:"start,omitempty"`        // For paginated calls, from which entity to start rendering the page.
	Template        *BuildTypeLocator       `json:"template,omitempty"`     // Base template locator.
	TemplateFlag    bool                    `json:"templateFlag,omitempty"` // Is a template.
	UUID            string                  `json:"uuid,omitempty"`
	VcsRoot         *VcsRootLocator         `json:"vcsRoot,omitempty"`         // VCS root locator.
	VcsRootInstance *VcsRootInstanceLocator `json:"vcsRootInstance,omitempty"` // VCS root instance locator.
}

// Represents a locator string for filtering Branch entities.
type BranchLocator struct {
	Branched  string            `json:"branched,omitempty"`  // Is feature branch.
	Build     *BuildLocator     `json:"build,omitempty"`     // Build locator.
	BuildType *BuildTypeLocator `json:"buildType,omitempty"` // Build type locator.
	Default_  string            `json:"default,omitempty"`   // Is default branch.
	Item      *ItemLocator      `json:"item,omitempty"`      // Supply multiple locators and return a union of the results.
	Name      string            `json:"name,omitempty"`
	Policy    string            `json:"policy,omitempty"`
}

// Represents a locator string for filtering VcsRootInstance entities.
type VcsRootInstanceLocator struct {
	AffectedProject   string            `json:"affectedProject,omitempty"` // Supply multiple locators and return a union of the results.
	Build             *BuildLocator     `json:"build,omitempty"`           // Build locator.
	BuildType         *BuildTypeLocator `json:"buildType,omitempty"`       // Build type locator.
	Count             int32             `json:"count,omitempty"`           // For paginated calls, how many entities to return per page.
	Id                int32             `json:"id,omitempty"`              // Entity ID.
	Item              *ItemLocator      `json:"item,omitempty"`            // Supply multiple locators and return a union of the results.
	LookupLimit       int32             `json:"lookupLimit,omitempty"`     // Limit processing to the latest `<lookupLimit>` entities.
	Project           *ProjectLocator   `json:"project,omitempty"`         // Project (direct parent) locator.
	Property          string            `json:"property,omitempty"`
	Start             int32             `json:"start,omitempty"`             // For paginated calls, from which entity to start rendering the page.
	Type              string            `json:"type,omitempty"`              // Type of VCS (e.g. jetbrains.git).
	VcsRoot           *VcsRootLocator   `json:"vcsRoot,omitempty"`           // VCS root locator.
	VersionedSettings bool              `json:"versionedSettings,omitempty"` // Is used for versioned settings.
}

// Represents a locator string for filtering CloudImage entities.
type CloudImageLocator struct {
	AffectedProject *ProjectLocator       `json:"affectedProject,omitempty"` // Project (direct or indirect parent) locator.
	Agent           *AgentLocator         `json:"agent,omitempty"`           // Agent locator.
	AgentPool       *AgentPoolLocator     `json:"agentPool,omitempty"`       // Agent pool locator.
	Id              string                `json:"id,omitempty"`
	Instance        *CloudInstanceLocator `json:"instance,omitempty"` // Cloud instance locator.
	Item            *ItemLocator          `json:"item,omitempty"`     // Supply multiple locators and return a union of the results.
	Name            string                `json:"name,omitempty"`
	Profile         *CloudProfileLocator  `json:"profile,omitempty"` // Cloud profile locator.
	Project         *ProjectLocator       `json:"project,omitempty"` // Project locator.
	Property        string                `json:"property,omitempty"`
}

// Represents a locator string for filtering Tag entities.
type TagLocator struct {
	Name    string `json:"name,omitempty"`
	Owner   string `json:"owner,omitempty"`
	Private string `json:"private,omitempty"`
}

// Represents a locator string for filtering Change entities.
type ChangeLocator struct {
	Build           *BuildLocator           `json:"build,omitempty"`     // Build locator.
	BuildType       *BuildTypeLocator       `json:"buildType,omitempty"` // Build type locator.
	Comment         string                  `json:"comment,omitempty"`
	Count           int32                   `json:"count,omitempty"` // For paginated calls, how many entities to return per page.
	File            string                  `json:"file,omitempty"`
	Id              int32                   `json:"id,omitempty"` // Entity ID.
	InternalVersion string                  `json:"internalVersion,omitempty"`
	Item            *ItemLocator            `json:"item,omitempty"`            // Supply multiple locators and return a union of the results.
	Pending         bool                    `json:"pending,omitempty"`         // Is pending.
	Project         *ProjectLocator         `json:"project,omitempty"`         // Project locator.
	SinceChange     string                  `json:"sinceChange,omitempty"`     // Commit SHA since which the changes should be returned.
	Start           int32                   `json:"start,omitempty"`           // For paginated calls, from which entity to start rendering the page.
	User            *UserLocator            `json:"user,omitempty"`            // User locator.
	Username        string                  `json:"username,omitempty"`        // VCS side username.
	VcsRoot         *VcsRootLocator         `json:"vcsRoot,omitempty"`         // VCS root locator.
	VcsRootInstance *VcsRootInstanceLocator `json:"vcsRootInstance,omitempty"` // VCS instance locator.
	Version         string                  `json:"version,omitempty"`
}

// Represents a locator string for filtering VcsRoot entities.
type VcsRootLocator struct {
	AffectedProject *ProjectLocator `json:"affectedProject,omitempty"` // Project (direct or indirect parent) locator.
	Count           int32           `json:"count,omitempty"`           // For paginated calls, how many entities to return per page.
	Id              int32           `json:"id,omitempty"`              // Entity ID.
	InternalId      string          `json:"internalId,omitempty"`
	Item            *ItemLocator    `json:"item,omitempty"`        // Supply multiple locators and return a union of the results.
	LookupLimit     int32           `json:"lookupLimit,omitempty"` // Limit processing to the latest `<lookupLimit>` entities.
	Name            string          `json:"name,omitempty"`
	Project         *ProjectLocator `json:"project,omitempty"` // Project (direct parent) locator.
	Property        string          `json:"property,omitempty"`
	Start           int32           `json:"start,omitempty"` // For paginated calls, from which entity to start rendering the page.
	Type            string          `json:"type,omitempty"`  // Type of VCS (e.g. jetbrains.git).
	Uuid            string          `json:"uuid,omitempty"`
}

// Represents a locator string for filtering CloudInstance entities.
type CloudInstanceLocator struct {
	AffectedProject *ProjectLocator      `json:"affectedProject,omitempty"` // Project (direct or indirect parent) locator.
	Agent           *AgentLocator        `json:"agent,omitempty"`           // Agent locator.
	Id              string               `json:"id,omitempty"`
	Instance        *CloudImageLocator   `json:"instance,omitempty"` // Cloud image locator.
	Item            *ItemLocator         `json:"item,omitempty"`     // Supply multiple locators and return a union of the results.
	NetworkAddress  string               `json:"networkAddress,omitempty"`
	Profile         *CloudProfileLocator `json:"profile,omitempty"` // Cloud profile locator.
	Project         *ProjectLocator      `json:"project,omitempty"` // Project locator.
	Property        string               `json:"property,omitempty"`
}

// Represents a locator string for filtering CloudProfile entities.
type CloudProfileLocator struct {
	AffectedProject *ProjectLocator    `json:"affectedProject,omitempty"` // Project (direct or indirect parent) locator.
	CloudProviderId string             `json:"cloudProviderId,omitempty"`
	Id              string             `json:"id,omitempty"`
	Instance        *CloudImageLocator `json:"instance,omitempty"` // Cloud image locator.
	Item            *ItemLocator       `json:"item,omitempty"`     // Supply multiple locators and return a union of the results.
	Name            string             `json:"name,omitempty"`
	Project         *ProjectLocator    `json:"project,omitempty"` // Project locator.
	Property        string             `json:"property,omitempty"`
}

// Represents a locator string for filtering Investigation entities.
type InvestigationLocator struct {
	AffectedProject   *ProjectLocator   `json:"affectedProject,omitempty"` // Project (direct or indirect parent) locator.
	Assignee          string            `json:"assignee,omitempty"`
	AssignmentProject *ProjectLocator   `json:"assignmentProject,omitempty"` // Project (direct parent) locator.
	BuildType         *BuildTypeLocator `json:"buildType,omitempty"`         // Build type locator.
	Count             int32             `json:"count,omitempty"`             // For paginated calls, how many entities to return per page.
	Item              *ItemLocator      `json:"item,omitempty"`              // Supply multiple locators and return a union of the results.
	LookupLimit       int32             `json:"lookupLimit,omitempty"`       // Limit processing to the latest `<lookupLimit>` entities.
	Problem           *ProblemLocator   `json:"problem,omitempty"`           // Problem locator.
	Reporter          string            `json:"reporter,omitempty"`
	Resolution        string            `json:"resolution,omitempty"`
	SinceDate         *time.Time        `json:"sinceDate,omitempty"` // yyyyMMddTHHmmss+ZZZZ
	Start             int32             `json:"start,omitempty"`     // For paginated calls, from which entity to start rendering the page.
	State             string            `json:"state,omitempty"`
	Test              *TestLocator      `json:"test,omitempty"` // Test locator.
	Type              string            `json:"type,omitempty"`
}

// Represents a locator string for filtering ProblemOccurrence entities.
type ProblemOccurrenceLocator struct {
	AffectedProject       *ProjectLocator `json:"affectedProject,omitempty"`       // Project (direct or indirect parent) locator.
	Build                 *BuildLocator   `json:"build,omitempty"`                 // Build locator.
	Count                 int32           `json:"count,omitempty"`                 // For paginated calls, how many entities to return per page.
	CurrentlyFailing      bool            `json:"currentlyFailing,omitempty"`      // Is currently failing.
	CurrentlyInvestigated bool            `json:"currentlyInvestigated,omitempty"` // Is currently investigated.
	CurrentlyMuted        bool            `json:"currentlyMuted,omitempty"`        // Is currently muted.
	Identity              string          `json:"identity,omitempty"`
	Item                  *ItemLocator    `json:"item,omitempty"`        // Supply multiple locators and return a union of the results.
	LookupLimit           int32           `json:"lookupLimit,omitempty"` // Limit processing to the latest `<lookupLimit>` entities.
	Muted                 bool            `json:"muted,omitempty"`       // Has ever been muted.
	Problem               string          `json:"problem,omitempty"`
	Start                 int32           `json:"start,omitempty"` // For paginated calls, from which entity to start rendering the page.
	Type                  string          `json:"type,omitempty"`
}

// Represents a locator string for filtering Problem entities.
type ProblemLocator struct {
	AffectedProject       *ProjectLocator `json:"affectedProject,omitempty"`       // Project (direct or indirect parent) locator.
	Build                 *BuildLocator   `json:"build,omitempty"`                 // Build locator.
	Count                 int32           `json:"count,omitempty"`                 // For paginated calls, how many entities to return per page.
	CurrentlyFailing      bool            `json:"currentlyFailing,omitempty"`      // Is currently failing.
	CurrentlyInvestigated bool            `json:"currentlyInvestigated,omitempty"` // Is currently investigated.
	CurrentlyMuted        bool            `json:"currentlyMuted,omitempty"`        // Is currently muted.
	Id                    int32           `json:"id,omitempty"`                    // Entity ID.
	Identity              string          `json:"identity,omitempty"`
	Item                  *ItemLocator    `json:"item,omitempty"`        // Supply multiple locators and return a union of the results.
	LookupLimit           int32           `json:"lookupLimit,omitempty"` // Limit processing to the latest `<lookupLimit>` entities.
	Start                 int32           `json:"start,omitempty"`       // For paginated calls, from which entity to start rendering the page.
	Type                  string          `json:"type,omitempty"`
}

// Represents a locator string for filtering Test entities.
type TestLocator struct {
	AffectedProject       *ProjectLocator `json:"affectedProject,omitempty"`       // Project (direct or indirect parent) locator.
	Count                 int32           `json:"count,omitempty"`                 // For paginated calls, how many entities to return per page.
	CurrentlyFailing      bool            `json:"currentlyFailing,omitempty"`      // Is currently failing.
	CurrentlyInvestigated bool            `json:"currentlyInvestigated,omitempty"` // Is currently investigated.
	CurrentlyMuted        bool            `json:"currentlyMuted,omitempty"`        // Is currently muted.
	Id                    int32           `json:"id,omitempty"`                    // Entity ID.
	Item                  *ItemLocator    `json:"item,omitempty"`                  // Supply multiple locators and return a union of the results.
	LookupLimit           int32           `json:"lookupLimit,omitempty"`           // Limit processing to the latest `<lookupLimit>` entities.
	MuteAffected          string          `json:"muteAffected,omitempty"`          // Build type locator (for finding out if this test is affected by mutes in build type).
	Name                  string          `json:"name,omitempty"`
	Start                 int32           `json:"start,omitempty"` // For paginated calls, from which entity to start rendering the page.
}

// Represents a locator string for filtering TestOccurrence entities.
type TestOccurrenceLocator struct {
	AffectedProject       *ProjectLocator   `json:"affectedProject,omitempty"` // Project (direct or indirect parent) locator.
	Branch                string            `json:"branch,omitempty"`
	Build                 *BuildLocator     `json:"build,omitempty"`                 // Build locator.
	BuildType             *BuildTypeLocator `json:"buildType,omitempty"`             // Build type locator.
	Count                 int32             `json:"count,omitempty"`                 // For paginated calls, how many entities to return per page.
	CurrentlyFailing      bool              `json:"currentlyFailing,omitempty"`      // Is currently failing.
	CurrentlyInvestigated bool              `json:"currentlyInvestigated,omitempty"` // Is currently investigated.
	CurrentlyMuted        string            `json:"currentlyMuted,omitempty"`
	Id                    int32             `json:"id,omitempty"`      // Entity ID.
	Ignored               bool              `json:"ignored,omitempty"` // Is ignored.
	IncludePersonal       bool              `json:"includePersonal,omitempty"`
	Item                  *ItemLocator      `json:"item,omitempty"`        // Supply multiple locators and return a union of the results.
	LookupLimit           int32             `json:"lookupLimit,omitempty"` // Limit processing to the latest `<lookupLimit>` entities.
	Muted                 bool              `json:"muted,omitempty"`       // Is muted.
	Name                  string            `json:"name,omitempty"`
	NewFailure            string            `json:"newFailure,omitempty"`
	Start                 int32             `json:"start,omitempty"` // For paginated calls, from which entity to start rendering the page.
	Status                string            `json:"status,omitempty"`
	Test                  *TestLocator      `json:"test,omitempty"` // Test locator.
}

// Represents a locator string for filtering Mute entities.
type MuteLocator struct {
	AffectedProject *ProjectLocator `json:"affectedProject,omitempty"` // Project (direct or indirect parent) locator.
	CreationDate    *time.Time      `json:"creationDate,omitempty"`    // yyyyMMddTHHmmss+ZZZZ
	Id              int32           `json:"id,omitempty"`
	Item            *ItemLocator    `json:"item,omitempty"`     // Supply multiple locators and return a union of the results.
	Problem         *ProblemLocator `json:"problem,omitempty"`  // Problem locator.
	Project         *ProjectLocator `json:"project,omitempty"`  // Project (direct parent) locator.
	Reporter        string          `json:"reporter,omitempty"` // User who muted this test.
	Resolution      string          `json:"resolution,omitempty"`
	Test            *TestLocator    `json:"test,omitempty"` // Test locator.
	Type            string          `json:"type,omitempty"`
	UnmuteDate      *time.Time      `json:"unmuteDate,omitempty"` // yyyyMMddTHHmmss+ZZZZ
}

func ProjectLocatorID(id string) ProjectLocator {
	return ProjectLocator{Id: id}
}

// Represents a locator string for filtering Project entities.
type ProjectLocator struct {
	AffectedProject *ProjectLocator   `json:"affectedProject,omitempty"` // Project (direct or indirect parent) locator.
	Archived        bool              `json:"archived,omitempty"`        // Is archived.
	Build           *BuildLocator     `json:"build,omitempty"`           // Build locator.
	BuildType       *BuildTypeLocator `json:"buildType,omitempty"`       // Build type locator.
	Count           int32             `json:"count,omitempty"`           // For paginated calls, how many entities to return per page.
	DefaultTemplate *BuildTypeLocator `json:"defaultTemplate,omitempty"` // Default template locator.
	Id              string            `json:"id,omitempty"`
	InternalId      string            `json:"internalId,omitempty"`
	Item            *ItemLocator      `json:"item,omitempty"` // Supply multiple locators and return a union of the results.
	Name            string            `json:"name,omitempty"`
	Pool            *AgentPoolLocator `json:"pool,omitempty"`           // Associated agent pool locator.
	Project         *ProjectLocator   `json:"project,omitempty"`        // Project (direct parent) locator.
	ProjectFeature  *ProjectLocator   `json:"projectFeature,omitempty"` // Project feature locator.
	Start           int32             `json:"start,omitempty"`          // For paginated calls, from which entity to start rendering the page.
	UUID            string            `json:"uuid,omitempty"`
	VcsRoot         *VcsRootLocator   `json:"vcsRoot,omitempty"` // VCS root locator.
}

// Represents a locator string for filtering User entities.
type UserLocator struct {
	AffectedGroup *UserGroupLocator `json:"affectedGroup,omitempty"` // User group (direct or indirect parent) locator.
	Email         string            `json:"email,omitempty"`
	Group         *UserGroupLocator `json:"group,omitempty"` // User group (direct parent) locator.
	Id            string            `json:"id,omitempty"`
	Item          *ItemLocator      `json:"item,omitempty"` // Supply multiple locators and return a union of the results.
	LastLogin     *time.Time        `json:"lastLogin,omitempty"`
	Name          string            `json:"name,omitempty"`
	Property      string            `json:"property,omitempty"`
	Role          string            `json:"role,omitempty"` // Role locator. https://www.jetbrains.com/help/teamcity/rest/userlocator.html#Properties
	Username      string            `json:"username,omitempty"`
}

// Represents a locator string for filtering Group entities.
type UserGroupLocator struct {
	Item *ItemLocator `json:"item,omitempty"` // Supply multiple locators and return a union of the results.
	Key  string       `json:"key,omitempty"`
	Name string       `json:"name,omitempty"`
}

type ItemLocator []interface{}
