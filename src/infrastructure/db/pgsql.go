package db

import (
	"github.com/jinzhu/gorm"
)

// SQLBDer ...
type SQLBDer interface {
	Save(interface{}) error
	Close() error
	Select(query interface{}, args ...interface{}) *gorm.DB
	AutoMigrate(interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Group(args string) *gorm.DB
	Table(table string) *gorm.DB
}

type sqlDB struct {
	db *gorm.DB
}

func (s sqlDB) Select(query interface{}, args ...interface{}) *gorm.DB {
	return s.db.Select(query, args...)
}

// Save ...
func (s sqlDB) Save(data interface{}) error {
	return s.db.Create(data).Error
}

// Close ...
func (s sqlDB) Close() error {
	return s.db.Close()
}

// AutoMigrate ...
func (s sqlDB) AutoMigrate(param interface{}) *gorm.DB {
	return s.db.AutoMigrate(param)
}

// Where ...
func (s sqlDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	return s.db.Where(query, args...)
}

// Group ...
func (s sqlDB) Group(args string) *gorm.DB {
	return s.db.Group(args)
}

// Table ...
func (s sqlDB) Table(table string) *gorm.DB {
	return s.db.Table(table)
}
