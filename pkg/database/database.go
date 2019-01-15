package database

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	// sqlite driver for test
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	dbTypeKey = "database.type"
	dbNameKey = "database.name"

	dbTypeDefault = "sqlite3"
	dbNameDefault = "test.db"

	// ErrDBNotFound means no matched record
	ErrDBNotFound = "record not found"
)

var db *gorm.DB

func init() {
	viper.SetDefault(dbTypeKey, dbTypeDefault)
	viper.SetDefault(dbNameKey, dbNameDefault)
}

// GetDB returns database connection
func GetDB() *gorm.DB {
	var err error

	if db != nil {
		return db
	}

	dbType := viper.GetString(dbTypeKey)
	dbName := viper.GetString(dbNameKey)

	db, err = gorm.Open(dbType, dbName)
	if err != nil {
		panic(err)
	}
	return db
}
