package trac

import "os"

// Represents a file.
type File struct {
	Name             string    `json:"name,omitempty"`
	FullName         string    `json:"fullName,omitempty"`
	Size             int64     `json:"size,omitempty"`
	ModificationTime string    `json:"modificationTime,omitempty"`
	Href             string    `json:"href,omitempty"`
	Parent           **os.File `json:"parent,omitempty"`
	Content          *Href     `json:"content,omitempty"`
	Children         *Files    `json:"children,omitempty"`
}

// Represents a list of File entities.
type Files struct {
	Count int32      `json:"count,omitempty"`
	Href  string     `json:"href,omitempty"`
	File  []*os.File `json:"file,omitempty"`
}
