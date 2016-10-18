package remoteworkitem

import "github.com/almighty/almighty-core/gormsupport"

// TrackerItem represents a remote tracker item
// Staging area before pushing to work item
type TrackerItem struct {
	gormsupport.Lifecycle
	ID uint64 `gorm:"primary_key"`
	// Remote item ID
	RemoteItemID string
	// the field values
	Item string
	// Batch ID for earch running of tracker query (UUID V4)
	BatchID string
	// FK to trackey query
	TrackerQueryID uint64 `gorm:"ForeignKey:TrackerQuery"`
}
