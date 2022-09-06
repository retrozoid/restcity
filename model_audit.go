package trac

// Represents an audit action.
type AuditAction struct {
	Name    string `json:"name,omitempty"`
	Id      string `json:"id,omitempty"`
	Pattern string `json:"pattern,omitempty"`
}

// Represents an audit event including a user and affected entities.
type AuditEvent struct {
	Id              string           `json:"id,omitempty"`
	Timestamp       string           `json:"timestamp,omitempty"`
	Comment         string           `json:"comment,omitempty"`
	Action          *AuditAction     `json:"action,omitempty"`
	RelatedEntities *RelatedEntities `json:"relatedEntities,omitempty"`
	User            *User            `json:"user,omitempty"`
}

// Represents a paginated list of AuditEvent entities.
type AuditEvents struct {
	Count      int32        `json:"count,omitempty"`
	NextHref   string       `json:"nextHref,omitempty"`
	PrevHref   string       `json:"prevHref,omitempty"`
	Href       string       `json:"href,omitempty"`
	AuditEvent []AuditEvent `json:"auditEvent,omitempty"`
}
