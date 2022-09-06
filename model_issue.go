package trac

// Represents an issue related to the specific change.
type Issue struct {
	Url string `json:"url,omitempty"`
	Id  string `json:"id,omitempty"`
}

// Represents a relation between the issue and the Changes entity.
type IssueUsage struct {
	Changes *Changes `json:"changes,omitempty"`
	Issue   *Issue   `json:"issue,omitempty"`
}

// Represents a list of IssueUsage entities.
type IssuesUsages struct {
	IssueUsage []IssueUsage `json:"issueUsage,omitempty"`
	Count      int32        `json:"count,omitempty"`
	Href       string       `json:"href,omitempty"`
}

// Represents a list of Issue entities.
type Issues struct {
	Issues []Issue `json:"issues,omitempty"`
}
