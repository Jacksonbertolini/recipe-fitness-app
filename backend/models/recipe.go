package models

import (
	"database/sql"
	"fitmeals/database"
	"time"
)

// ─── Structs ─────────────────────────────────────────────────────────────────

type Nutrition struct {
	Calories  int     `json:"calories"`
	ProteinG  float64 `json:"protein_g"`
	CarbsG    float64 `json:"carbs_g"`
	FatsG     float64 `json:"fats_g"`
	FiberG    float64 `json:"fiber_g"`
}

type Recipe struct {
	ID               int        `json:"id"`
	Name             string     `json:"name"`
	Description      string     `json:"description"`
	GoalType         string     `json:"goal_type"`
	PrepTimeMinutes  int        `json:"prep_time_minutes"`
	CookTimeMinutes  int        `json:"cook_time_minutes"`
	Servings         int        `json:"servings"`
	Ingredients      string     `json:"ingredients"`  // raw JSON string
	Instructions     string     `json:"instructions"` // raw JSON string
	ImageURL         string     `json:"image_url"`
	CreatedAt        time.Time  `json:"created_at"`
	Nutrition        *Nutrition `json:"nutrition,omitempty"`
	IsFavorited      bool       `json:"is_favorited"`
}

// ─── Queries ──────────────────────────────────────────────────────────────────

// ListRecipes returns all recipes (+ nutrition), optionally filtered.
// goal and search are both optional ("" = no filter).
func ListRecipes(goal, search string) ([]Recipe, error) {
	query := `
		SELECT r.id, r.name, r.description, r.goal_type,
		       r.prep_time_minutes, r.cook_time_minutes, r.servings,
		       r.ingredients, r.instructions, COALESCE(r.image_url, ''),
		       r.created_at,
		       n.calories, n.protein_g, n.carbs_g, n.fats_g, n.fiber_g
		FROM recipes r
		LEFT JOIN recipe_nutrition n ON n.recipe_id = r.id
		WHERE (? = '' OR r.goal_type = ?)
		  AND (? = '' OR r.name LIKE ?)
		ORDER BY r.id`

	searchLike := "%" + search + "%"
	rows, err := database.DB.Query(query, goal, goal, search, searchLike)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []Recipe
	for rows.Next() {
		r, n := Recipe{}, Nutrition{}
		err := rows.Scan(
			&r.ID, &r.Name, &r.Description, &r.GoalType,
			&r.PrepTimeMinutes, &r.CookTimeMinutes, &r.Servings,
			&r.Ingredients, &r.Instructions, &r.ImageURL, &r.CreatedAt,
			&n.Calories, &n.ProteinG, &n.CarbsG, &n.FatsG, &n.FiberG,
		)
		if err != nil {
			return nil, err
		}
		r.Nutrition = &n
		recipes = append(recipes, r)
	}
	return recipes, rows.Err()
}

// GetRecipeByID returns a single recipe with nutrition, or nil if not found.
func GetRecipeByID(id int) (*Recipe, error) {
	r, n := &Recipe{}, Nutrition{}
	row := database.DB.QueryRow(`
		SELECT r.id, r.name, r.description, r.goal_type,
		       r.prep_time_minutes, r.cook_time_minutes, r.servings,
		       r.ingredients, r.instructions, COALESCE(r.image_url, ''),
		       r.created_at,
		       n.calories, n.protein_g, n.carbs_g, n.fats_g, n.fiber_g
		FROM recipes r
		LEFT JOIN recipe_nutrition n ON n.recipe_id = r.id
		WHERE r.id = ?`, id)

	err := row.Scan(
		&r.ID, &r.Name, &r.Description, &r.GoalType,
		&r.PrepTimeMinutes, &r.CookTimeMinutes, &r.Servings,
		&r.Ingredients, &r.Instructions, &r.ImageURL, &r.CreatedAt,
		&n.Calories, &n.ProteinG, &n.CarbsG, &n.FatsG, &n.FiberG,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	r.Nutrition = &n
	return r, nil
}

// IsFavorited checks if a user has favorited a recipe.
func IsFavorited(userID, recipeID int) (bool, error) {
	var count int
	err := database.DB.QueryRow(
		`SELECT COUNT(*) FROM user_favorites WHERE user_id = ? AND recipe_id = ?`,
		userID, recipeID,
	).Scan(&count)
	return count > 0, err
}
