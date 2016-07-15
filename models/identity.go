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

// Describes a unique Person with the ALM
type Identity struct {
	ID        uuid.UUID `sql:"type:uuid default uuid_generate_v4()" gorm:"primary_key"` // This is the ID PK field
	CreatedAt time.Time
	DeletedAt *time.Time
	Emails    []User // has many Users
	FullName  string // The fullname of the Identity
	ImageURL  string // The image URL for this Identity
	UpdatedAt time.Time
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Identity) TableName() string {
	return "identities"

}

// IdentityDB is the implementation of the storage interface for
// Identity.
type IdentityDB struct {
	Db *gorm.DB
}

// NewIdentityDB creates a new storage type.
func NewIdentityDB(db *gorm.DB) *IdentityDB {
	return &IdentityDB{Db: db}
}

// DB returns the underlying database.
func (m *IdentityDB) DB() interface{} {
	return &m.Db
}

// IdentityStorage represents the storage interface.
type IdentityStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Identity, error)
	Get(ctx context.Context, id uuid.UUID) (*Identity, error)
	Add(ctx context.Context, identity *Identity) error
	Update(ctx context.Context, identity *Identity) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *IdentityDB) TableName() string {
	return "identities"

}

// CRUD Functions

// Get returns a single Identity as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *IdentityDB) Get(ctx context.Context, id uuid.UUID) (*Identity, error) {
	defer goa.MeasureSince([]string{"goa", "db", "identity", "get"}, time.Now())

	var native Identity
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &native, err
}

// List returns an array of Identity
func (m *IdentityDB) List(ctx context.Context) ([]*Identity, error) {
	defer goa.MeasureSince([]string{"goa", "db", "identity", "list"}, time.Now())

	var objs []*Identity
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *IdentityDB) Add(ctx context.Context, model *Identity) error {
	defer goa.MeasureSince([]string{"goa", "db", "identity", "add"}, time.Now())

	model.ID = uuid.NewV4()

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Identity", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *IdentityDB) Update(ctx context.Context, model *Identity) error {
	defer goa.MeasureSince([]string{"goa", "db", "identity", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Identity", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *IdentityDB) Delete(ctx context.Context, id uuid.UUID) error {
	defer goa.MeasureSince([]string{"goa", "db", "identity", "delete"}, time.Now())

	var obj Identity

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Identity", "error", err.Error())
		return err
	}

	return nil
}
