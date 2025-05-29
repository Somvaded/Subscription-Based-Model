package controllers

import (
	"context"
	"strconv"
	"time"

	"github.com/Somvaded/subscription-management/models"
	"github.com/Somvaded/subscription-management/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PlanController struct {
	Db *gorm.DB
}
func NewPlanController(db *gorm.DB) *PlanController {
	return &PlanController{
		Db: db,
	}
}
// CreatePlan handles the creation of a new subscription plan

func(pc *PlanController) CreatePlan(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Price    float64 `json:"price" binding:"required"`
		Features string `json:"features" binding:"required"`
		Duration int    `json:"duration" binding:"required"`
	}

	ctx , cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var plan models.Plan
	plan.Name = req.Name
	plan.Price = req.Price
	plan.Features = req.Features
	plan.Duration = req.Duration
	if err := repository.CreatePlan(ctx, plan, pc.Db); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create plan: " + err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "Plan created successfully", "plan": plan})
}

// GetPlanByID retrieves a subscription plan by its ID
func(pc *PlanController) GetPlanByID(c *gin.Context) {
	planID := c.Param("id")
	id , err := strconv.ParseUint(planID, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid plan ID format"})
		return
	}

	ctx , cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	plan , err := repository.GetPlanByID(ctx,uint(id),pc.Db)
	 if err != nil {
		c.JSON(404, gin.H{"error": "Plan not found"})
		return
	}
	c.JSON(200, plan)
}


// GetPlans retrieves all available subscription plans
func(pc *PlanController) GetPlans(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	plans, err := repository.GetPlans(ctx, pc.Db)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve plans: " + err.Error()})
		return
	}
	c.JSON(200, plans)
}