package handlers

import (
	"net/http"
	"strconv"

	"fitmeals/models"

	"github.com/gin-gonic/gin"
)

// GET /api/favorites
func ListFavorites(c *gin.Context) {
	userID := c.GetInt("user_id")

	recipes, err := models.ListFavorites(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch favorites"})
		return
	}

	if recipes == nil {
		recipes = []models.Recipe{}
	}

	c.JSON(http.StatusOK, recipes)
}

// POST /api/favorites/:recipeId
func AddFavorite(c *gin.Context) {
	userID := c.GetInt("user_id")

	recipeID, err := strconv.Atoi(c.Param("recipeId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid recipe id"})
		return
	}

	// Verify the recipe exists before inserting
	recipe, err := models.GetRecipeByID(recipeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}
	if recipe == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
		return
	}

	if err := models.AddFavorite(userID, recipeID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not add favorite"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "recipe added to favorites"})
}

// DELETE /api/favorites/:recipeId
func RemoveFavorite(c *gin.Context) {
	userID := c.GetInt("user_id")

	recipeID, err := strconv.Atoi(c.Param("recipeId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid recipe id"})
		return
	}

	if err := models.RemoveFavorite(userID, recipeID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not remove favorite"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "recipe removed from favorites"})
}
