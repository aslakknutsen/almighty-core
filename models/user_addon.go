package models

import (
	"time"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
)

// ByEmails is a gorm filter for a Belongs To relationship.
func ByEmails(emails []string) func(db *gorm.DB) *gorm.DB {
	if len(emails) > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("email in (?)", emails)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// WithIdentity is a gorm filter for preloading the Identity relationship.
func WithIdentity() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Identity")

	}
}

// Query expose an open ended Query model
func (u *UserDB) Query(funcs ...func(*gorm.DB) *gorm.DB) ([]*User, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "query"}, time.Now())
	var objs []*User

	err := u.Db.Scopes(funcs...).Table(u.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return objs, nil

}
