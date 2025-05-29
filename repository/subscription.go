package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Somvaded/subscription-management/models"
	"gorm.io/gorm"
)

// retry attempts to execute a function multiple times with a delay between attempts
func retry(attempts int, sleep time.Duration, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		time.Sleep(sleep)
	}
	return fmt.Errorf("all attempts failed %v",err)
}


// CreateSubscription creates a new subscription in the database
func CreateSubscription(ctx context.Context,s *models.Subscription, DB *gorm.DB) error {
	err := retry(3, 2*time.Second, func() error {
		return DB.WithContext(ctx).Create(s).Error
	})
	return err
}


// GetSubscriptionByUserID retrieves a subscription by user ID from the database
func GetSubscriptionByUserID(ctx context.Context,userID string, DB *gorm.DB) (*models.Subscription, error) {
	var s models.Subscription
	err := retry(3, 2*time.Second, func() error {
		return  DB.WithContext(ctx).Preload("Plan").Where("user_id = ?", userID).First(&s).Error
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("subscription not found for user %s: %w", userID, err)
		}
		return nil, fmt.Errorf("error retrieving subscription for user %s: %w", userID, err)
	}
	return &s, nil
}

// UpdateSubscription updates an existing subscription in the database
func UpdateSubscription(ctx context.Context,s *models.Subscription, DB *gorm.DB) error {
	err := retry(3, 2*time.Second, func() error {
		return DB.WithContext(ctx).Save(s).Error
	})
	return err
}


// DeleteSubscription deletes a subscription by user ID from the database
func DeleteSubscription(ctx context.Context,userID string, DB *gorm.DB) error {
	err := retry(3, 2*time.Second, func() error {
		return DB.WithContext(ctx).Where("user_id = ?", userID).Delete(&models.Subscription{}).Error
	})
	return err
}