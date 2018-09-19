package database

import (
	"log"

	"github.com/jinzhu/gorm"
	// Extends sql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// Client : gorm database instance
	Client *gorm.DB
)

func main() {}

// NewDatabaseClient : initializes a gorm.DB client
func NewDatabaseClient(driver string, dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(driver, dsn)
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}
	Client = db
	return db, nil
}
