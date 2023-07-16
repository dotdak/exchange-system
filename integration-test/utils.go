//go:build integration

package integration

import (
	"gorm.io/gorm"
)

func Teardown(db *gorm.DB) {
	db.Exec("DROP TABLE wagers")
	db.Exec("DROP TABLE buys")
}

func Truncate(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE wagers")
	db.Exec("TRUNCATE TABLE buys")
}
