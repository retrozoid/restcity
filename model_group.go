package trac

// Represents a user group.
type Group struct {
	Key          string      `json:"key,omitempty"`
	Name         string      `json:"name,omitempty"`
	Href         string      `json:"href,omitempty"`
	Description  string      `json:"description,omitempty"`
	ParentGroups *Groups     `json:"parent-groups,omitempty"`
	ChildGroups  *Groups     `json:"child-groups,omitempty"`
	Users        *Users      `json:"users,omitempty"`
	Roles        *Roles      `json:"roles,omitempty"`
	Properties   *Properties `json:"properties,omitempty"`
}

// Represents a list of Group entities.
type Groups struct {
	Count int32   `json:"count,omitempty"`
	Group []Group `json:"group,omitempty"`
}
