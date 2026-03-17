package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /api/favorites
func ListFavorites(c *gin.Context) {
	userID := c.GetInt("user_id")

	// TODO: join user_favorites with recipes + recipe_nutrition where user_id = userID
	_ = userID
	c.JSON(http.StatusOK, []gin.H{
		{
			"id":        1,
			"name":      "High-Protein Chicken Bowl",
			"goal_type": "weight_gain",
			"nutrition": gin.H{
				"calories":  650,
				"protein_g": 55.0,
			},
			"is_favorited": true,
		},
	})
}

// POST /api/favorites/:recipeId
func AddFavorite(c *gin.Context) {
	userID := c.GetInt("user_id")
	recipeID := c.Param("recipeId")

	// TODO: insert into user_favorites (user_id, recipe_id), ignore duplicate
	_ = userID
	_ = recipeID
	c.JSON(http.StatusCreated, gin.H{"message": "recipe added to favorites"})
}

// DELETE /api/favorites/:recipeId
func RemoveFavorite(c *gin.Context) {
	userID := c.GetInt("user_id")
	recipeID := c.Param("recipeId")

	// TODO: delete from user_favorites where user_id = userID and recipe_id = recipeID
	_ = userID
	_ = recipeID
	c.JSON(http.StatusOK, gin.H{"message": "recipe removed from favorites"})
}
