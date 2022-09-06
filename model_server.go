package trac

// Represents various details of this server including the installation version.
type Server struct {
	Version        string `json:"version,omitempty"`
	VersionMajor   int32  `json:"versionMajor,omitempty"`
	VersionMinor   int32  `json:"versionMinor,omitempty"`
	StartTime      string `json:"startTime,omitempty"`
	CurrentTime    string `json:"currentTime,omitempty"`
	BuildNumber    string `json:"buildNumber,omitempty"`
	BuildDate      string `json:"buildDate,omitempty"`
	InternalId     string `json:"internalId,omitempty"`
	Role           string `json:"role,omitempty"`
	WebUrl         string `json:"webUrl,omitempty"`
	Projects       *Href  `json:"projects,omitempty"`
	VcsRoots       *Href  `json:"vcsRoots,omitempty"`
	Builds         *Href  `json:"builds,omitempty"`
	Users          *Href  `json:"users,omitempty"`
	UserGroups     *Href  `json:"userGroups,omitempty"`
	Agents         *Href  `json:"agents,omitempty"`
	BuildQueue     *Href  `json:"buildQueue,omitempty"`
	AgentPools     *Href  `json:"agentPools,omitempty"`
	Investigations *Href  `json:"investigations,omitempty"`
	Mutes          *Href  `json:"mutes,omitempty"`
	ArtifactsUrl   string `json:"artifactsUrl,omitempty"`
}

// Represents a license key details.
type LicenseKey struct {
	Valid               bool   `json:"valid,omitempty"`
	Active              bool   `json:"active,omitempty"`
	Expired             bool   `json:"expired,omitempty"`
	Obsolete            bool   `json:"obsolete,omitempty"`
	ExpirationDate      string `json:"expirationDate,omitempty"`
	MaintenanceEndDate  string `json:"maintenanceEndDate,omitempty"`
	Type_               string `json:"type,omitempty"`
	Servers             int32  `json:"servers,omitempty"`
	Agents              int32  `json:"agents,omitempty"`
	UnlimitedAgents     bool   `json:"unlimitedAgents,omitempty"`
	BuildTypes          int32  `json:"buildTypes,omitempty"`
	UnlimitedBuildTypes bool   `json:"unlimitedBuildTypes,omitempty"`
	ErrorDetails        string `json:"errorDetails,omitempty"`
	Key                 string `json:"key,omitempty"`
	RawType             string `json:"rawType,omitempty"`
}

// Represents a list of LicenseKey entities.
type LicenseKeys struct {
	Count      int32        `json:"count,omitempty"`
	Href       string       `json:"href,omitempty"`
	LicenseKey []LicenseKey `json:"licenseKey,omitempty"`
}

// Represents license state details (available build configurations, agents, etc.).
type LicensingData struct {
	LicenseUseExceeded         bool         `json:"licenseUseExceeded,omitempty"`
	MaxAgents                  int32        `json:"maxAgents,omitempty"`
	UnlimitedAgents            bool         `json:"unlimitedAgents,omitempty"`
	MaxBuildTypes              int32        `json:"maxBuildTypes,omitempty"`
	UnlimitedBuildTypes        bool         `json:"unlimitedBuildTypes,omitempty"`
	BuildTypesLeft             int32        `json:"buildTypesLeft,omitempty"`
	ServerLicenseType          string       `json:"serverLicenseType,omitempty"`
	ServerEffectiveReleaseDate string       `json:"serverEffectiveReleaseDate,omitempty"`
	AgentsLeft                 int32        `json:"agentsLeft,omitempty"`
	LicenseKeys                *LicenseKeys `json:"licenseKeys,omitempty"`
}
