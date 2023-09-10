// database/config.go
package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to the database")
	}
	DB.AutoMigrate(&User{}, &Photo{})
}
