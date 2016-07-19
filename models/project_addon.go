package models

import (
	"time"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
)

// ByName is a gorm filter for a Belongs To relationship.
func ByName(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name = ?", name)

	}
}

// Query expose an open ended Query model
func (u *ProjectDB) Query(funcs ...func(*gorm.DB) *gorm.DB) ([]*Project, error) {
	defer goa.MeasureSince([]string{"goa", "db", "project", "query"}, time.Now())
	var objs []*Project

	err := u.Db.Scopes(funcs...).Table(u.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return objs, nil
}
