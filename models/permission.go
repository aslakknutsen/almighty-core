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

// Describes a single permissions and it's relation to a team
type Permission struct {
	ID        uuid.UUID `sql:"type:uuid default uuid_generate_v4()" gorm:"primary_key"` // This is the ID PK field
	CreatedAt time.Time
	DeletedAt *time.Time
	Name      string // The string value/name of the permission used in auth token
	UpdatedAt time.Time
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Permission) TableName() string {
	return "permissions"

}

// PermissionDB is the implementation of the storage interface for
// Permission.
type PermissionDB struct {
	Db *gorm.DB
}

// NewPermissionDB creates a new storage type.
func NewPermissionDB(db *gorm.DB) *PermissionDB {
	return &PermissionDB{Db: db}
}

// DB returns the underlying database.
func (m *PermissionDB) DB() interface{} {
	return &m.Db
}

// PermissionStorage represents the storage interface.
type PermissionStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Permission, error)
	Get(ctx context.Context, id uuid.UUID) (*Permission, error)
	Add(ctx context.Context, permission *Permission) error
	Update(ctx context.Context, permission *Permission) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *PermissionDB) TableName() string {
	return "permissions"

}

// CRUD Functions

// Get returns a single Permission as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *PermissionDB) Get(ctx context.Context, id uuid.UUID) (*Permission, error) {
	defer goa.MeasureSince([]string{"goa", "db", "permission", "get"}, time.Now())

	var native Permission
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &native, err
}

// List returns an array of Permission
func (m *PermissionDB) List(ctx context.Context) ([]*Permission, error) {
	defer goa.MeasureSince([]string{"goa", "db", "permission", "list"}, time.Now())

	var objs []*Permission
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *PermissionDB) Add(ctx context.Context, model *Permission) error {
	defer goa.MeasureSince([]string{"goa", "db", "permission", "add"}, time.Now())

	model.ID = uuid.NewV4()

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Permission", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *PermissionDB) Update(ctx context.Context, model *Permission) error {
	defer goa.MeasureSince([]string{"goa", "db", "permission", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Permission", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *PermissionDB) Delete(ctx context.Context, id uuid.UUID) error {
	defer goa.MeasureSince([]string{"goa", "db", "permission", "delete"}, time.Now())

	var obj Permission

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Permission", "error", err.Error())
		return err
	}

	return nil
}
