package handlers

import (
	"net/http"

	"fitmeals/models"

	"github.com/gin-gonic/gin"
)

type UpdateGoalRequest struct {
	GoalType       string `json:"goal_type" binding:"required,oneof=weight_gain weight_loss"`
	TargetCalories int    `json:"target_calories" binding:"required,min=500,max=10000"`
}

// GET /api/user/profile
func GetProfile(c *gin.Context) {
	userID := c.GetInt("user_id")

	user, err := models.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	goal, err := models.GetUserGoal(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"goal":  goal, // nil when not set yet — frontend handles this
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

	if err := models.UpsertUserGoal(userID, req.GoalType, req.TargetCalories); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update goal"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "goal updated successfully"})
}
