package models

import "time"

type Plan struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Features  string  `json:"features"`
	Duration  int     `json:"duration"`
	CreatedAt time.Time
}

type Subscription struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"user_id"`
	PlanID    uint      `json:"plan_id"`
	Status    string    `json:"status"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Plan      Plan `json:"plan" gorm:"foreignKey:PlanID"`
}


type SubscriptionRequest struct {
	UserID string `json:"user_id" binding:"required,uuid"`
	PlanID uint   `json:"plan_id" binding:"required,gt=0"`
}