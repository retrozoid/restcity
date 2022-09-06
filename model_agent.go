package trac

// Represents a known agent instance.
type Agent struct {
	Id                     int32                `json:"id,omitempty"`
	Name                   string               `json:"name,omitempty"`
	TypeId                 int32                `json:"typeId,omitempty"`
	Connected              bool                 `json:"connected,omitempty"`
	Enabled                bool                 `json:"enabled,omitempty"`
	Authorized             bool                 `json:"authorized,omitempty"`
	Uptodate               bool                 `json:"uptodate,omitempty"`
	Outdated               bool                 `json:"outdated,omitempty"`
	PluginsOutdated        bool                 `json:"pluginsOutdated,omitempty"`
	JavaOutdated           bool                 `json:"javaOutdated,omitempty"`
	Ip                     string               `json:"ip,omitempty"`
	Protocol               string               `json:"protocol,omitempty"`
	Version                string               `json:"version,omitempty"`
	CurrentAgentVersion    string               `json:"currentAgentVersion,omitempty"`
	LastActivityTime       string               `json:"lastActivityTime,omitempty"`
	IdleSinceTime          string               `json:"idleSinceTime,omitempty"`
	DisconnectionComment   string               `json:"disconnectionComment,omitempty"`
	RegistrationTimestamp  string               `json:"registrationTimestamp,omitempty"`
	Host                   string               `json:"host,omitempty"`
	CpuRank                int32                `json:"cpuRank,omitempty"`
	Port                   int32                `json:"port,omitempty"`
	Href                   string               `json:"href,omitempty"`
	WebUrl                 string               `json:"webUrl,omitempty"`
	Build                  *Build               `json:"build,omitempty"`
	Links                  *Links               `json:"links,omitempty"`
	EnabledInfo            *EnabledInfo         `json:"enabledInfo,omitempty"`
	AuthorizedInfo         *AuthorizedInfo      `json:"authorizedInfo,omitempty"`
	Properties             *Properties          `json:"properties,omitempty"`
	CloudInstance          *CloudInstance       `json:"cloudInstance,omitempty"`
	CloudImage             *CloudImage          `json:"cloudImage,omitempty"`
	Environment            *Environment         `json:"environment,omitempty"`
	Pool                   *AgentPool           `json:"pool,omitempty"`
	CompatibilityPolicy    *CompatibilityPolicy `json:"compatibilityPolicy,omitempty"`
	CompatibleBuildTypes   *BuildTypes          `json:"compatibleBuildTypes,omitempty"`
	IncompatibleBuildTypes *Compatibilities     `json:"incompatibleBuildTypes,omitempty"`
	Builds                 *Builds              `json:"builds,omitempty"`
	Locator                string               `json:"locator,omitempty"`
}

// Represents an agent pool instance.
type AgentPool struct {
	Id           int32     `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Href         string    `json:"href,omitempty"`
	MaxAgents    int32     `json:"maxAgents,omitempty"`
	OwnerProject *Project  `json:"ownerProject,omitempty"`
	Projects     *Projects `json:"projects,omitempty"`
	Agents       *Agents   `json:"agents,omitempty"`
	Locator      string    `json:"locator,omitempty"`
}

// Represents a paginated list of AgentPool entities.
type AgentPools struct {
	Count     int32       `json:"count,omitempty"`
	Href      string      `json:"href,omitempty"`
	NextHref  string      `json:"nextHref,omitempty"`
	PrevHref  string      `json:"prevHref,omitempty"`
	AgentPool []AgentPool `json:"agentPool,omitempty"`
}

// Represents a paginated list of Agent entities.
type Agents struct {
	Count    int32   `json:"count,omitempty"`
	NextHref string  `json:"nextHref,omitempty"`
	PrevHref string  `json:"prevHref,omitempty"`
	Href     string  `json:"href,omitempty"`
	Agent    []Agent `json:"agent,omitempty"`
}

// Represents agent authorization data.
type AuthorizedInfo struct {
	Status  bool     `json:"status,omitempty"`
	Comment *Comment `json:"comment,omitempty"`
}

// Represents a compatibility relation between the agent and build configuration.
type Compatibility struct {
	Compatible        bool          `json:"compatible,omitempty"`
	Agent             *Agent        `json:"agent,omitempty"`
	BuildType         *BuildType    `json:"buildType,omitempty"`
	UnmetRequirements *Requirements `json:"unmetRequirements,omitempty"`
}

// Represents a list of Compatibility entities.
type Compatibilities struct {
	Count         int32           `json:"count,omitempty"`
	Compatibility []Compatibility `json:"compatibility,omitempty"`
}

// Represents a build configuration run policy and included build configurations.
type CompatibilityPolicy struct {
	Policy     string      `json:"policy,omitempty"`
	BuildTypes *BuildTypes `json:"buildTypes,omitempty"`
}

// Represents the current enablement status of the agent.
type EnabledInfo struct {
	Status           bool     `json:"status,omitempty"`
	Comment          *Comment `json:"comment,omitempty"`
	StatusSwitchTime string   `json:"statusSwitchTime,omitempty"`
}

// Represents the details of the agent's OS.
type Environment struct {
	OsType string `json:"osType,omitempty"`
	OsName string `json:"osName,omitempty"`
}

// Represents a list of unmet requirements of a build.
type Requirements struct {
	Description string `json:"description,omitempty"`
}
