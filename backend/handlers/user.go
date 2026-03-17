package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateGoalRequest struct {
	GoalType       string `json:"goal_type" binding:"required,oneof=weight_gain weight_loss"`
	TargetCalories int    `json:"target_calories" binding:"required,min=500,max=10000"`
}

// GET /api/user/profile
func GetProfile(c *gin.Context) {
	userID := c.GetInt("user_id")

	// TODO: query users + user_goals tables by userID
	c.JSON(http.StatusOK, gin.H{
		"id":    userID,
		"name":  "Test User",
		"email": "test@example.com",
		"goal": gin.H{
			"goal_type":       "weight_gain",
			"target_calories": 2800,
		},
	})
}

// PUT /api/user/goal
func UpdateGoal(c *gin.Context) {
	userID := c.GetInt("user_id")

	var req UpdateGoalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: upsert into user_goals table for userID
	_ = userID
	c.JSON(http.StatusOK, gin.H{"message": "goal updated successfully"})
}
