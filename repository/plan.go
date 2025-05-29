package repository

import (
	"context"
	"time"

	"github.com/Somvaded/subscription-management/models"
	"gorm.io/gorm"
)



//createplan creates a new subscription plan in the database
func CreatePlan(ctx context.Context,plan models.Plan,DB *gorm.DB) error {
	err := retry(3, 2*time.Second, func() error {
		return DB.WithContext(ctx).Create(&plan).Error
	})
	return err
}

// GetPlanByID retrieves a subscription plan by its ID from the database
func GetPlanByID(ctx context.Context,planID uint, DB *gorm.DB) (*models.Plan, error) {
	
	var plan models.Plan
	err := retry(3, 2*time.Second, func() error {
		return DB.WithContext(ctx).Where("id = ?", planID).First(&plan).Error
	})
	if err != nil {
		return nil, err
	}
	return &plan, nil
}

// GetPlans retrieves all subscription plans from the database
func GetPlans(ctx context.Context,DB *gorm.DB) ([]models.Plan, error) {
	var plans []models.Plan
	err := retry(3, 2*time.Second, func() error {
		return DB.WithContext(ctx).Find(&plans).Error
	})
	if err != nil {
		return nil, err
	}
	return plans, nil
}
