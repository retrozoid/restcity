package trac

// Represents a plugin.
type Plugin struct {
	Name        string      `json:"name,omitempty"`
	DisplayName string      `json:"displayName,omitempty"`
	Version     string      `json:"version,omitempty"`
	LoadPath    string      `json:"loadPath,omitempty"`
	Parameters  *Properties `json:"parameters,omitempty"`
}

// Represents a list of Plugin entities.
type Plugins struct {
	Count  int32    `json:"count,omitempty"`
	Plugin []Plugin `json:"plugin,omitempty"`
}
