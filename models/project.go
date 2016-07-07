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
	"github.com/goadesign/goa/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"time"
)

// This is the Project model
type Project struct {
	ID        uuid.UUID `sql:"type:uuid" gorm:"primary_key"` // This is the ID PK field
	CreatedAt time.Time
	DeletedAt *time.Time
	Name      string
	UpdatedAt time.Time
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Project) TableName() string {
	return "projects"

}

// ProjectDB is the implementation of the storage interface for
// Project.
type ProjectDB struct {
	Db *gorm.DB
}

// NewProjectDB creates a new storage type.
func NewProjectDB(db *gorm.DB) *ProjectDB {
	return &ProjectDB{Db: db}
}

// DB returns the underlying database.
func (m *ProjectDB) DB() interface{} {
	return &m.Db
}

// ProjectStorage represents the storage interface.
type ProjectStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Project, error)
	Get(ctx context.Context, id uuid.UUID) (*Project, error)
	Add(ctx context.Context, project *Project) error
	Update(ctx context.Context, project *Project) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *ProjectDB) TableName() string {
	return "projects"

}

// CRUD Functions

// Get returns a single Project as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *ProjectDB) Get(ctx context.Context, id uuid.UUID) (*Project, error) {
	defer goa.MeasureSince([]string{"goa", "db", "project", "get"}, time.Now())

	var native Project
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &native, err
}

// List returns an array of Project
func (m *ProjectDB) List(ctx context.Context) ([]*Project, error) {
	defer goa.MeasureSince([]string{"goa", "db", "project", "list"}, time.Now())

	var objs []*Project
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *ProjectDB) Add(ctx context.Context, model *Project) error {
	defer goa.MeasureSince([]string{"goa", "db", "project", "add"}, time.Now())

	model.ID = uuid.NewV4()

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Project", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *ProjectDB) Update(ctx context.Context, model *Project) error {
	defer goa.MeasureSince([]string{"goa", "db", "project", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Project", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *ProjectDB) Delete(ctx context.Context, id uuid.UUID) error {
	defer goa.MeasureSince([]string{"goa", "db", "project", "delete"}, time.Now())

	var obj Project

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Project", "error", err.Error())
		return err
	}

	return nil
}
