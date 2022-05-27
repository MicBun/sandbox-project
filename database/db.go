package database

import (
	"sandbox/core"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(core.User{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(core.Order{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(core.ServiceProvider{}); err != nil {
		return err
	}
	return nil
}
