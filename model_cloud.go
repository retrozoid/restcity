package trac

// Represents a cloud instance image saved with a profile.
type CloudImage struct {
	Id                  string          `json:"id,omitempty"`
	Name                string          `json:"name,omitempty"`
	Href                string          `json:"href,omitempty"`
	Profile             *CloudProfile   `json:"profile,omitempty"`
	Instances           *CloudInstances `json:"instances,omitempty"`
	ErrorMessage        string          `json:"errorMessage,omitempty"`
	AgentTypeId         int32           `json:"agentTypeId,omitempty"`
	AgentPoolId         int32           `json:"agentPoolId,omitempty"`
	OperatingSystemName string          `json:"operatingSystemName,omitempty"`
	Locator             string          `json:"locator,omitempty"`
}

// Represents a paginated list of CloudImage entities.
type CloudImages struct {
	Count      int32        `json:"count,omitempty"`
	NextHref   string       `json:"nextHref,omitempty"`
	PrevHref   string       `json:"prevHref,omitempty"`
	Href       string       `json:"href,omitempty"`
	CloudImage []CloudImage `json:"cloudImage,omitempty"`
}

// Represents a running cloud instance.
type CloudInstance struct {
	Id             string      `json:"id,omitempty"`
	Name           string      `json:"name,omitempty"`
	State          string      `json:"state,omitempty"`
	StartDate      string      `json:"startDate,omitempty"`
	NetworkAddress string      `json:"networkAddress,omitempty"`
	Href           string      `json:"href,omitempty"`
	Image          *CloudImage `json:"image,omitempty"`
	Agent          *Agent      `json:"agent,omitempty"`
	ErrorMessage   string      `json:"errorMessage,omitempty"`
}

// Represents a paginated list of CloudInstance entities.
type CloudInstances struct {
	Count         int32           `json:"count,omitempty"`
	NextHref      string          `json:"nextHref,omitempty"`
	PrevHref      string          `json:"prevHref,omitempty"`
	Href          string          `json:"href,omitempty"`
	CloudInstance []CloudInstance `json:"cloudInstance,omitempty"`
}

// Represents a cloud agent profile.
type CloudProfile struct {
	Id              string       `json:"id,omitempty"`
	Name            string       `json:"name,omitempty"`
	CloudProviderId string       `json:"cloudProviderId,omitempty"`
	Href            string       `json:"href,omitempty"`
	Project         *Project     `json:"project,omitempty"`
	Images          *CloudImages `json:"images,omitempty"`
}

// Represents a paginated list of CloudProfile entities.
type CloudProfiles struct {
	Count        int32          `json:"count,omitempty"`
	NextHref     string         `json:"nextHref,omitempty"`
	PrevHref     string         `json:"prevHref,omitempty"`
	Href         string         `json:"href,omitempty"`
	CloudProfile []CloudProfile `json:"cloudProfile,omitempty"`
}
