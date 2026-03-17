package handlers

import (
	"net/http"
	"strconv"

	"fitmeals/models"

	"github.com/gin-gonic/gin"
)

// GET /api/recipes
func ListRecipes(c *gin.Context) {
	goal := c.Query("goal")
	search := c.Query("search")

	recipes, err := models.ListRecipes(goal, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch recipes"})
		return
	}

	// If the caller is authenticated, mark which recipes they've favorited
	if userID, exists := c.Get("user_id"); exists {
		uid := userID.(int)
		for i := range recipes {
			favorited, err := models.IsFavorited(uid, recipes[i].ID)
			if err == nil {
				recipes[i].IsFavorited = favorited
			}
		}
	}

	// Return empty array instead of null when there are no results
	if recipes == nil {
		recipes = []models.Recipe{}
	}

	c.JSON(http.StatusOK, recipes)
}

// GET /api/recipes/:id
func GetRecipe(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid recipe id"})
		return
	}

	recipe, err := models.GetRecipeByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch recipe"})
		return
	}
	if recipe == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
		return
	}

	// Mark as favorited if the caller is authenticated
	if userID, exists := c.Get("user_id"); exists {
		favorited, err := models.IsFavorited(userID.(int), recipe.ID)
		if err == nil {
			recipe.IsFavorited = favorited
		}
	}

	c.JSON(http.StatusOK, recipe)
}
