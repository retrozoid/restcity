package trac

// Represents a branch version.
type BranchVersion struct {
	Version      string  `json:"version,omitempty"`
	LastActivity string  `json:"lastActivity,omitempty"`
	GroupFlag    bool    `json:"groupFlag,omitempty"`
	InternalName string  `json:"internalName,omitempty"`
	Builds       *Builds `json:"builds,omitempty"`
	Unspecified  bool    `json:"unspecified,omitempty"`
	Name         string  `json:"name,omitempty"`
	Default_     bool    `json:"default,omitempty"`
	Active       bool    `json:"active,omitempty"`
}

// Represents a list of Branch entities.
type Branches struct {
	Count  int32    `json:"count,omitempty"`
	Href   string   `json:"href,omitempty"`
	Branch []Branch `json:"branch,omitempty"`
}
