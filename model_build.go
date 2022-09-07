package trac

func (b Build) WithProperty(p Property) Build {
	if b.Properties == nil {
		b.Properties = new(Properties)
	}
	b.Properties.Add(p)
	return b
}

func (b Build) WithTag(p Tag) Build {
	if b.Tags == nil {
		b.Tags = new(Tags)
	}
	b.Tags.Add(p)
	return b
}

// Represents a build instance.
type Build struct {
	Id                          int64                   `json:"id,omitempty"`
	TaskId                      int64                   `json:"taskId,omitempty"`
	BuildTypeId                 string                  `json:"buildTypeId,omitempty"`
	BuildTypeInternalId         string                  `json:"buildTypeInternalId,omitempty"`
	Number                      string                  `json:"number,omitempty"`
	Status                      string                  `json:"status,omitempty"`
	State                       string                  `json:"state,omitempty"`
	Running                     bool                    `json:"running,omitempty"`
	Composite                   bool                    `json:"composite,omitempty"`
	FailedToStart               bool                    `json:"failedToStart,omitempty"`
	Personal                    bool                    `json:"personal,omitempty"`
	PercentageComplete          int32                   `json:"percentageComplete,omitempty"`
	BranchName                  string                  `json:"branchName,omitempty"`
	DefaultBranch               bool                    `json:"defaultBranch,omitempty"`
	UnspecifiedBranch           bool                    `json:"unspecifiedBranch,omitempty"`
	History                     bool                    `json:"history,omitempty"`
	Pinned                      bool                    `json:"pinned,omitempty"`
	Href                        string                  `json:"href,omitempty"`
	WebUrl                      string                  `json:"webUrl,omitempty"`
	QueuePosition               int32                   `json:"queuePosition,omitempty"`
	LimitedChangesCount         int32                   `json:"limitedChangesCount,omitempty"`
	ArtifactsDirectory          string                  `json:"artifactsDirectory,omitempty"`
	Links                       *Links                  `json:"links,omitempty"`
	StatusText                  string                  `json:"statusText,omitempty"`
	BuildType                   *BuildType              `json:"buildType,omitempty"`
	Comment                     *Comment                `json:"comment,omitempty"`
	Tags                        *Tags                   `json:"tags,omitempty"`
	PinInfo                     *Comment                `json:"pinInfo,omitempty"`
	User                        *User                   `json:"user,omitempty"`
	StartEstimate               string                  `json:"startEstimate,omitempty"`
	WaitReason                  string                  `json:"waitReason,omitempty"`
	FinishEstimate              string                  `json:"finishEstimate,omitempty"`
	DelayedByBuild              *Build                  `json:"delayedByBuild,omitempty"`
	PlannedAgent                *Agent                  `json:"plannedAgent,omitempty"`
	RunningInfo                 *ProgressInfo           `json:"running-info,omitempty"`
	CanceledInfo                *Comment                `json:"canceledInfo,omitempty"`
	QueuedDate                  string                  `json:"queuedDate,omitempty"`
	StartDate                   string                  `json:"startDate,omitempty"`
	FinishDate                  string                  `json:"finishDate,omitempty"`
	Triggered                   *TriggeredBy            `json:"triggered,omitempty"`
	LastChanges                 *Changes                `json:"lastChanges,omitempty"`
	Changes                     *Changes                `json:"changes,omitempty"`
	Revisions                   *Revisions              `json:"revisions,omitempty"`
	VersionedSettingsRevision   *Revision               `json:"versionedSettingsRevision,omitempty"`
	ArtifactDependencyChanges   *BuildChanges           `json:"artifactDependencyChanges,omitempty"`
	Agent                       *Agent                  `json:"agent,omitempty"`
	CompatibleAgents            *Agents                 `json:"compatibleAgents,omitempty"`
	TestOccurrences             *TestOccurrences        `json:"testOccurrences,omitempty"`
	ProblemOccurrences          *ProblemOccurrences     `json:"problemOccurrences,omitempty"`
	Artifacts                   *Files                  `json:"artifacts,omitempty"`
	RelatedIssues               *IssuesUsages           `json:"relatedIssues,omitempty"`
	Properties                  *Properties             `json:"properties,omitempty"`
	ResultingProperties         *Properties             `json:"resultingProperties,omitempty"`
	Attributes                  *Entries                `json:"attributes,omitempty"`
	Statistics                  *Properties             `json:"statistics,omitempty"`
	Metadata                    *Datas                  `json:"metadata,omitempty"`
	SnapshotDependencies        *Builds                 `json:"snapshot-dependencies,omitempty"`
	ArtifactDependencies        *Builds                 `json:"artifact-dependencies,omitempty"`
	CustomArtifactDependencies  *ArtifactDependencies   `json:"custom-artifact-dependencies,omitempty"`
	SettingsHash                string                  `json:"settingsHash,omitempty"`
	CurrentSettingsHash         string                  `json:"currentSettingsHash,omitempty"`
	ModificationId              string                  `json:"modificationId,omitempty"`
	ChainModificationId         string                  `json:"chainModificationId,omitempty"`
	ReplacementIds              *Items                  `json:"replacementIds,omitempty"`
	Related                     *Related                `json:"related,omitempty"`
	TriggeringOptions           *BuildTriggeringOptions `json:"triggeringOptions,omitempty"`
	UsedByOtherBuilds           bool                    `json:"usedByOtherBuilds,omitempty"`
	StatusChangeComment         *Comment                `json:"statusChangeComment,omitempty"`
	VcsLabels                   []VcsLabel              `json:"vcsLabels,omitempty"`
	DetachedFromAgent           bool                    `json:"detachedFromAgent,omitempty"`
	FinishOnAgentDate           string                  `json:"finishOnAgentDate,omitempty"`
	Customized                  bool                    `json:"customized,omitempty"`
	Customization               *Customizations         `json:"customization,omitempty"`
	ChangesCollectingInProgress bool                    `json:"changesCollectingInProgress,omitempty"`
	Locator                     string                  `json:"locator,omitempty"`
}

// Represents a paginated list of Build entities.
type Builds struct {
	Count    int32   `json:"count,omitempty"`
	Href     string  `json:"href,omitempty"`
	NextHref string  `json:"nextHref,omitempty"`
	PrevHref string  `json:"prevHref,omitempty"`
	Build    []Build `json:"build,omitempty"`
}

// Represents a branch on which this build has been started.
type Branch struct {
	Name         string  `json:"name,omitempty"`
	InternalName string  `json:"internalName,omitempty"`
	Default_     bool    `json:"default,omitempty"`
	Unspecified  bool    `json:"unspecified,omitempty"`
	Active       bool    `json:"active,omitempty"`
	LastActivity string  `json:"lastActivity,omitempty"`
	GroupFlag    bool    `json:"groupFlag,omitempty"`
	Builds       *Builds `json:"builds,omitempty"`
}

// Represents a cancel request for the specific build.
type BuildCancelRequest struct {
	Comment        string `json:"comment,omitempty"`
	ReaddIntoQueue bool   `json:"readdIntoQueue,omitempty"`
}

// Represents the dependency/queue settings with which this build has been started.
type BuildTriggeringOptions struct {
	CleanSources                          bool        `json:"cleanSources,omitempty"`
	CleanSourcesInAllDependencies         bool        `json:"cleanSourcesInAllDependencies,omitempty"`
	RebuildAllDependencies                bool        `json:"rebuildAllDependencies,omitempty"`
	RebuildFailedOrIncompleteDependencies bool        `json:"rebuildFailedOrIncompleteDependencies,omitempty"`
	QueueAtTop                            bool        `json:"queueAtTop,omitempty"`
	FreezeSettings                        bool        `json:"freezeSettings,omitempty"`
	TagDependencies                       bool        `json:"tagDependencies,omitempty"`
	RebuildDependencies                   *BuildTypes `json:"rebuildDependencies,omitempty"`
}

// Represents the pinned status of this build.
type PinInfo struct {
	Status  bool     `json:"status,omitempty"`
	Comment *Comment `json:"comment,omitempty"`
}

// Represents a progress estimate of this build.
type ProgressInfo struct {
	PercentageComplete    int32  `json:"percentageComplete,omitempty"`
	ElapsedSeconds        int64  `json:"elapsedSeconds,omitempty"`
	EstimatedTotalSeconds int64  `json:"estimatedTotalSeconds,omitempty"`
	LeftSeconds           int64  `json:"leftSeconds,omitempty"`
	CurrentStageText      string `json:"currentStageText,omitempty"`
	Outdated              bool   `json:"outdated,omitempty"`
	ProbablyHanging       bool   `json:"probablyHanging,omitempty"`
	LastActivityTime      string `json:"lastActivityTime,omitempty"`
	OutdatedReasonBuild   *Build `json:"outdatedReasonBuild,omitempty"`
}

// Represents a single build tag.
type Tag struct {
	Name    string `json:"name,omitempty"`
	Owner   *User  `json:"owner,omitempty"`
	Private bool   `json:"private,omitempty"`
}

// Represents a list of Tag entities.
type Tags struct {
	Count int32 `json:"count,omitempty"`
	Tag   []Tag `json:"tag,omitempty"`
}

func (p *Tags) Add(v Tag) {
	p.Tag = append(p.Tag, v)
	p.Count++
}

// Represents the user/trigger/dependency which caused this build to start.
type TriggeredBy struct {
	Type_       string      `json:"type,omitempty"`
	Details     string      `json:"details,omitempty"`
	Date        string      `json:"date,omitempty"`
	DisplayText string      `json:"displayText,omitempty"`
	RawValue    string      `json:"rawValue,omitempty"`
	User        *User       `json:"user,omitempty"`
	Build       *Build      `json:"build,omitempty"`
	BuildType   *BuildType  `json:"buildType,omitempty"`
	Properties  *Properties `json:"properties,omitempty"`
}

// Represents build customizations (artifact dependency overrides, custom parameters or changesets).
type Customizations struct {
	Parameters           map[string]string `json:"parameters,omitempty"`
	Changes              map[string]string `json:"changes,omitempty"`
	ArtifactDependencies map[string]string `json:"artifactDependencies,omitempty"`
}

// Represents a VCS-side label of this build's sources.
type VcsLabel struct {
	Text            string           `json:"text,omitempty"`
	FailureReason   string           `json:"failureReason,omitempty"`
	Status          string           `json:"status,omitempty"`
	BuildId         int64            `json:"buildId,omitempty"`
	VcsRootInstance *VcsRootInstance `json:"vcs-root-instance,omitempty"`
}

// Represents a list of VcsLabel entities.
type VcsLabels struct {
	Count    int32      `json:"count,omitempty"`
	VcsLabel []VcsLabel `json:"vcsLabel,omitempty"`
}
