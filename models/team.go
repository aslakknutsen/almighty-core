//************************************************************************//
// API "alm": Models
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/almighty/almighty-core/design
// --out=$(GOPATH)/src/github.com/almighty/almighty-core
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package models

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"time"
)

// Describes a Team and how users share e.g. Permissions
type Team struct {
	ID          uuid.UUID `sql:"type:uuid default uuid_generate_v4()" gorm:"primary_key"` // This is the ID PK field
	CreatedAt   time.Time
	DeletedAt   *time.Time
	Identities  []Identity   `gorm:"many2many:team_identity"` // many to many Team/Identity
	Name        string       // The display name of the team
	Permissions []Permission `gorm:"many2many:team_permission"` // many to many Team/Permission
	Projects    []Project    `gorm:"many2many:team_project"`    // many to many Team/Project
	UpdatedAt   time.Time
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Team) TableName() string {
	return "teams"

}

// TeamDB is the implementation of the storage interface for
// Team.
type TeamDB struct {
	Db *gorm.DB
}

// NewTeamDB creates a new storage type.
func NewTeamDB(db *gorm.DB) *TeamDB {
	return &TeamDB{Db: db}
}

// DB returns the underlying database.
func (m *TeamDB) DB() interface{} {
	return &m.Db
}

// TeamStorage represents the storage interface.
type TeamStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Team, error)
	Get(ctx context.Context, id uuid.UUID) (*Team, error)
	Add(ctx context.Context, team *Team) error
	Update(ctx context.Context, team *Team) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *TeamDB) TableName() string {
	return "teams"

}

// CRUD Functions

// Get returns a single Team as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *TeamDB) Get(ctx context.Context, id uuid.UUID) (*Team, error) {
	defer goa.MeasureSince([]string{"goa", "db", "team", "get"}, time.Now())

	var native Team
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &native, err
}

// List returns an array of Team
func (m *TeamDB) List(ctx context.Context) ([]*Team, error) {
	defer goa.MeasureSince([]string{"goa", "db", "team", "list"}, time.Now())

	var objs []*Team
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *TeamDB) Add(ctx context.Context, model *Team) error {
	defer goa.MeasureSince([]string{"goa", "db", "team", "add"}, time.Now())

	model.ID = uuid.NewV4()

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Team", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *TeamDB) Update(ctx context.Context, model *Team) error {
	defer goa.MeasureSince([]string{"goa", "db", "team", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Team", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *TeamDB) Delete(ctx context.Context, id uuid.UUID) error {
	defer goa.MeasureSince([]string{"goa", "db", "team", "delete"}, time.Now())

	var obj Team

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Team", "error", err.Error())
		return err
	}

	return nil
}
