package trac

import "time"

// Represents a user.
type User struct {
	Username    string       `json:"username,omitempty"`
	Name        string       `json:"name,omitempty"`
	Id          int64        `json:"id,omitempty"`
	Email       string       `json:"email,omitempty"`
	LastLogin   string       `json:"lastLogin,omitempty"`
	Password    string       `json:"password,omitempty"`
	HasPassword bool         `json:"hasPassword,omitempty"`
	Realm       string       `json:"realm,omitempty"`
	Href        string       `json:"href,omitempty"`
	Properties  *Properties  `json:"properties,omitempty"`
	Roles       *Roles       `json:"roles,omitempty"`
	Groups      *Groups      `json:"groups,omitempty"`
	Locator     string       `json:"locator,omitempty"`
	Avatars     *UserAvatars `json:"avatars,omitempty"`
	Enabled2FA  bool         `json:"enabled2FA,omitempty"`
}

// Represents a list of User entities.
type Users struct {
	Count int32  `json:"count,omitempty"`
	User  []User `json:"user,omitempty"`
}

// Represents a group of links to the user's avatars
type UserAvatars struct {
	UrlToSize20 string `json:"urlToSize20,omitempty"`
	UrlToSize28 string `json:"urlToSize28,omitempty"`
	UrlToSize32 string `json:"urlToSize32,omitempty"`
	UrlToSize40 string `json:"urlToSize40,omitempty"`
	UrlToSize56 string `json:"urlToSize56,omitempty"`
	UrlToSize64 string `json:"urlToSize64,omitempty"`
	UrlToSize80 string `json:"urlToSize80,omitempty"`
}

// Represents a relation between the specific permission and a project.
type PermissionAssignment struct {
	Permission *Permission `json:"permission,omitempty"`
	Project    *Project    `json:"project,omitempty"`
}

// Represents a list of PermissionAssignment entities.
type PermissionAssignments struct {
	Count                int32                  `json:"count,omitempty"`
	PermissionAssignment []PermissionAssignment `json:"permissionAssignment,omitempty"`
}

// Represents permission restrictions of an authentication token.
type PermissionRestriction struct {
	IsGlobalScope bool        `json:"isGlobalScope,omitempty"`
	Project       *Project    `json:"project,omitempty"`
	Permission    *Permission `json:"permission,omitempty"`
}

// Represents a list of PermissionRestriction entities.
type PermissionRestrictions struct {
	Count                 int32                   `json:"count,omitempty"`
	PermissionRestriction []PermissionRestriction `json:"permissionRestriction,omitempty"`
}

// Represents a permission.
type Permission struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Global bool   `json:"global,omitempty"`
}

// Represents a user/group role.
type Role struct {
	RoleId string `json:"roleId,omitempty"`
	Scope  string `json:"scope,omitempty"`
	Href   string `json:"href,omitempty"`
}

// Represents a list of Role entities.
type Roles struct {
	Role []Role `json:"role,omitempty"`
}

// Represents an authentication token.
type Token struct {
	Name                   string                  `json:"name,omitempty"`
	CreationTime           time.Time               `json:"creationTime,omitempty"`
	Value                  string                  `json:"value,omitempty"`
	ExpirationTime         time.Time               `json:"expirationTime,omitempty"`
	PermissionRestrictions *PermissionRestrictions `json:"permissionRestrictions,omitempty"`
}

// Represents a list of Token entities.
type Tokens struct {
	Count int32   `json:"count,omitempty"`
	Token []Token `json:"token,omitempty"`
}
