package models

import (
	"github.com/almighty/almighty-core/compare"
	"github.com/almighty/almighty-core/gormsupport"
)

// WorkItem represents a work item as it is stored in the database
type WorkItem struct {
	gormsupport.Lifecycle
	ID uint64 `gorm:"primary_key"`
	// Id of the type of this work item
	Type string
	// Version for optimistic concurrency control
	Version int
	// the field values
	Fields Fields `sql:"type:jsonb"`
}

// Ensure WorkItem implements the Equaler interface
var _ compare.Equaler = WorkItem{}
var _ compare.Equaler = (*WorkItem)(nil)

// Equal returns true if two WorkItem objects are equal; otherwise false is returned.
func (self WorkItem) Equal(u compare.Equaler) bool {
	other, ok := u.(WorkItem)
	if !ok {
		return false
	}
	if !self.Lifecycle.Equal(other.Lifecycle) {
		return false
	}
	if self.Type != other.Type {
		return false
	}
	if self.Version != other.Version {
		return false
	}
	return self.Fields.Equal(other.Fields)
}
