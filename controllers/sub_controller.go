package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Somvaded/subscription-management/models"
	"github.com/Somvaded/subscription-management/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubscriptionController struct {
	DB *gorm.DB
}

func NewSubscriptionController(db *gorm.DB) *SubscriptionController {
	return &SubscriptionController{
		DB: db,
	}
}

// CreateSubscription handles the creation of a new subscription
func (sc *SubscriptionController) CreateSubscription(c *gin.Context) {
	var req models.SubscriptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx , cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	plan, err := repository.GetPlanByID(ctx , req.PlanID, sc.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Plan not found "+ err.Error()})
		return
	}
	var sub models.Subscription
	sub.PlanID = plan.ID
	sub.UserID = req.UserID
	sub.Status = "ACTIVE"
	sub.StartDate = time.Now()
	sub.EndDate = sub.StartDate.AddDate(0,0,plan.Duration) 
		
	err = repository.CreateSubscription(ctx , &sub , sc.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subscription "+ err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Subscription created successfully",
	})
	
}

// GetSubscriptionByUserID retrieves a subscription by user ID
func (sc *SubscriptionController) GetSubscriptionByUserID(c *gin.Context) {
	userID := c.Param("userId")
	if  _, err := uuid.Parse(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}
	ctx ,cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	s, err := repository.GetSubscriptionByUserID(ctx,userID, sc.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		return
	}
	c.JSON(http.StatusOK, s)
}

// UpdateSubscription updates an existing subscription
func (sc *SubscriptionController) UpdateSubscription(c *gin.Context) {
	userID := c.Param("userId")
	if  _, err := uuid.Parse(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}
	ctx ,cancel := context.WithTimeout(c.Request.Context(), 12*time.Second)
	defer cancel()
	s, err := repository.GetSubscriptionByUserID(ctx , userID,sc.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		return
	}
	var payload models.SubscriptionRequest
	c.ShouldBindJSON(&payload)
	s.PlanID = payload.PlanID
	s.UpdatedAt = time.Now()
	s.Status = "ACTIVE"
	plan, err := repository.GetPlanByID(ctx ,s.PlanID, sc.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Plan not found "+ err.Error()})
		return
	}
	s.EndDate = time.Now().AddDate(0, 0, plan.Duration)
	err = repository.UpdateSubscription(ctx, s, sc.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update subscription "+ err.Error()})
		return
	}
	c.JSON(http.StatusOK, s)
}

// DeleteSubscription cancels a subscription by user ID
func (sc *SubscriptionController) DeleteSubscription(c *gin.Context) {
	userID := c.Param("userId")
	if  _, err := uuid.Parse(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}
	ctx ,cancel := context.WithTimeout(c.Request.Context(), 12*time.Second)
	defer cancel()
	repository.DeleteSubscription(ctx , userID,sc.DB)
	c.JSON(http.StatusOK, gin.H{"status": "cancelled"})
}
