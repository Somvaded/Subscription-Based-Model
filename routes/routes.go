package routes

import (
	"github.com/Somvaded/subscription-management/controllers"
	"github.com/Somvaded/subscription-management/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {


	
	// Initialize controllers
	planController := controllers.NewPlanController(db)
	subscriptionController := controllers.NewSubscriptionController(db)

	// Middleware for JWT authentication
	router.Use(middleware.JWTAuth())
	
	// Plan routes
	planGroup := router.Group("/plans")
	{
		planGroup.POST("/", planController.CreatePlan)
		planGroup.GET("/:id", planController.GetPlanByID)
		planGroup.GET("/", planController.GetPlans)
	}

	// Subscription routes
	subscriptionGroup := router.Group("/subscriptions")
	{
		subscriptionGroup.POST("/", subscriptionController.CreateSubscription)
		subscriptionGroup.GET("/:id", subscriptionController.GetSubscriptionByUserID)
		subscriptionGroup.POST(":id",subscriptionController.UpdateSubscription)
		subscriptionGroup.DELETE("/:id",subscriptionController.DeleteSubscription)
	}

}