package main

import (
	"fitmeals/handlers"
	"fitmeals/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "FitMeals API running"})
	})

	api := r.Group("/api")

	// Public: Auth
	api.POST("/auth/register", handlers.Register)
	api.POST("/auth/login", handlers.Login)

	// Public: Recipes
	api.GET("/recipes", handlers.ListRecipes)
	api.GET("/recipes/:id", handlers.GetRecipe)

	// Protected routes
	protected := api.Group("/")
	protected.Use(middleware.AuthRequired())

	protected.GET("/user/profile", handlers.GetProfile)
	protected.PUT("/user/goal", handlers.UpdateGoal)

	protected.GET("/favorites", handlers.ListFavorites)
	protected.POST("/favorites/:recipeId", handlers.AddFavorite)
	protected.DELETE("/favorites/:recipeId", handlers.RemoveFavorite)

	log.Println("FitMeals API running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
