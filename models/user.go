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

// Describes a User(single email) in any system
type User struct {
	ID         uuid.UUID `sql:"type:uuid default uuid_generate_v4()" gorm:"primary_key"` // This is the ID PK field
	CreatedAt  time.Time
	DeletedAt  *time.Time
	Email      string    `sql:"unique_index"` // This is the unique email field
	IdentityID uuid.UUID `sql:"type:uuid"`    // Belongs To Identity
	UpdatedAt  time.Time
	Identity   Identity
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m User) TableName() string {
	return "users"

}

// UserDB is the implementation of the storage interface for
// User.
type UserDB struct {
	Db *gorm.DB
}

// NewUserDB creates a new storage type.
func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{Db: db}
}

// DB returns the underlying database.
func (m *UserDB) DB() interface{} {
	return &m.Db
}

// UserStorage represents the storage interface.
type UserStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*User, error)
	Get(ctx context.Context, id uuid.UUID) (*User, error)
	Add(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *UserDB) TableName() string {
	return "users"

}

// Belongs To Relationships

// UserFilterByIdentity is a gorm filter for a Belongs To relationship.
func UserFilterByIdentity(identityID uuid.UUID, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if identityID > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("identity_id = ?", identityID)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single User as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *UserDB) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "get"}, time.Now())

	var native User
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &native, err
}

// List returns an array of User
func (m *UserDB) List(ctx context.Context) ([]*User, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "list"}, time.Now())

	var objs []*User
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *UserDB) Add(ctx context.Context, model *User) error {
	defer goa.MeasureSince([]string{"goa", "db", "user", "add"}, time.Now())

	model.ID = uuid.NewV4()

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding User", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *UserDB) Update(ctx context.Context, model *User) error {
	defer goa.MeasureSince([]string{"goa", "db", "user", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating User", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *UserDB) Delete(ctx context.Context, id uuid.UUID) error {
	defer goa.MeasureSince([]string{"goa", "db", "user", "delete"}, time.Now())

	var obj User

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting User", "error", err.Error())
		return err
	}

	return nil
}
