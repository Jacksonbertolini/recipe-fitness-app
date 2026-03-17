package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Mock recipe data matching the DB schema
var mockRecipes = []gin.H{
	{
		"id":                1,
		"name":              "High-Protein Chicken Bowl",
		"description":       "A calorie-dense bowl packed with lean protein and complex carbs.",
		"goal_type":         "weight_gain",
		"prep_time_minutes": 10,
		"cook_time_minutes": 25,
		"servings":          2,
		"image_url":         "",
		"ingredients": []gin.H{
			{"name": "chicken breast", "amount": "300", "unit": "g"},
			{"name": "brown rice", "amount": "200", "unit": "g"},
			{"name": "olive oil", "amount": "2", "unit": "tbsp"},
		},
		"instructions": []string{
			"Cook brown rice according to package instructions.",
			"Season chicken with salt, pepper, and garlic powder.",
			"Heat olive oil in a pan over medium-high heat.",
			"Cook chicken 6-7 minutes per side until cooked through.",
			"Slice chicken and serve over rice.",
		},
		"nutrition": gin.H{
			"calories":  650,
			"protein_g": 55.0,
			"carbs_g":   70.0,
			"fats_g":    14.0,
			"fiber_g":   3.0,
		},
		"is_favorited": false,
	},
	{
		"id":                2,
		"name":              "Light Greek Salad",
		"description":       "A fresh, low-calorie salad with lean protein to support weight loss.",
		"goal_type":         "weight_loss",
		"prep_time_minutes": 15,
		"cook_time_minutes": 0,
		"servings":          1,
		"image_url":         "",
		"ingredients": []gin.H{
			{"name": "cucumber", "amount": "1", "unit": "whole"},
			{"name": "cherry tomatoes", "amount": "100", "unit": "g"},
			{"name": "feta cheese", "amount": "30", "unit": "g"},
			{"name": "grilled chicken", "amount": "120", "unit": "g"},
		},
		"instructions": []string{
			"Dice cucumber and halve tomatoes.",
			"Combine vegetables in a bowl.",
			"Add crumbled feta and sliced grilled chicken.",
			"Drizzle with lemon juice and a teaspoon of olive oil.",
		},
		"nutrition": gin.H{
			"calories":  320,
			"protein_g": 35.0,
			"carbs_g":   12.0,
			"fats_g":    14.0,
			"fiber_g":   2.5,
		},
		"is_favorited": false,
	},
}

// GET /api/recipes
func ListRecipes(c *gin.Context) {
	goal := c.Query("goal")   // "weight_gain" | "weight_loss"
	search := c.Query("search")

	// TODO: replace with DB query using goal + search filters
	result := []gin.H{}
	for _, r := range mockRecipes {
		if goal != "" && r["goal_type"] != goal {
			continue
		}
		if search != "" {
			name, _ := r["name"].(string)
			if len(name) < len(search) {
				continue
			}
			// simple case-insensitive contains check (stub)
			_ = search // TODO: real filtering
		}
		result = append(result, r)
	}

	c.JSON(http.StatusOK, result)
}

// GET /api/recipes/:id
func GetRecipe(c *gin.Context) {
	id := c.Param("id")

	// TODO: look up recipe by id in DB
	for _, r := range mockRecipes {
		rid := r["id"]
		if rid == 1 && id == "1" || rid == 2 && id == "2" {
			c.JSON(http.StatusOK, r)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
}
