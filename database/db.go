package database

import (
	"github.com/Somvaded/subscription-management/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectToDBwithRetry attempts to connect to the database with a specified number of retries.
func ConnectToDBwithRetry(dburl string, maxtries int) (*gorm.DB, error){
	var db *gorm.DB
	var err error
	for i:= 0 ;i <maxtries;i++{
		db,err = gorm.Open(postgres.Open(dburl),&gorm.Config{})
		if err == nil {
			db.AutoMigrate(&models.Plan{}, &models.Subscription{})
			return db ,nil
		}
	}
	return nil,err
}