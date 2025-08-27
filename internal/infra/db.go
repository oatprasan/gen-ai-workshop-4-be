package infra

import (
	"log"

	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteDB(path string) *gorm.DB {
	dbpath := filepath.Clean(path)
	db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed connect sqlite: %v", err)
	}
	return db
}
