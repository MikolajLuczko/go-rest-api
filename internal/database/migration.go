package database

import (
	"github.com/MikolajLuczko/go-rest-api/internal/transaction"
	"github.com/jinzhu/gorm"
)

// creates the transaction table in our db
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&transaction.Transaction{}); result.Error != nil {
		return result.Error
	}
	return nil
}
